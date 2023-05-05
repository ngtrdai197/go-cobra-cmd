# Go Cobra

Go version: go 1.20

## Directory structure

```
├── main.go
├── .env
├── config
│   ├── config.go
├── cmd
│   ├── root.go
│   └── worker.go
├── pkg
│   ├── redis
│   └── worker
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
