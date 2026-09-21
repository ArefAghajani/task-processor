package worker

import (
	"testing"

	"task-processor/task"
)

func TestSimpleProcessorProcess(t *testing.T) {
	processor := SimpleProcessor{}

	t.Run("valid task", func(t *testing.T) {
		ta := task.NewTask(1, "task 1")

		err := processor.Process(ta)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("task without name", func(t *testing.T) {
		ta := task.NewTask(1, "")

		err := processor.Process(ta)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
