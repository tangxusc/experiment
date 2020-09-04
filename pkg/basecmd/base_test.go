package basecmd

import (
	"testing"
	"time"
)

func TestNewCommand(t *testing.T) {
	command, _, cancelFunc := NewCommand()
	go func() {
		time.Sleep(time.Second * 10)
		cancelFunc()
	}()
	err := command.Execute()
	if err != nil {
		panic(err.Error())
	}
}
