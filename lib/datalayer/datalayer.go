package datalayer

import (
	"fmt"
	"time"

	"github.com/dollarshaveclub/furan/generated/lib"
	"github.com/dollarshaveclub/furan/lib/db"
	"github.com/gocql/gocql"
	"github.com/golang/protobuf/proto"
)

// DataLayer describes an object that interacts with the persistant data store
type DataLayer interface {
	CreateBuild(*lib.BuildRequest) (gocql.UUID, error)
	GetBuildByID(gocql.UUID) (*lib.BuildStatusResponse, error)
	SetBuildFlags(gocql.UUID, map[string]bool) error
	SetBuildCompletedTimestamp(gocql.UUID) error
	SetBuildState(gocql.UUID, lib.BuildStatusResponse_BuildState) error
	DeleteBuild(gocql.UUID) error
	SetBuildTimeMetric(gocql.UUID, string) error
	SetDockerImageSizesMetric(gocql.UUID, int64, int64) error
	SaveBuildOutput(gocql.UUID, []lib.BuildEvent, string) error
	GetBuildOutput(gocql.UUID, string) ([]lib.BuildEvent, error)
}

// DBLayer is an DataLayer instance that interacts with the Cassandra database
type DBLayer struct {
	s *gocql.Session
}

// NewDBLayer returns a data layer object
func NewDBLayer(s *gocql.Session) *DBLayer {
	return &DBLayer{s: s}
}

// CreateBuild inserts a new build into the DB returning the ID
func (dl *DBLayer) CreateBuild(req *lib.BuildRequest) (gocql.UUID, error) {
	q := `INSERT INTO builds_by_id (id, request, state, finished, failed, cancelled, started)
        VALUES (?,{github_repo: ?, dockerfile_path: ?, tags: ?, tag_with_commit_sha: ?, ref: ?,
					push_registry_repo: ?, push_s3_region: ?, push_s3_bucket: ?,
					push_s3_key_prefix: ?},?,?,?,?,?);`
	id, err := gocql.RandomUUID()
	if err != nil {
		return id, err
	}
	udt := db.UDTFromBuildRequest(req)
	err = dl.s.Query(q, id, udt.GithubRepo, udt.DockerfilePath, udt.Tags, udt.TagWithCommitSha, udt.Ref,
		udt.PushRegistryRepo, udt.PushS3Region, udt.PushS3Bucket, udt.PushS3KeyPrefix,
		lib.BuildStatusResponse_STARTED.String(), false, false, false, time.Now()).Exec()
	if err != nil {
		return id, err
	}
	q = `INSERT INTO build_metrics_by_id (id) VALUES (?);`
	err = dl.s.Query(q, id).Exec()
	if err != nil {
		return id, err
	}
	q = `INSERT INTO build_events_by_id (id) VALUES (?);`
	return id, dl.s.Query(q, id).Exec()
}

// GetBuildByID fetches a build object from the DB
func (dl *DBLayer) GetBuildByID(id gocql.UUID) (*lib.BuildStatusResponse, error) {
	q := `SELECT request, state, finished, failed, cancelled, started, completed,
	      duration FROM builds_by_id WHERE id = ?;`
	var udt db.BuildRequestUDT
	var state string
	var started, completed time.Time
	bi := &lib.BuildStatusResponse{
		BuildId: id.String(),
	}
	err := dl.s.Query(q, id).Scan(&udt, &state, &bi.Finished, &bi.Failed,
		&bi.Cancelled, &started, &completed, &bi.Duration)
	if err != nil {
		return bi, err
	}
	bi.State = db.BuildStateFromString(state)
	bi.BuildRequest = db.BuildRequestFromUDT(&udt)
	bi.Started = started.Format(time.RFC3339)
	bi.Completed = completed.Format(time.RFC3339)
	return bi, nil
}

// SetBuildFlags sets the boolean flags on the build object
// Caller must ensure that the flags passed in are valid
func (dl *DBLayer) SetBuildFlags(id gocql.UUID, flags map[string]bool) error {
	var err error
	q := `UPDATE builds_by_id SET %v = ? WHERE id = ?;`
	for k, v := range flags {
		err = dl.s.Query(fmt.Sprintf(q, k), v, id).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBuildCompletedTimestamp sets the completed timestamp on a build to time.Now()
func (dl *DBLayer) SetBuildCompletedTimestamp(id gocql.UUID) error {
	var started time.Time
	now := time.Now()
	q := `SELECT started FROM builds_by_id WHERE id = ?;`
	err := dl.s.Query(q, id).Scan(&started)
	if err != nil {
		return err
	}
	duration := now.Sub(started).Seconds()
	q = `UPDATE builds_by_id SET completed = ?, duration = ? WHERE id = ?;`
	return dl.s.Query(q, now, duration, id).Exec()
}

// SetBuildState sets the state of a build
func (dl *DBLayer) SetBuildState(id gocql.UUID, state lib.BuildStatusResponse_BuildState) error {
	q := `UPDATE builds_by_id SET state = ? WHERE id = ?;`
	return dl.s.Query(q, state.String(), id).Exec()
}

// DeleteBuild removes a build from the DB.
// Only used in case of queue full when we can't actually do a build
func (dl *DBLayer) DeleteBuild(id gocql.UUID) error {
	q := `DELETE FROM builds_by_id WHERE id = ?;`
	err := dl.s.Query(q, id).Exec()
	if err != nil {
		return err
	}
	q = `DELETE FROM build_metrics_by_id WHERE id = ?;`
	return dl.s.Query(q, id).Exec()
}

// SetBuildTimeMetric sets a build metric to time.Now()
// metric is the name of the column to update
// if metric is a *_completed column, it will also compute and persist the duration
func (dl *DBLayer) SetBuildTimeMetric(id gocql.UUID, metric string) error {
	var started time.Time
	now := time.Now()
	getstarted := true
	var startedcolumn string
	var durationcolumn string
	switch metric {
	case "docker_build_completed":
		startedcolumn = "docker_build_started"
		durationcolumn = "docker_build_duration"
	case "push_completed":
		startedcolumn = "push_started"
		durationcolumn = "push_duration"
	case "clean_completed":
		startedcolumn = "clean_started"
		durationcolumn = "clean_duration"
	default:
		getstarted = false
	}
	q := `UPDATE build_metrics_by_id SET %v = ? WHERE id = ?;`
	err := dl.s.Query(fmt.Sprintf(q, metric), now, id).Exec()
	if err != nil {
		return err
	}
	if getstarted {
		q := `SELECT %v FROM build_metrics_by_id WHERE id = ?;`
		err := dl.s.Query(fmt.Sprintf(q, startedcolumn), id).Scan(&started)
		if err != nil {
			return err
		}
		duration := now.Sub(started).Seconds()
		q = `UPDATE build_metrics_by_id SET %v = ? WHERE id = ?;`
		return dl.s.Query(fmt.Sprintf(q, durationcolumn), duration, id).Exec()
	}
	return nil
}

// SetDockerImageSizesMetric sets the docker image sizes for a build
func (dl *DBLayer) SetDockerImageSizesMetric(id gocql.UUID, size int64, vsize int64) error {
	q := `UPDATE build_metrics_by_id SET docker_image_size = ?, docker_image_vsize = ? WHERE id = ?;`
	return dl.s.Query(q, size, vsize, id).Exec()
}

// SaveBuildOutput serializes an array of stream events to the database
func (dl *DBLayer) SaveBuildOutput(id gocql.UUID, output []lib.BuildEvent, column string) error {
	serialized := make([][]byte, len(output))
	var err error
	var b []byte
	for i, e := range output {
		b, err = proto.Marshal(&e)
		if err != nil {
			return err
		}
		serialized[i] = b
	}
	q := `UPDATE build_events_by_id SET %v = ? WHERE id = ?;`
	return dl.s.Query(fmt.Sprintf(q, column), serialized, id.String()).Exec()
}

// GetBuildOutput returns an array of stream events from the database
func (dl *DBLayer) GetBuildOutput(id gocql.UUID, column string) ([]lib.BuildEvent, error) {
	var rawoutput [][]byte
	output := []lib.BuildEvent{}
	q := `SELECT %v FROM build_events_by_id WHERE id = ?;`
	err := dl.s.Query(fmt.Sprintf(q, column), id).Scan(&rawoutput)
	if err != nil {
		return output, err
	}
	for _, rawevent := range rawoutput {
		event := lib.BuildEvent{}
		err = proto.Unmarshal(rawevent, &event)
		if err != nil {
			return output, err
		}
		output = append(output, event)
	}
	return output, nil
}
