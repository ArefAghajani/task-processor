package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"task-processor/task"
	"task-processor/worker"
)

func main() {
	tasksCount := flag.Int("tasks", 20, "number of tasks")
	workersCount := flag.Int("workers", 3, "number of workers")

	flag.Parse()

	if *tasksCount < 1 {
		fmt.Println("tasks must be greater than 0")
		os.Exit(1)
	}

	if *workersCount < 1 {
		fmt.Println("workers must be greater than 0")
		os.Exit(1)
	}

	fmt.Println("Task Processor")
	fmt.Println("----------------")
	fmt.Printf("Tasks: %d\n", *tasksCount)
	fmt.Printf("Workers: %d\n\n", *workersCount)

	tasks := make(chan task.Task)
	results := make(chan error, *tasksCount)

	processor := worker.SimpleProcessor{}

	var wg sync.WaitGroup

	for i := 1; i <= *workersCount; i++ {
		w := worker.NewWorker(i, processor)

		wg.Add(1)

		go func(w worker.Worker) {
			defer wg.Done()

			w.Run(tasks, results)
		}(w)
	}

	go func() {
		defer close(tasks)

		for i := 1; i <= *tasksCount; i++ {
			t := task.NewTask(
				i,
				fmt.Sprintf("task %d", i),
			)

			tasks <- t
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	failed := 0

	for err := range results {
		if err != nil {
			failed++
			fmt.Printf("Task failed: %v\n", err)
		}
	}

	fmt.Println()
	fmt.Printf("Completed: %d\n", *tasksCount-failed)
	fmt.Printf("Failed: %d\n", failed)
}
