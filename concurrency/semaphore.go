// Package concurrency provides concurrency control utilities
package concurrency

import (
	"sync"
)

// Semaphore provides a counting semaphore implementation
type Semaphore struct {
	c  chan struct{}
	wg sync.WaitGroup
}

// NewSemaphore returns a new Semaphore initialized to the given value.
func NewSemaphore(maxCount int) *Semaphore {
	return &Semaphore{c: make(chan struct{}, maxCount)}
}

// Acquire acquires a permit, blocking until it becomes available or ctx is done.
func (s *Semaphore) Acquire(delta int) {
	s.wg.Add(delta)
	for i := 0; i < delta; i++ {
		s.c <- struct{}{}
	}
}

// Release releases a permit.
func (s *Semaphore) Release() {
	<-s.c
	s.wg.Done()
}

// Wait blocks until all permits have been released.
func (s *Semaphore) Wait() {
	s.wg.Wait()
}

// AcquireWithFunc gets the semaphore and executes the callback function with arguments
func (s *Semaphore) AcquireWithFunc(f func(args ...interface{}), args ...interface{}) {
	go func() {
		defer s.Release()
		s.Acquire(1)
		f(args...)
	}()
}
