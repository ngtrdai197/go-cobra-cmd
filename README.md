# Go Cobra

Go version: go 1.20

## Directory structure

```
├── cmd
│   ├── root.go
│   ├── worker.go
│   ├── client.go
│   ├── kafka-consumer.go
│   ├── kafka-producer.go
├── config
│   ├── config.go
├── pkg
│   ├── redis
│       ├── redis.go
│   └── worker
│       ├── distributor.go
│       ├── processor.go
│       ├── email_delivery.go
│       ├── ...
│
├── main.go
├── .env.example
├── Makefile
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

- Run worker

```shell
go run main.go worker-cmd
```
