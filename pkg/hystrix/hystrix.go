package hystrix

import (
	"context"
	"fmt"
)

var CancelContextError = fmt.Errorf("cancel context")

type Run func() (interface{}, error)

/*
熔断器
*/
type Hystrix struct {
	limitOutChan <-chan struct{}               //<-ring
	executors    <-chan chan<- func()          //<-ring
	circuitIn    chan<- struct{}               // ->能写入
	circuitOut   <-chan error                  // <-能读出(快速失败)
	metrics      chan<- CommandExecutorMetrics //defer ->ring
}

//TODO:创建熔断器

func (h *Hystrix) Do(ctx context.Context, run Run) (result interface{}, err error) {
	do, errchan := h.GoDo(ctx, run)
	select {
	case s := <-do:
		return s, nil
	case e := <-errchan:
		return nil, e
	}
}

/**
如何返回呢?
*/
func (h *Hystrix) GoDo(ctx context.Context, run Run) (<-chan interface{}, <-chan error) {
	result := make(chan interface{}, 1)
	errChan := make(chan error, 1)
	//defer func() {
	//	close(result)
	//	close(errChan)
	//}()
	select {
	case <-ctx.Done():
		errChan <- CancelContextError
		return result, errChan
	case <-h.limitOutChan: //通过限流
		select {
		case <-ctx.Done():
			errChan <- CancelContextError
			return result, errChan
		case h.circuitIn <- struct{}{}: //未熔断
			select {
			case <-ctx.Done():
				errChan <- CancelContextError
				return result, errChan
			case executor := <-h.executors: //从池中获取执行器
				runFun := func() {
					runResult, err := run()
					if err != nil {
						errChan <- err
						return
					}
					result <- runResult
				}
				select {
				case <-ctx.Done():
					errChan <- CancelContextError
				case executor <- runFun: //执行run
				}
			}
		case out := <-h.circuitOut: //如果熔断,则从此读取熔断
			errChan <- out
			return result, errChan
		}
	}
	//统计
	go func() {
		h.metrics <- CommandExecutorMetrics{error: len(errChan) > 0}
	}()
	return result, errChan
}

