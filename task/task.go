package task

type Task struct {
	ID   int
	Name string
}

func NewTask(id int, name string) Task {
	return Task{
		ID:   id,
		Name: name,
	}
}
