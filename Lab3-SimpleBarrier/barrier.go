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
// Issues:
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	count  int
	wg     sync.WaitGroup
	mu     sync.Mutex
	notify chan struct{}
}

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, barrier *Barrier) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	barrier.mu.Lock()
	barrier.count--
	if barrier.count == 0 {
		close(barrier.notify)
	}
	barrier.mu.Unlock()

	<-barrier.notify

	//we wait here until everyone has completed part A
	fmt.Println("Part B", goNum)
	barrier.wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var barrier Barrier
	barrier.wg.Add(totalRoutines)
	barrier.count = totalRoutines
	barrier.notify = make(chan struct{})

	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &barrier)
	}

	barrier.wg.Wait() //wait for everyone to finish before exiting
}
