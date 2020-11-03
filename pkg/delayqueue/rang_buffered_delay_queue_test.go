package delayqueue

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewDelayQueue(t *testing.T) {
	//创建延迟消息
	dm := NewRangeBufferedDelayQueue()
	//添加任务
	err := dm.AddTask(time.Now().Add(time.Second*4), Task{
		Fn: func(args ...interface{}) {
			fmt.Println(args...)
		},
		Params: []interface{}{1, 2, 3},
	})
	fmt.Println(err)
	err = dm.AddTask(time.Now().Add(time.Second*7), Task{
		Fn: func(args ...interface{}) {
			fmt.Println(args...)
		},
		Params: []interface{}{4, 5, 6},
	})
	fmt.Println(err)

	timeout, _ := context.WithTimeout(context.TODO(), time.Second*13)
	dm.Start(timeout)
}
