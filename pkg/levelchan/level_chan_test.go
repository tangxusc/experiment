package levelchan

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLevelChan(t *testing.T) {
	c1 := make(chan interface{}, 10)
	c2 := make(chan interface{}, 10)
	c3 := make(chan interface{}, 10)
	go func(ch ...chan interface{}) {
		for i, c := range ch {
			for j := 0; j < 5; j++ {
				c <- fmt.Sprintf("%v,正常的写入", i)
			}
		}
	}(c1, c2, c3)
	time.Sleep(time.Second)
	fmt.Println("chan len:", len(c1), len(c2), len(c3))

	chans := NewChans()
	chans.Append(1, c1)
	chans.Append(2, c2)
	chans.Append(3, c3)

	cancel, cancelFunc := context.WithCancel(context.TODO())
	ch := chans.Read(cancel)

	go func(ch ...chan interface{}) {
		time.Sleep(5 * time.Second)
		for _, c := range ch {
			for j := 0; j < 5; j++ {
				c <- fmt.Sprintf("%v,延迟10s的写入", "c2")
			}
		}
	}(c2)

	go func(ch ...chan interface{}) {
		time.Sleep(5 * time.Second)
		for _, c := range ch {
			for j := 0; j < 5; j++ {
				c <- fmt.Sprintf("%v,延迟10s的写入", "c3")
			}
		}
	}(c3)

	go func(ch ...chan interface{}) {
		time.Sleep(5 * time.Second)
		for _, c := range ch {
			for j := 0; j < 5; j++ {
				c <- fmt.Sprintf("%v,延迟10s的写入", "c1")
			}
		}
	}(c1)

	go func() {
		time.Sleep(10 * time.Second)
		cancelFunc()
	}()

	//先读取之前的
	for i := 0; i < 15; i++ {
		e := <-ch
		fmt.Println(e)
	}
	time.Sleep(time.Second * 7)
	for i := range ch {
		fmt.Println(i)
	}
}

