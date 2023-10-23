package task

import (
	"testing"
	"time"
)

func TestTaskExecute(t *testing.T) {
	var tests = []struct {
		id       int
		duration time.Duration
		want     string
	}{
		{1, 1 * time.Second, "Task 1 executed"},
		{2, 2 * time.Second, "Task 2 executed"},
	}

	for _, tt := range tests {
		task := Task{Id: tt.id, Duration: tt.duration}

		start := time.Now()
		got := task.Execute()
		duration := time.Since(start)

		if got != tt.want {
			t.Errorf("Task.Execute() = %q; want %q", got, tt.want)
		}

		if duration < tt.duration || duration > tt.duration+10*time.Millisecond {
			t.Errorf("Task %d executed in %v; want %v", tt.id, duration, tt.duration)
		}
	}
}
