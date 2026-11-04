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

func writer(n, offset, loopCount int, wg *sync.WaitGroup, b *Buffer) {
	defer wg.Done()

	for i := range loopCount {
		b.mu.Lock()
		b.data = i + offset
		b.mu.Unlock()
		fmt.Println("Writer ", n, " wrote:", i+offset)
		X := time.Duration(rand.Intn(2000))
		time.Sleep(X * time.Millisecond)
	}
}

func main() {
	readerCount := 5
	writerCount := 3
	loopCount := 10

	buffer := new(Buffer)
	rwg := new(sync.WaitGroup)
	wwg := new(sync.WaitGroup)

	wwg.Add(writerCount)
	for i := range writerCount {
		go writer(i+1, i*loopCount, loopCount, wwg, buffer)
	}

	rwg.Add(readerCount)
	for i := range readerCount {
		go reader(i+1, rwg, buffer)
	}

	// after all writers are done, signal quit to readers (to end program & prevent infinite reading)
	wwg.Wait()
	buffer.mu.Lock()
	buffer.quit = true
	buffer.mu.Unlock()
	rwg.Wait()
}
