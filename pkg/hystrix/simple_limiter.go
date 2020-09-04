package hystrix

import (
	"context"
	"time"
)

type SimpleLimiter struct {
	Out chan struct{}
}

func (l *SimpleLimiter) GetOut() <-chan struct{} {
	return l.Out
}

func (l *SimpleLimiter) run(ctx context.Context) {
	go func() {
		for {
			select {
			case l.Out <- struct{}{}:
			case <-ctx.Done():
				return
			}
			//todo:应该有更好的方法
			//考虑统一的时钟,通过时钟计数方式
			time.Sleep(time.Second)
		}
	}()
}

func NewSimpleLimiter() *SimpleLimiter {
	s := &SimpleLimiter{
		Out: make(chan struct{}),
	}
	return s
}
