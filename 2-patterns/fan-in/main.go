package main

import (
	"fmt"
	"sync"
)

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(inputs))
	for _, input := range inputs {
		go output(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func simulateWorker(id int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			out <- id*10 + i
		}
		close(out)
	}()
	return out
}

func main() {
	worker1 := simulateWorker(1)
	worker2 := simulateWorker(2)

	consolidated := fanIn(worker1, worker2)

	for n := range consolidated {
		fmt.Println(n)
	}
}
