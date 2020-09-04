package hystrix

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var SimpleCircuitError = fmt.Errorf("已熔断")

type SimpleCircuit struct {
	in  chan struct{}
	out chan error

	interval time.Duration
	cancel   context.Context
	c        context.CancelFunc
	run      bool
	timer    *time.Timer

	lock int32 //0 unlock 1 lock
}

func NewSimpleCircuit() *SimpleCircuit {
	return &SimpleCircuit{
		in:       make(chan struct{}),
		out:      make(chan error),
		interval: time.Second * 2,
	}
}

func (s *SimpleCircuit) Run(ctx context.Context) {
	//read in
	go func() {
		for {
			if loadInt32 := atomic.LoadInt32(&s.lock); loadInt32 == 1 {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case <-s.in:

			}
		}
	}()

	//write out
	//go func() {
	//	for {
	//		if !s.err {
	//			continue
	//		}
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case s.out <- SimpleCircuitError:
	//
	//		}
	//	}
	//}()

}

func (s *SimpleCircuit) Circuit(ctx context.Context, e CommandExecutorMetrics) {
	//todo:run变量在goroutine同步
	if loadInt32 := atomic.LoadInt32(&s.lock); loadInt32 == 0 {
		s.timer = time.NewTimer(s.interval)
		s.cancel, s.c = context.WithCancel(ctx)
		s.startCircuit(ctx)
		s.run = true
	}
	if !s.timer.Stop() {
		select {
		case <-s.timer.C:
		default:
		}
	}
	s.timer.Reset(time.Second * 2)

}

func (s *SimpleCircuit) startCircuit(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				goto ex
			case <-s.timer.C:
				goto ex
			default:
				select {
				case s.out <- SimpleCircuitError:
				default:
					runtime.Gosched()
				}
			}
		}
	ex:
		s.timer.Stop()
		swaped := false
		for !swaped {
			swaped = atomic.CompareAndSwapInt32(&s.lock, 1, 0)
		}
		return
	}()

}
