package delayqueue

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewDelayQueue(t *testing.T) {
	//创建延迟消息
	dm := NewDelayQueue()
	//添加任务
	err := dm.AddTask(time.Now().Add(time.Second*10), "test1", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{1, 2, 3})
	fmt.Println(err)
	err = dm.AddTask(time.Now().Add(time.Second*12), "test2", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{4, 5, 6})
	fmt.Println(err)

	timeout, _ := context.WithTimeout(context.TODO(), time.Second*13)
	dm.Start(timeout)
}
