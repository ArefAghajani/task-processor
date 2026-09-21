package worker

import (
	"errors"
	"fmt"

	"task-processor/task"
)

type Processor interface {
	Process(task.Task) error
}

type SimpleProcessor struct{}

func (p SimpleProcessor) Process(t task.Task) error {
	if t.Name == "" {
		return errors.New("task name is not defined")
	}

	return nil
}

type Worker struct {
	ID        int
	Processor Processor
}

func NewWorker(id int, processor Processor) Worker {
	return Worker{
		ID:        id,
		Processor: processor,
	}
}

func (w Worker) Run(tasks <-chan task.Task, results chan<- error) {
	for t := range tasks {
		fmt.Printf("Worker %d → Processing task %d\n", w.ID, t.ID)

		err := w.Processor.Process(t)

		results <- err
	}
}
