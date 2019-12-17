package lib

import "sync"

type mutexCounter struct {
	sync.RWMutex
	val int64
}

func (i *mutexCounter) Add() {
	i.Lock()
	i.val += 1
	i.Unlock()
}

func (i *mutexCounter) Value() int64 {
	i.RLock()
	v := i.val
	i.RUnlock()

	return v
}

func CountMutex() int64 {
	wg := sync.WaitGroup{}
	wg.Add(CountGorotine)

	counter := mutexCounter{}

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
