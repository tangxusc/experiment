package levelchan

import (
	"context"
	"sort"
)

type LevelChansItem struct {
	ch    <-chan interface{}
	level uint
}

type LevelChans struct {
	chans []*LevelChansItem
}

func (c *LevelChans) Len() int {
	return len(c.chans)
}

func (c *LevelChans) Less(i, j int) bool {
	return c.chans[i].level < c.chans[j].level
}

func (c *LevelChans) Swap(i, j int) {
	c.chans[i], c.chans[j] = c.chans[j], c.chans[i]
}

func (c *LevelChans) Append(level uint, ch <-chan interface{}) {
	c.chans = append(c.chans, &LevelChansItem{
		ch:    ch,
		level: level,
	})
	sort.Sort(c)
}

func (c *LevelChans) Read(todo context.Context) <-chan interface{} {
	out := make(chan interface{})
	if len(c.chans) == 0 {
		close(out)
		return out
	}
	go func() {
		var current <-chan interface{}
		index := 0
		for {
			if len(c.chans) <= index {
				index = 0
			}
			current = c.chans[index].ch
			select {
			case <-todo.Done():
				close(out)
				return
			default:
				select {
				case e := <-current:
					out <- e
					index = 0
				default:
					index++
				}
			}
		}
	}()

	return out
}

func NewChans() *LevelChans {
	return &LevelChans{}
}
