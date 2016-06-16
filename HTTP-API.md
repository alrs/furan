HTTP API
========

The native API for Furan is [gRPC](http://www.grpc.io) however a RESTful HTTPS
adapter is available for testing convenience.


POST ``/build``
---------------

Body:

```javascript
{
  "build": {
    "github_repo": "dollarshaveclub/foobar",
    "dockerfile_path": "Dockerfile"  // optional: defaults to 'Dockerfile'
    "tags": ["master"],   // tags only (not including repo/name)
    "tag_with_commit_sha": true,
    "ref": "master",   // commit SHA or branch or tag
  },
  "push": {
    "registry": {
      "repo": "quay.io/dollarshaveclub/foobar"
    },
    // OR
    "s3": {
      "region": "us-west-2",
      "bucket": "foobar",
      "key_prefix": "myfolder/"
    }
  }
}
```

Response:

```json
{
  "buildId": "56aeda1c-6736-4657-8440-77972b103fee",
  "error": {}
}
```

GET ``/build/{id}``
-------------------

Response:

```javascript
{
  "build_id": "56aeda1c-6736-4657-8440-77972b103fee",
  "request": {
    "build": {
      "github_repo": "dollarshaveclub/foobar",
      "tags": ["master"],
      "tag_with_commit_sha": true,
      "ref": "master",
    },
    "push": {
      "registry": {
        "repo": "quay.io/dollarshaveclub/foobar"
      }
    },
  "state": "building",
  "failed": false,
  "started": "2016-05-19T07:33:54.691073",
  "completed": "",
  "duration": ""
}
```

Possible states:
  - building
  - pushing
  - success
  - buildFailure
  - pushFailure


GET ``http://localhost:{healthcheck port}/health``
---------------

Healthcheck. Listens on localhost only.

Response:

- 200 OK: everything is fine
- 429 Too Many Requests: queue is full
