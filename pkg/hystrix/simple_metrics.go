package hystrix

import (
	"context"
)

type CommandExecutorMetrics struct {
	error bool
}

type SimpleMetrics struct {
	metricsIn chan CommandExecutorMetrics
	circuit   *SimpleCircuit
}

func NewSimpleMetrics(circuit *SimpleCircuit) *SimpleMetrics {
	return &SimpleMetrics{
		metricsIn: make(chan CommandExecutorMetrics),
		circuit:   circuit,
	}
}

func (m *SimpleMetrics) Run(ctx context.Context) {
	go func() {
		//timer := time.NewTimer(time.Second * 2)
		//var cancel context.Context
		//var c context.CancelFunc
		//run := false
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-m.metricsIn:
				if !e.error {
					continue
				}

				//TODO:通过计算各种指标后进行熔断

				m.circuit.Circuit(ctx, e)
				//if !run {
				//	cancel, c = context.WithCancel(ctx)
				//	go m.StopErrorChanWithCancel(timer, cancel, c, func() {
				//		run = false
				//	})
				//	//开启熔断
				//	go m.StartErrorChan(cancel)
				//	run = true
				//}
				//if !timer.Stop() {
				//	select {
				//	case <-timer.C:
				//	default:
				//	}
				//}
				//timer.Reset(time.Second * 2)
			}
		}
	}()
}

//func (m *SimpleMetrics) StopErrorChanWithCancel(timer *time.Timer, cancel context.Context, c context.CancelFunc, call func()) {
//	select {
//	case <-cancel.Done():
//	case <-timer.C:
//		//恢复未熔断状态
//		c()
//		call()
//	}
//}
//
//func (m *SimpleMetrics) StartErrorChan(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		default:
//			select {
//			case <-ctx.Done():
//			default:
//				//out
//				m.metricsOut <- CommandExecutorMetrics{}
//			}
//		}
//	}
//}
