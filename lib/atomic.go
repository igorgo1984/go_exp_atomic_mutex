package lib

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	val int64
}

func (i *atomicCounter) Add() {
	atomic.AddInt64(&i.val, 1)
	// Without be fast. But need run atomic iteration
	runtime.Gosched()
}

func (i *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&i.val)
}

func CountAtomic() int64 {
	wg := sync.WaitGroup{}
	wg.Add(CountGoroutine)

	counter := atomicCounter{}

	for i := 0; i < CountGoroutine; i++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < CountInc; i++ {
				counter.Add()
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	return counter.Value()
}
