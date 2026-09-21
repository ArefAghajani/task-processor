package task

import "testing"

func TestNewTask(t *testing.T) {
	got := NewTask(1, "task 1")

	if got.ID != 1 {
		t.Errorf("expected ID 1, got %d", got.ID)
	}

	if got.Name != "task 1" {
		t.Errorf("expected name %q, got %q", "task 1", got.Name)
	}
}
