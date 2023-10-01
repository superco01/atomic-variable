package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

const iteration = 100

func BenchmarkAtomic(t *testing.B) {
	var wg sync.WaitGroup
	var counter int64

	wg.Add(iteration)

	for i := 0; i < iteration; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iteration; i++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
}

func BenchmarkNonAtomic(t *testing.B) {
	var wg sync.WaitGroup
	var counter int64

	wg.Add(iteration)

	for i := 0; i < iteration; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iteration; i++ {
				counter++
			}
		}()
	}
}
