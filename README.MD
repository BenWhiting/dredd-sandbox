# dredd-sandbox

This repo is a sandbox for learning how Dredd works

## HOW TO DREDD

### Reading

https://github.com/apiaryio/dredd
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md
https://dredd.org/en/latest/quickstart.html#document-your-api

### Repository layout

cmd/server.go - RESTful golang server with two endpoints
Dockerfile - docker file for the server.go
docker-compose.yml - starts both the golang server under test and the Dredd container
makefile - a simple go build and compose up command
swagger.yml - openAPI 2.0 document for server.go

### Running the tests

```bash
    make build
    make compose
```

### What just happened

make build - Creates the server.go executable
make compose - starts up both containers and runs the Dredd test against the swagger.yml file

There should be one passing and one failing test due to incorrect return bodies
