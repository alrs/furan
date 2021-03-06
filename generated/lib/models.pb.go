// Code generated by protoc-gen-go.
// source: models.proto
// DO NOT EDIT!

/*
Package lib is a generated protocol buffer package.

It is generated from these files:
	models.proto

It has these top-level messages:
	BuildDefinition
	PushRegistryDefinition
	PushS3Definition
	PushDefinition
	BuildRequest
	BuildStatusRequest
	BuildCancelRequest
	BuildRequestResponse
	BuildCancelResponse
	BuildStatusResponse
	BuildEventError
	BuildEvent
*/
package lib

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildStatusResponse_BuildState int32

const (
	BuildStatusResponse_STARTED       BuildStatusResponse_BuildState = 0
	BuildStatusResponse_BUILDING      BuildStatusResponse_BuildState = 1
	BuildStatusResponse_PUSHING       BuildStatusResponse_BuildState = 2
	BuildStatusResponse_SUCCESS       BuildStatusResponse_BuildState = 3
	BuildStatusResponse_BUILD_FAILURE BuildStatusResponse_BuildState = 4
	BuildStatusResponse_PUSH_FAILURE  BuildStatusResponse_BuildState = 5
	BuildStatusResponse_NOT_NECESSARY BuildStatusResponse_BuildState = 6
)

var BuildStatusResponse_BuildState_name = map[int32]string{
	0: "STARTED",
	1: "BUILDING",
	2: "PUSHING",
	3: "SUCCESS",
	4: "BUILD_FAILURE",
	5: "PUSH_FAILURE",
	6: "NOT_NECESSARY",
}
var BuildStatusResponse_BuildState_value = map[string]int32{
	"STARTED":       0,
	"BUILDING":      1,
	"PUSHING":       2,
	"SUCCESS":       3,
	"BUILD_FAILURE": 4,
	"PUSH_FAILURE":  5,
	"NOT_NECESSARY": 6,
}

func (x BuildStatusResponse_BuildState) String() string {
	return proto.EnumName(BuildStatusResponse_BuildState_name, int32(x))
}
func (BuildStatusResponse_BuildState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9, 0}
}

type BuildEventError_ErrorType int32

const (
	BuildEventError_FATAL     BuildEventError_ErrorType = 0
	BuildEventError_TEMPORARY BuildEventError_ErrorType = 1
	BuildEventError_NO_ERROR  BuildEventError_ErrorType = 2
)

var BuildEventError_ErrorType_name = map[int32]string{
	0: "FATAL",
	1: "TEMPORARY",
	2: "NO_ERROR",
}
var BuildEventError_ErrorType_value = map[string]int32{
	"FATAL":     0,
	"TEMPORARY": 1,
	"NO_ERROR":  2,
}

func (x BuildEventError_ErrorType) String() string {
	return proto.EnumName(BuildEventError_ErrorType_name, int32(x))
}
func (BuildEventError_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

type BuildEvent_EventType int32

const (
	BuildEvent_DOCKER_BUILD_STREAM BuildEvent_EventType = 0
	BuildEvent_DOCKER_PUSH_STREAM  BuildEvent_EventType = 1
	BuildEvent_LOG                 BuildEvent_EventType = 2
)

var BuildEvent_EventType_name = map[int32]string{
	0: "DOCKER_BUILD_STREAM",
	1: "DOCKER_PUSH_STREAM",
	2: "LOG",
}
var BuildEvent_EventType_value = map[string]int32{
	"DOCKER_BUILD_STREAM": 0,
	"DOCKER_PUSH_STREAM":  1,
	"LOG":                 2,
}

func (x BuildEvent_EventType) String() string {
	return proto.EnumName(BuildEvent_EventType_name, int32(x))
}
func (BuildEvent_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type BuildDefinition struct {
	GithubRepo       string   `protobuf:"bytes,1,opt,name=github_repo,json=githubRepo" json:"github_repo,omitempty"`
	DockerfilePath   string   `protobuf:"bytes,2,opt,name=dockerfile_path,json=dockerfilePath" json:"dockerfile_path,omitempty"`
	Ref              string   `protobuf:"bytes,3,opt,name=ref" json:"ref,omitempty"`
	Tags             []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	TagWithCommitSha bool     `protobuf:"varint,5,opt,name=tag_with_commit_sha,json=tagWithCommitSha" json:"tag_with_commit_sha,omitempty"`
}

func (m *BuildDefinition) Reset()                    { *m = BuildDefinition{} }
func (m *BuildDefinition) String() string            { return proto.CompactTextString(m) }
func (*BuildDefinition) ProtoMessage()               {}
func (*BuildDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildDefinition) GetGithubRepo() string {
	if m != nil {
		return m.GithubRepo
	}
	return ""
}

func (m *BuildDefinition) GetDockerfilePath() string {
	if m != nil {
		return m.DockerfilePath
	}
	return ""
}

func (m *BuildDefinition) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *BuildDefinition) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *BuildDefinition) GetTagWithCommitSha() bool {
	if m != nil {
		return m.TagWithCommitSha
	}
	return false
}

type PushRegistryDefinition struct {
	Repo string `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
}

func (m *PushRegistryDefinition) Reset()                    { *m = PushRegistryDefinition{} }
func (m *PushRegistryDefinition) String() string            { return proto.CompactTextString(m) }
func (*PushRegistryDefinition) ProtoMessage()               {}
func (*PushRegistryDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PushRegistryDefinition) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

type PushS3Definition struct {
	Region    string `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	Bucket    string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty"`
	KeyPrefix string `protobuf:"bytes,3,opt,name=key_prefix,json=keyPrefix" json:"key_prefix,omitempty"`
}

func (m *PushS3Definition) Reset()                    { *m = PushS3Definition{} }
func (m *PushS3Definition) String() string            { return proto.CompactTextString(m) }
func (*PushS3Definition) ProtoMessage()               {}
func (*PushS3Definition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushS3Definition) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *PushS3Definition) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *PushS3Definition) GetKeyPrefix() string {
	if m != nil {
		return m.KeyPrefix
	}
	return ""
}

type PushDefinition struct {
	Registry *PushRegistryDefinition `protobuf:"bytes,1,opt,name=registry" json:"registry,omitempty"`
	S3       *PushS3Definition       `protobuf:"bytes,2,opt,name=s3" json:"s3,omitempty"`
}

func (m *PushDefinition) Reset()                    { *m = PushDefinition{} }
func (m *PushDefinition) String() string            { return proto.CompactTextString(m) }
func (*PushDefinition) ProtoMessage()               {}
func (*PushDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PushDefinition) GetRegistry() *PushRegistryDefinition {
	if m != nil {
		return m.Registry
	}
	return nil
}

func (m *PushDefinition) GetS3() *PushS3Definition {
	if m != nil {
		return m.S3
	}
	return nil
}

type BuildRequest struct {
	Build        *BuildDefinition `protobuf:"bytes,1,opt,name=build" json:"build,omitempty"`
	Push         *PushDefinition  `protobuf:"bytes,2,opt,name=push" json:"push,omitempty"`
	SkipIfExists bool             `protobuf:"varint,3,opt,name=skip_if_exists,json=skipIfExists" json:"skip_if_exists,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BuildRequest) GetBuild() *BuildDefinition {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *BuildRequest) GetPush() *PushDefinition {
	if m != nil {
		return m.Push
	}
	return nil
}

func (m *BuildRequest) GetSkipIfExists() bool {
	if m != nil {
		return m.SkipIfExists
	}
	return false
}

type BuildStatusRequest struct {
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *BuildStatusRequest) Reset()                    { *m = BuildStatusRequest{} }
func (m *BuildStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildStatusRequest) ProtoMessage()               {}
func (*BuildStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BuildStatusRequest) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

type BuildCancelRequest struct {
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *BuildCancelRequest) Reset()                    { *m = BuildCancelRequest{} }
func (m *BuildCancelRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildCancelRequest) ProtoMessage()               {}
func (*BuildCancelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BuildCancelRequest) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

type BuildRequestResponse struct {
	BuildId string `protobuf:"bytes,2,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *BuildRequestResponse) Reset()                    { *m = BuildRequestResponse{} }
func (m *BuildRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildRequestResponse) ProtoMessage()               {}
func (*BuildRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BuildRequestResponse) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

type BuildCancelResponse struct {
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *BuildCancelResponse) Reset()                    { *m = BuildCancelResponse{} }
func (m *BuildCancelResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildCancelResponse) ProtoMessage()               {}
func (*BuildCancelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BuildCancelResponse) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

type BuildStatusResponse struct {
	BuildId      string                         `protobuf:"bytes,1,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	BuildRequest *BuildRequest                  `protobuf:"bytes,2,opt,name=build_request,json=buildRequest" json:"build_request,omitempty"`
	State        BuildStatusResponse_BuildState `protobuf:"varint,3,opt,name=state,enum=lib.BuildStatusResponse_BuildState" json:"state,omitempty"`
	Finished     bool                           `protobuf:"varint,4,opt,name=finished" json:"finished,omitempty"`
	Failed       bool                           `protobuf:"varint,5,opt,name=failed" json:"failed,omitempty"`
	Cancelled    bool                           `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Started      string                         `protobuf:"bytes,7,opt,name=started" json:"started,omitempty"`
	Completed    string                         `protobuf:"bytes,8,opt,name=completed" json:"completed,omitempty"`
	Duration     float64                        `protobuf:"fixed64,9,opt,name=duration" json:"duration,omitempty"`
}

func (m *BuildStatusResponse) Reset()                    { *m = BuildStatusResponse{} }
func (m *BuildStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildStatusResponse) ProtoMessage()               {}
func (*BuildStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BuildStatusResponse) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *BuildStatusResponse) GetBuildRequest() *BuildRequest {
	if m != nil {
		return m.BuildRequest
	}
	return nil
}

func (m *BuildStatusResponse) GetState() BuildStatusResponse_BuildState {
	if m != nil {
		return m.State
	}
	return BuildStatusResponse_STARTED
}

func (m *BuildStatusResponse) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func (m *BuildStatusResponse) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *BuildStatusResponse) GetCancelled() bool {
	if m != nil {
		return m.Cancelled
	}
	return false
}

func (m *BuildStatusResponse) GetStarted() string {
	if m != nil {
		return m.Started
	}
	return ""
}

func (m *BuildStatusResponse) GetCompleted() string {
	if m != nil {
		return m.Completed
	}
	return ""
}

func (m *BuildStatusResponse) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type BuildEventError struct {
	IsError   bool                      `protobuf:"varint,1,opt,name=is_error,json=isError" json:"is_error,omitempty"`
	ErrorType BuildEventError_ErrorType `protobuf:"varint,2,opt,name=error_type,json=errorType,enum=lib.BuildEventError_ErrorType" json:"error_type,omitempty"`
}

func (m *BuildEventError) Reset()                    { *m = BuildEventError{} }
func (m *BuildEventError) String() string            { return proto.CompactTextString(m) }
func (*BuildEventError) ProtoMessage()               {}
func (*BuildEventError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BuildEventError) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *BuildEventError) GetErrorType() BuildEventError_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return BuildEventError_FATAL
}

type BuildEvent struct {
	EventError    *BuildEventError     `protobuf:"bytes,1,opt,name=event_error,json=eventError" json:"event_error,omitempty"`
	EventType     BuildEvent_EventType `protobuf:"varint,2,opt,name=event_type,json=eventType,enum=lib.BuildEvent_EventType" json:"event_type,omitempty"`
	BuildId       string               `protobuf:"bytes,3,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	Message       string               `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	BuildFinished bool                 `protobuf:"varint,5,opt,name=build_finished,json=buildFinished" json:"build_finished,omitempty"`
}

func (m *BuildEvent) Reset()                    { *m = BuildEvent{} }
func (m *BuildEvent) String() string            { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()               {}
func (*BuildEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BuildEvent) GetEventError() *BuildEventError {
	if m != nil {
		return m.EventError
	}
	return nil
}

func (m *BuildEvent) GetEventType() BuildEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return BuildEvent_DOCKER_BUILD_STREAM
}

func (m *BuildEvent) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *BuildEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BuildEvent) GetBuildFinished() bool {
	if m != nil {
		return m.BuildFinished
	}
	return false
}

func init() {
	proto.RegisterType((*BuildDefinition)(nil), "lib.BuildDefinition")
	proto.RegisterType((*PushRegistryDefinition)(nil), "lib.PushRegistryDefinition")
	proto.RegisterType((*PushS3Definition)(nil), "lib.PushS3Definition")
	proto.RegisterType((*PushDefinition)(nil), "lib.PushDefinition")
	proto.RegisterType((*BuildRequest)(nil), "lib.BuildRequest")
	proto.RegisterType((*BuildStatusRequest)(nil), "lib.BuildStatusRequest")
	proto.RegisterType((*BuildCancelRequest)(nil), "lib.BuildCancelRequest")
	proto.RegisterType((*BuildRequestResponse)(nil), "lib.BuildRequestResponse")
	proto.RegisterType((*BuildCancelResponse)(nil), "lib.BuildCancelResponse")
	proto.RegisterType((*BuildStatusResponse)(nil), "lib.BuildStatusResponse")
	proto.RegisterType((*BuildEventError)(nil), "lib.BuildEventError")
	proto.RegisterType((*BuildEvent)(nil), "lib.BuildEvent")
	proto.RegisterEnum("lib.BuildStatusResponse_BuildState", BuildStatusResponse_BuildState_name, BuildStatusResponse_BuildState_value)
	proto.RegisterEnum("lib.BuildEventError_ErrorType", BuildEventError_ErrorType_name, BuildEventError_ErrorType_value)
	proto.RegisterEnum("lib.BuildEvent_EventType", BuildEvent_EventType_name, BuildEvent_EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FuranExecutor service

type FuranExecutorClient interface {
	StartBuild(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildRequestResponse, error)
	GetBuildStatus(ctx context.Context, in *BuildStatusRequest, opts ...grpc.CallOption) (*BuildStatusResponse, error)
	MonitorBuild(ctx context.Context, in *BuildStatusRequest, opts ...grpc.CallOption) (FuranExecutor_MonitorBuildClient, error)
	CancelBuild(ctx context.Context, in *BuildCancelRequest, opts ...grpc.CallOption) (*BuildCancelResponse, error)
}

type furanExecutorClient struct {
	cc *grpc.ClientConn
}

func NewFuranExecutorClient(cc *grpc.ClientConn) FuranExecutorClient {
	return &furanExecutorClient{cc}
}

func (c *furanExecutorClient) StartBuild(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildRequestResponse, error) {
	out := new(BuildRequestResponse)
	err := grpc.Invoke(ctx, "/lib.FuranExecutor/StartBuild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *furanExecutorClient) GetBuildStatus(ctx context.Context, in *BuildStatusRequest, opts ...grpc.CallOption) (*BuildStatusResponse, error) {
	out := new(BuildStatusResponse)
	err := grpc.Invoke(ctx, "/lib.FuranExecutor/GetBuildStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *furanExecutorClient) MonitorBuild(ctx context.Context, in *BuildStatusRequest, opts ...grpc.CallOption) (FuranExecutor_MonitorBuildClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FuranExecutor_serviceDesc.Streams[0], c.cc, "/lib.FuranExecutor/MonitorBuild", opts...)
	if err != nil {
		return nil, err
	}
	x := &furanExecutorMonitorBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FuranExecutor_MonitorBuildClient interface {
	Recv() (*BuildEvent, error)
	grpc.ClientStream
}

type furanExecutorMonitorBuildClient struct {
	grpc.ClientStream
}

func (x *furanExecutorMonitorBuildClient) Recv() (*BuildEvent, error) {
	m := new(BuildEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *furanExecutorClient) CancelBuild(ctx context.Context, in *BuildCancelRequest, opts ...grpc.CallOption) (*BuildCancelResponse, error) {
	out := new(BuildCancelResponse)
	err := grpc.Invoke(ctx, "/lib.FuranExecutor/CancelBuild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FuranExecutor service

type FuranExecutorServer interface {
	StartBuild(context.Context, *BuildRequest) (*BuildRequestResponse, error)
	GetBuildStatus(context.Context, *BuildStatusRequest) (*BuildStatusResponse, error)
	MonitorBuild(*BuildStatusRequest, FuranExecutor_MonitorBuildServer) error
	CancelBuild(context.Context, *BuildCancelRequest) (*BuildCancelResponse, error)
}

func RegisterFuranExecutorServer(s *grpc.Server, srv FuranExecutorServer) {
	s.RegisterService(&_FuranExecutor_serviceDesc, srv)
}

func _FuranExecutor_StartBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuranExecutorServer).StartBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.FuranExecutor/StartBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuranExecutorServer).StartBuild(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuranExecutor_GetBuildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuranExecutorServer).GetBuildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.FuranExecutor/GetBuildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuranExecutorServer).GetBuildStatus(ctx, req.(*BuildStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuranExecutor_MonitorBuild_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FuranExecutorServer).MonitorBuild(m, &furanExecutorMonitorBuildServer{stream})
}

type FuranExecutor_MonitorBuildServer interface {
	Send(*BuildEvent) error
	grpc.ServerStream
}

type furanExecutorMonitorBuildServer struct {
	grpc.ServerStream
}

func (x *furanExecutorMonitorBuildServer) Send(m *BuildEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _FuranExecutor_CancelBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuranExecutorServer).CancelBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.FuranExecutor/CancelBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuranExecutorServer).CancelBuild(ctx, req.(*BuildCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FuranExecutor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lib.FuranExecutor",
	HandlerType: (*FuranExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBuild",
			Handler:    _FuranExecutor_StartBuild_Handler,
		},
		{
			MethodName: "GetBuildStatus",
			Handler:    _FuranExecutor_GetBuildStatus_Handler,
		},
		{
			MethodName: "CancelBuild",
			Handler:    _FuranExecutor_CancelBuild_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorBuild",
			Handler:       _FuranExecutor_MonitorBuild_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "models.proto",
}

func init() { proto.RegisterFile("models.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 943 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x6d, 0x6f, 0xe2, 0x46,
	0x10, 0xc6, 0x40, 0x02, 0x1e, 0x08, 0xf1, 0x6d, 0xae, 0x39, 0xe7, 0xfa, 0x16, 0xb9, 0x3d, 0x5d,
	0x54, 0xb5, 0xf4, 0x0a, 0xea, 0x9b, 0x74, 0xfd, 0x40, 0x88, 0x93, 0xa2, 0x26, 0x01, 0x2d, 0x44,
	0x55, 0x3f, 0x59, 0x06, 0x2f, 0x78, 0xc5, 0x8b, 0x5d, 0xef, 0xba, 0x0d, 0x3f, 0xa1, 0x7f, 0xa0,
	0x9f, 0xef, 0x27, 0xf4, 0x47, 0xf4, 0x87, 0x55, 0xbb, 0x5e, 0xfc, 0x92, 0xdc, 0xa9, 0xf7, 0x05,
	0xed, 0x3c, 0x33, 0x8f, 0xe7, 0x99, 0x9d, 0x99, 0x05, 0x9a, 0xeb, 0xc0, 0x23, 0x2b, 0xd6, 0x0e,
	0xa3, 0x80, 0x07, 0xa8, 0xb2, 0xa2, 0x53, 0xeb, 0x1f, 0x0d, 0x0e, 0xcf, 0x63, 0xba, 0xf2, 0x2e,
	0xc8, 0x9c, 0x6e, 0x28, 0xa7, 0xc1, 0x06, 0x7d, 0x0a, 0x8d, 0x05, 0xe5, 0x7e, 0x3c, 0x75, 0x22,
	0x12, 0x06, 0xa6, 0x76, 0xaa, 0x9d, 0xe9, 0x18, 0x12, 0x08, 0x93, 0x30, 0x40, 0x2f, 0xe1, 0xd0,
	0x0b, 0x66, 0x4b, 0x12, 0xcd, 0xe9, 0x8a, 0x38, 0xa1, 0xcb, 0x7d, 0xb3, 0x2c, 0x83, 0x5a, 0x19,
	0x3c, 0x72, 0xb9, 0x8f, 0x0c, 0xa8, 0x44, 0x64, 0x6e, 0x56, 0xa4, 0x53, 0x1c, 0x11, 0x82, 0x2a,
	0x77, 0x17, 0xcc, 0xac, 0x9e, 0x56, 0xce, 0x74, 0x2c, 0xcf, 0xe8, 0x2b, 0x38, 0xe2, 0xee, 0xc2,
	0xf9, 0x93, 0x72, 0xdf, 0x99, 0x05, 0xeb, 0x35, 0xe5, 0x0e, 0xf3, 0x5d, 0x73, 0xef, 0x54, 0x3b,
	0xab, 0x63, 0x83, 0xbb, 0x8b, 0x5f, 0x29, 0xf7, 0xfb, 0xd2, 0x31, 0xf6, 0x5d, 0xeb, 0x4b, 0x38,
	0x1e, 0xc5, 0xcc, 0xc7, 0x64, 0x41, 0x19, 0x8f, 0xb6, 0x39, 0xe1, 0x08, 0xaa, 0x39, 0xc5, 0xf2,
	0x6c, 0xb9, 0x60, 0x88, 0xe8, 0x71, 0x37, 0x17, 0x77, 0x0c, 0xfb, 0x11, 0x59, 0xd0, 0x60, 0xa3,
	0x22, 0x95, 0x25, 0xf0, 0x69, 0x3c, 0x5b, 0x12, 0xae, 0xca, 0x51, 0x16, 0xfa, 0x18, 0x60, 0x49,
	0xb6, 0x4e, 0x18, 0x91, 0x39, 0xbd, 0x57, 0xd5, 0xe8, 0x4b, 0xb2, 0x1d, 0x49, 0xc0, 0x0a, 0xa1,
	0x25, 0x52, 0xe4, 0x12, 0x7c, 0x0f, 0xf5, 0x48, 0xc9, 0x93, 0x29, 0x1a, 0x9d, 0x0f, 0xdb, 0x2b,
	0x3a, 0x6d, 0xbf, 0x5d, 0x37, 0x4e, 0x83, 0xd1, 0x0b, 0x28, 0xb3, 0xae, 0xcc, 0xde, 0xe8, 0x7c,
	0x90, 0x52, 0xf2, 0xe2, 0x71, 0x99, 0x75, 0xad, 0xbf, 0x34, 0x68, 0xca, 0xae, 0x61, 0xf2, 0x7b,
	0x4c, 0x18, 0x47, 0x5f, 0xc0, 0xde, 0x54, 0xd8, 0x2a, 0xdb, 0x53, 0x49, 0x7d, 0xd0, 0x57, 0x9c,
	0x84, 0xa0, 0x97, 0x50, 0x0d, 0x63, 0xe6, 0xab, 0x2c, 0x47, 0x69, 0x96, 0x5c, 0xa4, 0x0c, 0x40,
	0x9f, 0x43, 0x8b, 0x2d, 0x69, 0xe8, 0xd0, 0xb9, 0x43, 0xee, 0x29, 0xe3, 0x4c, 0x96, 0x5e, 0xc7,
	0x4d, 0x81, 0x0e, 0xe6, 0xb6, 0xc4, 0xac, 0xaf, 0x01, 0xc9, 0x44, 0x63, 0xee, 0xf2, 0x98, 0xed,
	0x04, 0x9d, 0x40, 0x5d, 0x66, 0x73, 0xa8, 0xa7, 0x2e, 0xb9, 0x26, 0xed, 0x81, 0x97, 0x12, 0xfa,
	0xee, 0x66, 0x46, 0x56, 0xef, 0x41, 0xf8, 0x06, 0x9e, 0xe6, 0x8b, 0xc5, 0x84, 0x85, 0xc1, 0x86,
	0x91, 0x02, 0xa5, 0x5c, 0xa4, 0xbc, 0x82, 0xa3, 0x42, 0x8e, 0xb7, 0x30, 0x1e, 0x24, 0xf9, 0xb7,
	0xa2, 0x28, 0xbb, 0x3a, 0xfe, 0x97, 0x82, 0xbe, 0x83, 0x83, 0xc4, 0x15, 0x25, 0xc2, 0xd4, 0x8d,
	0x3e, 0xc9, 0x2e, 0x7f, 0xa7, 0xb8, 0x39, 0xcd, 0x37, 0xeb, 0x47, 0xd8, 0x63, 0xdc, 0xe5, 0x44,
	0x5e, 0x67, 0xab, 0xf3, 0x59, 0x16, 0x5f, 0xcc, 0x9d, 0x61, 0x04, 0x27, 0x0c, 0xf4, 0x1c, 0xea,
	0xa2, 0x49, 0xcc, 0x27, 0x9e, 0x59, 0x95, 0xcd, 0x48, 0x6d, 0x31, 0xbd, 0x73, 0x97, 0xae, 0x88,
	0xa7, 0x36, 0x47, 0x59, 0xe8, 0x23, 0xd0, 0x67, 0xf2, 0x1a, 0x84, 0x6b, 0x5f, 0xba, 0x32, 0x00,
	0x99, 0x50, 0x63, 0xdc, 0x8d, 0x38, 0xf1, 0xcc, 0x5a, 0x52, 0x9e, 0x32, 0x25, 0x2f, 0x58, 0x87,
	0x2b, 0x22, 0x7c, 0xf5, 0x64, 0xe8, 0x53, 0x40, 0x28, 0xf1, 0xe2, 0xc8, 0x15, 0xe3, 0x62, 0xea,
	0xa7, 0xda, 0x99, 0x86, 0x53, 0xdb, 0xda, 0x02, 0x64, 0xd2, 0x51, 0x03, 0x6a, 0xe3, 0x49, 0x0f,
	0x4f, 0xec, 0x0b, 0xa3, 0x84, 0x9a, 0x50, 0x3f, 0xbf, 0x1b, 0x5c, 0x5f, 0x0c, 0x6e, 0xaf, 0x0c,
	0x4d, 0xb8, 0x46, 0x77, 0xe3, 0x9f, 0x85, 0x51, 0x96, 0x71, 0x77, 0xfd, 0xbe, 0x3d, 0x1e, 0x1b,
	0x15, 0xf4, 0x04, 0x0e, 0x64, 0x9c, 0x73, 0xd9, 0x1b, 0x5c, 0xdf, 0x61, 0xdb, 0xa8, 0x22, 0x03,
	0x9a, 0x22, 0x38, 0x45, 0xf6, 0x44, 0xd0, 0xed, 0x70, 0xe2, 0xdc, 0xda, 0x82, 0xd4, 0xc3, 0xbf,
	0x19, 0xfb, 0xd6, 0x9b, 0xdd, 0x7b, 0x66, 0xff, 0x41, 0x36, 0xdc, 0x8e, 0xa2, 0x20, 0x12, 0x2d,
	0xa4, 0xcc, 0x21, 0xe2, 0x2c, 0x5b, 0x58, 0xc7, 0x35, 0xca, 0x12, 0xd7, 0x4f, 0x00, 0x12, 0x77,
	0xf8, 0x36, 0x24, 0xb2, 0x7f, 0xad, 0xce, 0x27, 0x59, 0x3f, 0xb2, 0x8f, 0xb4, 0xe5, 0xef, 0x64,
	0x1b, 0x12, 0xac, 0x93, 0xdd, 0xd1, 0xea, 0x82, 0x9e, 0xe2, 0x48, 0x87, 0xbd, 0xcb, 0xde, 0xa4,
	0x77, 0x6d, 0x94, 0xd0, 0x01, 0xe8, 0x13, 0xfb, 0x66, 0x34, 0xc4, 0x42, 0x94, 0x26, 0x8a, 0xbe,
	0x1d, 0x3a, 0x36, 0xc6, 0x43, 0x6c, 0x94, 0xad, 0x37, 0x65, 0x75, 0x3d, 0xf2, 0xeb, 0xe8, 0x5b,
	0x68, 0x10, 0x71, 0xc8, 0x09, 0x2c, 0x2c, 0x70, 0xa6, 0x01, 0x03, 0xc9, 0x8a, 0xfa, 0x01, 0x12,
	0x2b, 0xaf, 0xfc, 0xe4, 0x01, 0xab, 0x2d, 0x7f, 0x95, 0xe8, 0xdd, 0xb1, 0x30, 0xd1, 0x95, 0xe2,
	0x44, 0x9b, 0x50, 0x5b, 0x13, 0xc6, 0xdc, 0x05, 0x91, 0xd3, 0xa5, 0xe3, 0x9d, 0x89, 0x5e, 0x40,
	0x2b, 0x21, 0xa5, 0xe3, 0x97, 0x0c, 0x59, 0xb2, 0x01, 0x97, 0x0a, 0xb4, 0x6c, 0xd0, 0xd3, 0x9c,
	0xe8, 0x19, 0x1c, 0x5d, 0x0c, 0xfb, 0xbf, 0xd8, 0xd8, 0x49, 0x5a, 0x39, 0x9e, 0x60, 0xbb, 0x77,
	0x63, 0x94, 0xd0, 0x31, 0x20, 0xe5, 0x90, 0x0d, 0x55, 0xb8, 0x86, 0x6a, 0x50, 0xb9, 0x1e, 0x5e,
	0x19, 0xe5, 0xce, 0xdf, 0x65, 0x38, 0xb8, 0x8c, 0x23, 0x77, 0x63, 0xdf, 0x93, 0x59, 0xcc, 0x83,
	0x08, 0xbd, 0x06, 0x18, 0x8b, 0xb9, 0x94, 0xc5, 0xa1, 0xc7, 0x2b, 0xf6, 0xfc, 0xe4, 0xf1, 0xd6,
	0xa9, 0x35, 0xb2, 0x4a, 0xc8, 0x86, 0xd6, 0x15, 0xe1, 0xb9, 0x15, 0x43, 0xcf, 0x1e, 0x2f, 0x5d,
	0xf2, 0x1d, 0xf3, 0x5d, 0xdb, 0x68, 0x95, 0xd0, 0x6b, 0x68, 0xde, 0x04, 0x1b, 0xca, 0x83, 0x28,
	0x91, 0xf1, 0xce, 0x8f, 0x1c, 0x3e, 0x68, 0x84, 0x55, 0x7a, 0xa5, 0xa1, 0x73, 0x68, 0x24, 0xcf,
	0xd1, 0x23, 0x72, 0xe1, 0x25, 0xcc, 0x2b, 0x28, 0x3e, 0x5f, 0x56, 0x69, 0xba, 0x2f, 0xff, 0xba,
	0xbb, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x65, 0x0f, 0x24, 0xca, 0x07, 0x00, 0x00,
}
