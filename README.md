# autogen

## Requirements
Before installing make sure you have the required dependencies installed:
- Go version v1.17

## Installation
```sh
go install github.com/jerry-yt-chen/autogen@latest
```

## Create new project:

```sh
mkdir <project-name>
cd <project-name>
autogen new
```

## Directory Structure

```shell
├── Makefile
├── README.md
├── build
│   ├── Dockerfile
│   └── entrypoint.sh
├── cmd
│   └── main.go
├── configs
│   ├── config.go
│   └── config.yaml
├── deployments
│   └── docker-compose.yaml
├── go.mod
└── internal
    ├── dispatcher
    │   └── user
    │       ├── impl.go
    │       └── user.go
    ├── domain
    │   └── user
    │       ├── entity
    │       │   └── entity.go
    │       ├── model
    │       │   └── model.go
    │       └── repo
    │           ├── repository.go
    │           └── respositoryStub.go
    ├── framework
    │   ├── engine
    │   │   ├── engine.go
    │   │   └── gin
    │   │       ├── ginEngine.go
    │   │       └── render
    │   │           ├── error.go
    │   │           └── response.go
    │   ├── middlewares
    │   │   └── context.go
    │   └── router
    │       ├── ginRouter.go
    │       └── router.go
    ├── injector
    │   ├── api
    │   │   ├── dispatcher.go
    │   │   ├── httpServer.go
    │   │   ├── receiver.go
    │   │   ├── router.go
    │   │   └── translator.go
    │   ├── domain
    │   │   └── repository.go
    │   ├── injector.go
    │   ├── wire.go
    │   └── wire_gen.go
    ├── receiver
    │   ├── example
    │   │   ├── example.go
    │   │   └── impl.go
    │   └── receiver.go
    └── translator
        ├── example
        │   ├── example.go
        │   ├── impl.go
        │   ├── request
        │   │   └── getExampleRequest.go
        │   └── response
        │       └── getExampleResponse.go
        └── translator.go

```

## Test

```shell
cd <project-name>
make stop && make run
curl --request GET http://localhost:3000/api/v1/example
```

## TODO
- Support component service template
- cli should also support args
```shell
autogen new proj category
autoget new proj compnent
```