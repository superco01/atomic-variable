package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counterAtomic int64
	var counterNonAtomic int64
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func() {
			atomic.AddInt64(&counterAtomic, 1)
			counterNonAtomic++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter atomic result: ", atomic.LoadInt64(&counterAtomic))
	fmt.Println("counter result: ", counterNonAtomic)
}
