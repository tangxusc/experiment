package eventbus

import (
	"fmt"
	"github.com/vibrantbyte/go-antpath/antpath"
	"reflect"
)

type Event interface {
	GetKey() string
	GetData() interface{}
}

type StandardEvent struct {
	Key  string
	Data interface{}
}

func (s *StandardEvent) GetKey() string {
	return s.Key
}

func (s *StandardEvent) GetData() interface{} {
	return s.Data
}

type Consumer func(event Event) error

type Exchange interface {
	GetMatchKeyRouter(key string) ([]Consumer, error)
	AddConsumer(path string, consumer Consumer)
	DeleteConsumer(path string, consumer Consumer)
}

type EventBus struct {
	exchange Exchange
}

type AntPathExchange struct {
	matcher antpath.PathMatcher
	routers map[string][]Consumer
}

func (a *AntPathExchange) GetMatchKeyRouter(key string) ([]Consumer, error) {
	result := make([]Consumer, 0)
	for path, consumers := range a.routers {
		match := a.matcher.Match(path, key)
		if match {
			result = append(result, consumers...)
		}
	}
	return result, nil
}

func (a *AntPathExchange) AddConsumer(path string, consumer Consumer) {
	consumers, ok := a.routers[path]
	if !ok {
		consumers = make([]Consumer, 0)
	}
	consumers = append(consumers, consumer)
	a.routers[path] = consumers
}

func (a *AntPathExchange) DeleteConsumer(path string, consumer Consumer) {
	consumers, ok := a.routers[path]
	if !ok {
		return
	}
	for i, c := range consumers {
		if reflect.DeepEqual(c, consumer) {
			consumers = append(consumers[0:i], consumers[i+1:]...)
		}
	}
	a.routers[path] = consumers
}

var DefaultExchange = &AntPathExchange{
	matcher: antpath.New(),
	routers: make(map[string][]Consumer),
}

func NewEventBus() *EventBus {
	return &EventBus{
		exchange: DefaultExchange,
	}
}

func (e *EventBus) Register(path string, consumer Consumer) {
	e.exchange.AddConsumer(path, consumer)
}

func (e *EventBus) UnRegister(path string, consumer Consumer) {
	e.exchange.DeleteConsumer(path, consumer)
}

func (e *EventBus) SendEvent(event Event) error {
	consumers, err := e.exchange.GetMatchKeyRouter(event.GetKey())
	if err != nil {
		return err
	}
	for _, consumer := range consumers {
		go call(consumer, event)
	}
	return nil
}

func (e *EventBus) Send(key string, data interface{}) error {
	ev := &StandardEvent{
		Key:  key,
		Data: data,
	}
	return e.SendEvent(ev)
}

func call(consumer Consumer, event Event) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	err := consumer(event)
	if err != nil {
		fmt.Println(err.Error())
	}
}
