package lib

import (
	"runtime"
	"sync"
	"sync/atomic"
)

const (
	unlocked int32 = iota
	locked
)

// A atomicMutex must not be copied after first use.
type atomicMutex struct {
	state int32
}

func (lock *atomicMutex) Lock() {
	for atomic.LoadInt32(&lock.state) == locked{
		runtime.Gosched()
	}

	for !atomic.CompareAndSwapInt32(&lock.state, unlocked, locked) {
		runtime.Gosched()
	}
}

func (lock *atomicMutex) Unlock() {
	for atomic.LoadInt32(&lock.state) == unlocked {
		runtime.Gosched()
	}

	for !atomic.CompareAndSwapInt32(&lock.state, locked, unlocked) {
		runtime.Gosched()
	}
}

type atomicMutexCounter struct {
	val int64
	lock atomicMutex
}

func (i *atomicMutexCounter) Add() {
	i.lock.Lock()
	i.val += 1
	i.lock.Unlock()
}

func (i *atomicMutexCounter) Value() int64 {
	i.lock.Lock()
	v := i.val
	i.lock.Unlock()

	return v
}

func CountAtomicMutex() int64 {
	wg := sync.WaitGroup{}
	wg.Add(CountGoroutine)

	counter := atomicMutexCounter{}

	for i := 0; i < CountGoroutine; i++ {
		go func(wg *sync.WaitGroup, n int) {
			for i := 0; i < CountInc; i++ {
				counter.Add()
			}
			wg.Done()
		}(&wg, i)
	}

	wg.Wait()

	return counter.Value()
}
