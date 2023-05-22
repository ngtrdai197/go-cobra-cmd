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
│   │    ├── distributor.go
│   │    ├── processor.go
│   │    ├── email_delivery.go
│   │    ├── ...
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
