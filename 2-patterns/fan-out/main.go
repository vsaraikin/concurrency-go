package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateWork(job int) int {
	time.Sleep(time.Second)
	return job * 2
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result := simulateWork(job)
		fmt.Printf("Worker %d processed job %d resulting in %d\n", id, job, result)
		results <- result
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}
