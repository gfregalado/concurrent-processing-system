package main

import (
	"fmt"
	"github.com/gfregalado/concurrent-processing-system/task"
	"github.com/gfregalado/concurrent-processing-system/worker"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const numWorkers = 5
	const numTasks = 100

	tasks := make(chan *task.Task, numTasks)
	results := make(chan string, numTasks)
	errors := make(chan error, numTasks)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker.Handler(tasks, results, errors, &wg)
	}

	// Generate tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- &task.Task{
			Id:       i,
			Duration: time.Duration(rand.Intn(5-1)+1) * time.Second,
		}
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	doneResults := make(chan bool)
	doneErrors := make(chan bool)

	go func() {
		for res := range results {
			fmt.Println(res)
		}
		doneResults <- true
	}()

	go func() {
		for err := range errors {
			fmt.Println("Error:", err)
		}
		doneErrors <- true
	}()

	// Wait for both result handling and error handling to complete
	<-doneResults
	<-doneErrors
}
