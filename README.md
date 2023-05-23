# Go Cobra - Microservices Architecture

Go version: go 1.20

## Directory structure

```
├── cmd
│   ├── root.go
│   ├── worker.go
│   ├── gapi.go
│   ├── client.go
│   ├── kafka-consumer.go
│   ├── kafka-producer.go
│   ├── public-api.go
├── config
│   ├── config.go
├── util
│   ├── random.go
├── db
│   ├── migrations
│   │   ├── example_schema.down.sql
│   │   ├── example_schema.up.sql
│   ├── queries
│   │   ├── example.sql
│   ├── sqlc
│   │   ├── ...
├── pkg
│   ├── grpc
│   │   ├── pb
│   │   │   ├── example.pb.go
│   │   │   ├── ...
│   │   ├── proto
│   │   │   ├── example.proto
│   │   │   ├── ...
│   │   ├── converter.go
│   │   ├── server.go
│   │   ├── example.go
│   ├── redis
│   │   ├── redis.go
│   ├── kafka
│   │   ├── consumer.go
│   │   ├── producer.go
│   │   ├── ...
│   ├── public-api
│   │   ├── server.go
│   │   ├── example.go
│   │   ├── ...
│   ├── worker
│   │   ├── distributor.go
│   │   ├── processor.go
│   │   ├── email_delivery.go
│   │   ├── ...
│
├── main.go
├── .env.example
├── Makefile
├── sqlc.yaml
├── ...
...
```

## Getting started

### Install Dependencies

From the project root, run:

```shell
go build ./...
go mod tidy
```

- Run public api

```shell
 make public-api
```

- Run worker

```shell
make worker
```

### Setup gRPC

```shell
# Macos
brew install protobuf

# I’ll make sure it’s installed:
protoc --version

# Install the protocol compiler plugins for Go using the following commands:
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

# Update your PATH so that the protoc compiler can find the plugins:
export PATH="$PATH:$(go env GOPATH)/bin"
```

If you want to test if the methods in the service of gRPC are running, you can install [BloomRPC Github](https://github.com/bloomrpc/bloomrpc) to try or [Home brew](https://formulae.brew.sh/cask/bloomrpc)
