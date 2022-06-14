package multichannel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

//1m24.2352155s
func TestMultiChannel(t *testing.T) {
	c1 := make(chan interface{}, 1000)
	c2 := make(chan interface{}, 1000)
	sig := make(chan interface{}, 1000000)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-c1:
				c2 <- i
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-c2:
				fmt.Println(i)
				sig <- i
			}
		}
	}(ctx)
	time.Sleep(time.Second)

	now := time.Now()
	for i := 0; i < 1000000; i++ {
		c1 <- i
		<-sig
	}
	sub := time.Now().Sub(now)
	fmt.Printf("%v \n", sub)
	cancelFunc()
}

//1m27.2415094s
func TestMultiLock(t *testing.T) {
	mutex1 := sync.Mutex{}
	mutex2 := sync.Mutex{}
	sig := make(chan interface{}, 1000000)

	now := time.Now()
	for i := 0; i < 1000000; i++ {
		go func(i int) {
			mutex1.Lock()
			mutex2.Lock()
			fmt.Println(i)
			mutex1.Unlock()
			mutex2.Unlock()
			sig <- i
		}(i)
		<-sig
	}
	sub := time.Now().Sub(now)
	fmt.Printf("%v \n", sub)
}

//1m23.05649s
func TestMultiLock2(t *testing.T) {
	mutex1 := sync.Mutex{}
	mutex2 := sync.Mutex{}
	//sig := make(chan interface{}, 1000000)

	now := time.Now()
	for i := 0; i < 1000000; i++ {
		//go func(i int) {
		mutex1.Lock()
		mutex2.Lock()
		fmt.Println(i)
		mutex1.Unlock()
		mutex2.Unlock()
		//sig <- i
		//}(i)
		//<-sig
	}
	sub := time.Now().Sub(now)
	fmt.Printf("%v \n", sub)
}
