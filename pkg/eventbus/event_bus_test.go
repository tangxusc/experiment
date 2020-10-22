package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func TestNewEventBus(t *testing.T) {
	bus := NewEventBus()
	bus.Register("/**/1", func(event Event) error {
		fmt.Printf("/**/1 print event:%v \n", event)
		return nil
	})
	bus.Register("/agg/1", func(event Event) error {
		fmt.Printf("/agg/1 print event:%v \n", event)
		return nil
	})
	bus.Register("/agg/2", func(event Event) error {
		fmt.Printf("/agg/2 print event:%v \n", event)
		return nil
	})
	err := bus.Send("/agg/1", "abcd")
	fmt.Printf("send event,result:%v \n", err)

	time.Sleep(time.Hour)
}
