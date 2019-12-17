package lib

import "sync"

type mutexCounter struct {
	// Simple mutex be fasted
	//sync.RWMutex
	sync.Mutex
	val int64
}

func (i *mutexCounter) Add() {
	i.Lock()
	i.val += 1
	i.Unlock()
}

func (i *mutexCounter) Value() int64 {
	//i.RLock()
	i.Lock()
	v := i.val
	i.Unlock()
	//i.RUnlock()

	return v
}

func CountMutex() int64 {
	wg := sync.WaitGroup{}
	wg.Add(CountGoroutine)

	counter := mutexCounter{}

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
