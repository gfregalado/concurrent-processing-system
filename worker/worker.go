package worker

import (
	"fmt"
	"github.com/gfregalado/concurrent-processing-system/task"
	"sync"
)

func Handler(tasks <-chan *task.Task, results chan<- string, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tasks {
		func() {
			defer func() {
				if r := recover(); r != nil {
					errors <- fmt.Errorf("task %d panicked", t.Id)
				}
			}()
			result := t.Execute()
			results <- result
		}()
	}
}
