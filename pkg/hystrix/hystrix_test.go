package hystrix

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestHystrix_GoDo(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	todo := context.TODO()
	limiter := NewSimpleLimiter()
	limiter.run(todo)
	executors := NewExecutors()
	executors.Run(todo)
	circuit := NewSimpleCircuit()
	metrics := NewSimpleMetrics(circuit)
	metrics.Run(todo)
	circuit.Run(todo)

	time.Sleep(2*time.Second)

	hystrix := &Hystrix{
		executors:    executors.GetChan(),
		limitOutChan: limiter.GetOut(),
		circuitIn:    circuit.in,
		circuitOut:   circuit.out,
		metrics:      metrics.metricsIn,
	}
	result, err := hystrix.Do(context.TODO(), func() (interface{}, error) {
		time.Sleep(time.Second * 5)
		return 1, nil
	})
	fmt.Println(result, err)

	time.Sleep(time.Second * 20)
}
