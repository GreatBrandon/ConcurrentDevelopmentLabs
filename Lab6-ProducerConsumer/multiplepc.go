// Author: Brandon Jaroszczak
// Student number: C00296052

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(num, count int, wg *sync.WaitGroup, buffer chan<- int) {
	defer wg.Done()
	for n := range count {
		buffer <- n + 100
		fmt.Println("Producer ", num, " produced: ", n+100)
	}
}

func consume(n int, wg *sync.WaitGroup, buffer <-chan int) {
	defer wg.Done()
	for {
		select {
		case x, ok := <-buffer:
			if !ok {
				return
			}
			fmt.Println("Consumer ", n, " consumed: ", x)
			// simulate work
			X := time.Duration(rand.Intn(2000))
			time.Sleep(X * time.Millisecond)
		}
	}
}

func main() {
	numProducers := 1
	numConsumers := 10
	producerLoopCount := 50
	buffer := make(chan int, 5)
	pwg := new(sync.WaitGroup)
	cwg := new(sync.WaitGroup)
	pwg.Add(numProducers)
	cwg.Add(numConsumers)

	for i := range numProducers {
		go produce(i+1, producerLoopCount, pwg, buffer)
	}
	for i := range numConsumers {
		go consume(i+1, cwg, buffer)
	}
	pwg.Wait() // wait until all producers are done before closing
	close(buffer)
	cwg.Wait() // wait until all consumers are done before exiting
}
