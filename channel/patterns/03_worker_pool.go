//go:build ignore

// Worker pool: a fixed number of goroutines take jobs from one channel.
// Each job is taken by exactly one worker, and at most 3 jobs run at the same time.
// It limits how much runs at once (DB connections, API rate limits, CPU cores).
//
// One close(jobs) ends the range in every worker.
//
// Run: go run ./channel/patterns/03_worker_pool.go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int) {
	for j := range jobs { // the workers compete for jobs on the same channel
		// do the job, e.g. send an email
		fmt.Printf("worker %d did job %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int)

	var wg sync.WaitGroup
	for id := range 3 { // 3 workers
		wg.Add(1)
		go func() {
			worker(id, jobs)
			wg.Done()
		}()
	}

	for j := range 9 {
		jobs <- j
	}
	close(jobs) // no more jobs: every worker's range ends

	wg.Wait() // wait for the workers to finish their last jobs
}
