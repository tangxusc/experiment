package delayqueue

import (
	"context"
	"errors"
	"time"
)

const max = 3600

type RangeBufferedDelayQueue struct {
	curIndex  int
	slots     [max][]*RangeBufferedTask
	time2task chan int
}

type RangeBufferedTask struct {
	*Task
	cycleNum int
}

func NewRangeBufferedDelayQueue() *RangeBufferedDelayQueue {
	dm := &RangeBufferedDelayQueue{
		curIndex:  0,
		time2task: make(chan int),
	}

	for i := 0; i < max; i++ {
		dm.slots[i] = make([]*RangeBufferedTask, 0)
	}

	return dm
}

func (dm *RangeBufferedDelayQueue) taskLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case index := <-dm.time2task:
			tasks := dm.slots[index]
			if len(tasks) <= 0 {
				continue
			}
			tmpTasks := make([]*RangeBufferedTask, 0, len(tasks))
			for k, v := range tasks {
				if v.cycleNum == 0 {
					go v.Fn(v.Params...)
				} else {
					v.cycleNum--
					tmpTasks = append(tmpTasks, tasks[k])

				}
			}
			//fmt.Println("当前:", len(tmpTasks))
			dm.slots[index] = tmpTasks
		}
	}
}

//启动延迟消息
func (dm *RangeBufferedDelayQueue) Start(ctx context.Context) {
	go dm.taskLoop(ctx)
	go dm.timeLoop(ctx)
	select {
	case <-ctx.Done():
		return
	}
}

//处理每1秒移动下标
func (dm *RangeBufferedDelayQueue) timeLoop(ctx context.Context) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			//判断当前下标，如果等于3599则重置为0，否则加1
			if dm.curIndex == max-1 {
				dm.curIndex = 0
			} else {
				dm.curIndex++
			}
			dm.time2task <- dm.curIndex
		}
	}
}

//添加任务
func (dm *RangeBufferedDelayQueue) AddTask(t time.Time, task Task) error {
	now := time.Now()
	if now.After(t) {
		return errors.New("目标时间不能在当前时间之前")
	}
	//当前时间与指定时间相差秒数
	subSecond := t.Unix() - now.Unix()
	//计算循环次数
	cycleNum := int(subSecond / max)
	//计算任务所在的slots的下标
	ix := subSecond % max
	//把任务加入tasks中
	dm.slots[ix] = append(dm.slots[ix], &RangeBufferedTask{
		Task:     &task,
		cycleNum: cycleNum,
	})
	return nil
}
