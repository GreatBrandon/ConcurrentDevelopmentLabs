// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
//GPL Licence
// MISSING:
// 1. Readme
// 2. Full licence info.
// 3. Comments
// 4. It can Deadlock!

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(2))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(2))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	if index == 4 { // 1 philosopher chooses differently
		forks[(index+1)%5] <- true
		forks[index] <- true
	} else {
		forks[index] <- true
		forks[(index+1)%5] <- true
	}
}

func putForks(index int, forks map[int]chan bool) {
	if index == 4 { // 1 philosopher chooses differently, prevents deadlock
		<-forks[index]
		<-forks[(index+1)%5]
	} else {
		<-forks[(index+1)%5]
		<-forks[index]
	}
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers
	wg.Wait() //wait here until everyone is done

} //main
