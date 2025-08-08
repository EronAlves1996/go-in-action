package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	ID   int
	Task func(int) error
}

func worker(id int,
	tasks <-chan Task,
	wg *sync.WaitGroup,
	_ *sync.Mutex,
	sharedCounter *int32,
) {
	defer wg.Done()
	for task := range tasks {
		err := task.Task(task.ID)
		if err != nil {
			fmt.Printf("Processed work id %d failed with error %s in worker %d\n", task.ID, err, id)
		} else {
			fmt.Printf("Processed work id %d in worker %d succesfully\n", task.ID, id)
		}
		atomic.AddInt32(sharedCounter, 1)
	}
}

func main() {
	const (
		numWorkers = 3
		maxTasks   = 10
		rateLimit  = 2 // tasks per second
	)

	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		sharedCounter int32
		taskQueue     = make(chan Task, maxTasks)
		rateLimiter   = time.Tick(time.Second / time.Duration(rateLimit))
	)

	// Start workers
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskQueue, &wg, &mu, &sharedCounter)
	}

	// Generate tasks
	for i := 1; i <= maxTasks; i++ {
		<-rateLimiter // Wait for rate limit
		taskQueue <- Task{ID: i, Task: func(id int) error {
			time.Sleep(time.Second * 1)
			if id%2 == 0 {
				return errors.New("an error ocurred")
			}
			return nil
		}}
		fmt.Printf("Submitted task %d\n", i)
	}

	close(taskQueue) // Close task queue to signal workers to exit
	wg.Wait()        // Wait for all workers to finish

	fmt.Printf("All tasks completed. Final counter value: %d\n", sharedCounter)
}
