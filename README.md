# Task Processor

A simple CLI application written in Go for processing multiple tasks concurrently using a limited number of workers.

This project is mainly built to practice **goroutines, channels, interfaces, and the Worker Pool pattern** in Go.

## How it works

You can specify the number of tasks and workers:

```bash
go run ./cmd/taskprocessor -tasks=20 -workers=3
```

Example output:

```text
Task Processor
----------------
Tasks: 20
Workers: 3

Worker 1 → Processing task 1
Worker 2 → Processing task 2
Worker 3 → Processing task 3
Worker 1 → Processing task 4
Worker 2 → Processing task 5
...
Completed: 20
Failed: 0
```

The exact order of workers may change because the tasks are processed concurrently.

## Project Structure

```text
task-processor/
├── go.mod
├── README.md
│
├── cmd/
│   └── taskprocessor/
│       └── main.go
│
├── task/
│   ├── task.go
│   └── task_test.go
│
└── worker/
    ├── worker.go
    └── worker_test.go
```

### Main parts

* `task/` — Defines the `Task` model.
* `worker/` — Contains the Worker and Processor logic.
* `cmd/taskprocessor/` — Entry point of the CLI application.
* `*_test.go` — Unit tests.

## Usage

Run with default values:

```bash
go run ./cmd/taskprocessor
```

Default configuration:

```text
Tasks: 20
Workers: 3
```

Or provide your own values:

```bash
go run ./cmd/taskprocessor -tasks=50 -workers=5
```

Available options:

| Option     | Default | Description       |
| ---------- | ------: | ----------------- |
| `-tasks`   |    `20` | Number of tasks   |
| `-workers` |     `3` | Number of workers |

Both values must be greater than `0`.

## Concurrency

The application uses a simple Worker Pool:

```text
             Tasks
               │
               ▼
         ┌───────────┐
         │  Channel  │
         └─────┬─────┘
               │
       ┌───────┼───────┐
       ▼       ▼       ▼
    Worker 1 Worker 2 Worker 3
       │       │       │
       └───────┼───────┘
               ▼
            Results
```

Tasks are sent through a channel and available workers pick them up.

`sync.WaitGroup` is used to wait until all workers finish.

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with the race detector:

```bash
go test -race ./...
```

## What this project demonstrates

* Structs and methods
* Interfaces
* Error handling
* Goroutines
* Channels
* sync.WaitGroup
* Worker Pool pattern
* Basic unit testing
* CLI flags

## Future Improvements

Some possible improvements:

* Add task retry support
* Add graceful shutdown with context
* Add task processing time
* Add better logging
* Add task priorities
* Add configurable retry limits

