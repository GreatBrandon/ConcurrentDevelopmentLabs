package main

import (
	"fmt"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func fib(N int) int {
	if N < 2 {
		return 1
	}

	return fib(N-1) + fib(N-2)
}

type FibTable struct {
	mu     sync.Mutex
	values map[int]int
}

func (t *FibTable) Get(n int) (int, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	value, ok := t.values[n]
	return value, ok
}

func (t *FibTable) Set(n int, value int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.values[n] = value
}

func parFib(N int, table *FibTable) int {
	var wg sync.WaitGroup
	var A, B int

	if N < 2 {
		return 1
	}

	if value, ok := table.Get(N); ok {
		return value
	}

	wg.Add(2)

	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N-1, table)
	}(N, &A)
	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N-2, table)
	}(N, &B)
	wg.Wait()

	result := A + B
	table.Set(N, result)
	return result
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	for i := 0; i < 10; i++ {
		Seq := fib(i * 5)
		var table FibTable
		table.values = make(map[int]int)
		par := parFib(i*5, &table)
		fmt.Println(Seq, "---", par)
	}
}
