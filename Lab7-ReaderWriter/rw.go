// Author: Brandon Jaroszczak
// Student number: C00296052

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Buffer struct {
	data int
	quit bool
	mu   sync.RWMutex
}

func reader(n int, wg *sync.WaitGroup, b *Buffer) {
	defer wg.Done()
	num := 0
	quit := false

	for {
		X := time.Duration(rand.Intn(500))
		time.Sleep(X * time.Millisecond)
		b.mu.RLock()
		num = b.data
		quit = b.quit
		b.mu.RUnlock()
		if quit {
			fmt.Println("Reader ", n, "read ", num, "and quit")
			return // stop reading after writer stopped writing
		}
		fmt.Println("Reader ", n, "read ", num)
	}
}

func writer(wg *sync.WaitGroup, b *Buffer) {
	defer wg.Done()

	for i := range 10 {
		b.mu.Lock()
		b.data = i
		b.mu.Unlock()
		fmt.Println("Writer wrote:", i)
		X := time.Duration(rand.Intn(2000))
		time.Sleep(X * time.Millisecond)
	}
	b.mu.Lock()
	b.quit = true
	b.mu.Unlock()
	fmt.Println("Writer signalled quit")
}

func main() {
	readerCount := 5
	buffer := new(Buffer)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go writer(wg, buffer)

	wg.Add(readerCount)
	for i := range readerCount {
		go reader(i+1, wg, buffer)
	}
	wg.Wait()
}
