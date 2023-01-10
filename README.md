# Todo Cli App

A Todo CLI app made with Go and SQLite.

## Run and build

Your need install Go (1.19+) and SQLite3. For run just clone this project and exec:

```go
go run .
```

This create the database file in project root. For build exec:

```go
go build -tags prod -o ./build/gtask_v2.0
```

Now the database file will be create in the path of executable.

## Commands

- gtask list
- gtask add `task`
- gtask remove `task_id`
