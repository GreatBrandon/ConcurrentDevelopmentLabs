package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(wg *sync.WaitGroup, buffer chan<- int) {
	defer wg.Done()
	for n := range 5 {
		buffer <- n
		fmt.Println("Produced ", n)
	}
	close(buffer)
}

func consume(wg *sync.WaitGroup, buffer <-chan int) {
	defer wg.Done()
	for {
		select {
		case x, ok := <-buffer:
			if !ok {
				return
			}
			fmt.Println("Consumed ", x)
			X := time.Duration(rand.Intn(2000))
			time.Sleep(X * time.Millisecond) //wait random time amount
		}
	}
}

func main() {
	buffer := make(chan int, 3)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go produce(wg, buffer)
	go consume(wg, buffer)
	wg.Wait()
}
