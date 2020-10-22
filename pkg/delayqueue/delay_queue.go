package delayqueue

import (
	"context"
	"errors"
	"time"
)

const max = 3600

type DelayQueue struct {
	curIndex  int
	slots     [max]map[string]*Task
	time2task chan int
}

type TaskFunc func(args ...interface{})

type Task struct {
	cycleNum int
	exec     TaskFunc
	params   []interface{}
}

func NewDelayQueue() *DelayQueue {
	dm := &DelayQueue{
		curIndex:  0,
		time2task: make(chan int),
	}

	for i := 0; i < max; i++ {
		dm.slots[i] = make(map[string]*Task)
	}

	return dm
}

func (dm *DelayQueue) taskLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case index := <-dm.time2task:
			tasks := dm.slots[index]
			if len(tasks) <= 0 {
				continue
			}
			for k, v := range tasks {
				if v.cycleNum == 0 {
					go v.exec(v.params...)
					delete(tasks, k)
				} else {
					v.cycleNum--
				}
			}
		}
	}
}

//启动延迟消息
func (dm *DelayQueue) Start(ctx context.Context) {
	go dm.taskLoop(ctx)
	go dm.timeLoop(ctx)
	select {
	case <-ctx.Done():
		return
	}
}

//处理每1秒移动下标
func (dm *DelayQueue) timeLoop(ctx context.Context) {
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
func (dm *DelayQueue) AddTask(t time.Time, key string, exec TaskFunc, params []interface{}) error {
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
	tasks := dm.slots[ix]
	if _, ok := tasks[key]; ok {
		return errors.New("该slots中已存在key为" + key + "的任务")
	}
	tasks[key] = &Task{
		cycleNum: cycleNum,
		exec:     exec,
		params:   params,
	}
	return nil
}
