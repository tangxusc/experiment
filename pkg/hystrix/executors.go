package hystrix

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type ExecutorsI interface {
	GetChan() <-chan chan<- func()
}

type ChanExecutors struct {
	ch chan chan<- func()
}

func (e *ChanExecutors) GetChan() <-chan chan<- func() {
	return e.ch
}

func (e *ChanExecutors) Run(ctx context.Context) {
	e.ch = make(chan chan<- func(), 100)
	for i := 0; i < cap(e.ch); i++ {
		go NewExecutor(e.ch, fmt.Sprintf("executor-%v", i)).run(ctx)
	}
}

func NewExecutors() *ChanExecutors {
	e := &ChanExecutors{}
	return e
}

type ChanExecutor struct {
	name       string
	globalChan chan chan<- func()
	ch         chan func()
}

func NewExecutor(globalChan chan chan<- func(), name string) *ChanExecutor {
	return &ChanExecutor{
		globalChan: globalChan,
		ch:         make(chan func()),
		name:       name,
	}
}

//自循环
func (e *ChanExecutor) run(ctx context.Context) {
	for {
		select {
		case e.globalChan <- e.ch:
			logrus.Debugf("[%v]将自身放入全局执行器中,全局执行器现有%v", e.name, len(e.globalChan))
		case <-ctx.Done():
			return
		case f := <-e.ch:
			//fmt.Printf()
			logrus.Debugf("[%v]收到执行任务[%T],准备执行", e.name, f)
			recoverInvoke(f, e.name)
			logrus.Debugf("[%v]执行任务[%T]完成", e.name, f)
		}
	}
}

func recoverInvoke(f func(), name string) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("[%v]invoke func(%T) error:%v", name, f, err)
		}
	}()
	f()
}
