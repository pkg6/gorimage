package gorimage

import (
	"context"
	"math"
	"sync"
)

type WaitGroup struct {
	Size int

	current chan struct{}
	wg      sync.WaitGroup
}

func NewWaitGroup(limit int) WaitGroup {
	size := math.MaxInt32 // 2^31 - 1
	if limit > 0 {
		size = limit
	}
	return WaitGroup{
		Size:    size,
		current: make(chan struct{}, size),
		wg:      sync.WaitGroup{},
	}
}
func (s *WaitGroup) Add() {
	_ = s.AddWithContext(context.Background())
}
func (s *WaitGroup) AddWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.current <- struct{}{}:
		break
	}
	s.wg.Add(1)
	return nil
}
func (s *WaitGroup) Done() {
	<-s.current
	s.wg.Done()
}
func (s *WaitGroup) Wait() {
	s.wg.Wait()
}
