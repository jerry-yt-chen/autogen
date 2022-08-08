# autogen

## Requirements
Before installing make sure you have the required dependencies installed:
- Go version v1.17

## Installation
```sh
go install github.com/17media/autogen@latest
```

## Create new project:

```shell
//support type: 'category', 'comp' 
autogen new <project-name> --type=<project-type>
```

## Test

- category (API service type)
```shell
autogen new coke --type=category
cd coke
make stop && make run
curl --request GET http://localhost:3000/api/v1/example
```

- category (API service type)
```shell
autogen new coke-comp --type=comp
cd coke-comp
make stop && make run
grpcurl -d '{"iam": "John"}' -plaintext  localhost:50051 greeter.v1.GreeterService/WhoAreYou
```
