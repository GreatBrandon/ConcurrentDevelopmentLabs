//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Barrier struct {
	count atomic.Int64
	n     int64
	//mu     sync.Mutex
	before chan int
	after  chan int
}

func New(n int64) *Barrier {
	b := Barrier{
		count:  atomic.Int64{},
		n:      n,
		before: make(chan int, 1),
		after:  make(chan int, 1),
	}
	// close 1st gate
	b.after <- 1
	return &b
}
func (b *Barrier) Before() {
	b.count.Add(1)
	if b.count.Load() == b.n {
		// close 2nd gate
		<-b.after
		// open 1st gate
		b.before <- 1
	}
	<-b.before
	b.before <- 1
}
func (b *Barrier) After() {
	b.count.Add(-1)
	if b.count.Load() == 0 {
		// close 1st gate
		<-b.before
		// open 2st gate
		b.after <- 1
	}
	<-b.after
	b.after <- 1
}

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, b *Barrier, loopCount int) bool {

	for i := 0; i < loopCount; i++ {
		fmt.Println("Part A: ", goNum)
		b.Before()
		fmt.Println("Part B: ", goNum)
		b.After()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 5
	loopCount := 5

	b := New(int64(totalRoutines))

	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, b, loopCount)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
