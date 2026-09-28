package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Barrier struct {
	count  int
	wg     sync.WaitGroup
	mu     sync.Mutex
	notify chan struct{}
}

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(barrier *Barrier, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)

	barrier.mu.Lock()
	barrier.count--
	if barrier.count == 0 {
		close(barrier.notify)
	}
	barrier.mu.Unlock()
	<-barrier.notify
	//Rendezvous here

	fmt.Println("PartB", Num)
	barrier.wg.Done()
	return true
}

func main() {
	var barrier Barrier
	threadCount := 5
	barrier.count = threadCount
	barrier.notify = make(chan struct{})

	barrier.wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&barrier, N)
	}
	barrier.wg.Wait() //wait here until everyone (10 go routines) is done

}
