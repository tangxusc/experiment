package delayqueue

import (
	"context"
	"time"
)

type Task struct {
	Fn     TaskFunc
	Params []interface{}
}

type TaskFunc func(args ...interface{})

type DelayQueue interface {
	Start(ctx context.Context)
	AddTask(t time.Time, task Task) error
}
