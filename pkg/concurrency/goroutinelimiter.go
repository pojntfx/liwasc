package concurrency

import (
	"sync/atomic"
)

type GoRoutineLimiter struct {
	maxGoRoutines     int64
	currentGoRoutines int64

	slotChan chan int64
}

func NewGoRoutineLimiter(max int64) *GoRoutineLimiter {
	slotChan := make(chan int64, max)
	for i := int64(0); i < max; i++ {
		slotChan <- i
	}

	return &GoRoutineLimiter{
		maxGoRoutines:     max,
		currentGoRoutines: 0,

		slotChan: slotChan,
	}
}

func (c *GoRoutineLimiter) Dispatch(dispatcher func()) {
	slot := <-c.slotChan

	atomic.AddInt64(&c.currentGoRoutines, 1)

	go func() {
		defer func() {
			c.slotChan <- slot

			atomic.AddInt64(&c.currentGoRoutines, -1)
		}()

		dispatcher()
	}()
}
