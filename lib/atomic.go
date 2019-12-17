package lib

import (
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	val int64
}

func (i *atomicCounter) Add() {
	atomic.AddInt64(&i.val, 1)
}

func (i *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&i.val)
}

func CountAtomic() int64 {
	wg := sync.WaitGroup{}
	wg.Add(CountGorotine)

	counter := atomicCounter{}

	for i := 0; i < CountGorotine; i++ {
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
