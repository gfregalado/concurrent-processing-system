package task

import (
	"fmt"
	"time"
)

type Task struct {
	Id       int
	Duration time.Duration
}

func (t *Task) Execute() string {
	time.Sleep(t.Duration)
	return fmt.Sprintf("Task %d executed", t.Id)
}
