package epoll

import (
	"fmt"
	"testing"
)

type Warr struct {
	s *string
}

func TestEpoll(t *testing.T) {
	s := "test"
	var w Warr
	w.s = &s
	fmt.Printf("%+v",*w.s)
	//net.Listen("tcp", ":8080")
}
