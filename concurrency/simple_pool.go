// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"errors"
)

type task func()

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(t task)
}

type WorkPool struct {
	tasks chan task
}

func (wp *WorkPool) Submit(t task) {
	// produce
	wp.tasks <- t
}
func (wp *WorkPool) Work() {
	// consume
	for task := range wp.tasks {
		task() // execute task
	}
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) (*WorkPool, error) {
	if maxConcurrent <= 0 {
		return nil, errors.New("maxConcurrent must be greater than zero")
	}
	workPool := &WorkPool{tasks: make(chan task)}

	for i := 0; i < maxConcurrent; i++ {
		go workPool.Work()
	}
	return workPool, nil

}
