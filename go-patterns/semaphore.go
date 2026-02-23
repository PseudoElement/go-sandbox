package main

import (
	"context"
)

type Semaphore struct {
	slots chan struct{}
	ctx   context.Context
}

func NewSemaphore(ctx context.Context, limit int) *Semaphore {
	return &Semaphore{
		slots: make(chan struct{}, limit),
		ctx:   ctx,
	}
}

func (s *Semaphore) Acquire() error {
	select {
	case s.slots <- struct{}{}:
		return nil
	case <-s.ctx.Done():
		return s.ctx.Err()
	}
}

func (s *Semaphore) Release() {
	<-s.slots
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.slots <- struct{}{}:
		return true
	default:
		return false
	}
}
