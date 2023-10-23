package worker

import (
	"fmt"
	"github.com/gfregalado/concurrent-processing-system/task"
	"sync"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	numTasks := 5
	tasks := make(chan *task.Task, numTasks)
	results := make(chan string, numTasks)
	errors := make(chan error, numTasks)

	var wg sync.WaitGroup
	wg.Add(1)

	go Handler(tasks, results, errors, &wg)

	expectedResults := make(map[string]bool)
	for i := 1; i <= numTasks; i++ {
		executionMessage := fmt.Sprintf("Task %d executed", i)
		tasks <- &task.Task{Id: i, Duration: 1 * time.Second}
		expectedResults[executionMessage] = true
	}
	close(tasks)

	wg.Wait()
	close(results)
	close(errors)

	// Verifying results and errors.
	for res := range results {
		if _, ok := expectedResults[res]; !ok {
			t.Errorf("unexpected result: %v", res)
		} else {
			delete(expectedResults, res)
		}
	}

	if len(expectedResults) != 0 {
		t.Errorf("not all expected results were received")
	}

	if len(errors) > 0 {
		t.Errorf("expected no errors, got %d error(s)", len(errors))
	}
}
