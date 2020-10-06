package concurrency

import (
	"sync/atomic"
)

type GoRoutineLimiter struct {
	maxGoRoutines     int32
	currentGoRoutines int32

	slotChan chan int32
}

func NewGoRoutineLimiter(max int32) *GoRoutineLimiter {
	slotChan := make(chan int32, max)
	for i := int32(0); i < max; i++ {
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

	atomic.AddInt32(&c.currentGoRoutines, 1)

	go func() {
		defer func() {
			c.slotChan <- slot

			atomic.AddInt32(&c.currentGoRoutines, -1)
		}()

		dispatcher()
	}()
}
