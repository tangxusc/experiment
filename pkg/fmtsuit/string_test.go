package fmtsuit

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestStruct(t *testing.T) {
	strcut := &TestStrcut{Name: "sss"}
	fmt.Printf("%v \n", strcut)
	fmt.Println(strcut)
}

func TestCommand(t *testing.T) {
	c := exec.Command("cmd", "/C", "ipconfig")
	//c := exec.Command("cmd", "/C", "ping", "www.baidu.com")
	_, w, err := os.Pipe()
	if err != nil {
		panic(err.Error())
	}
	c.Stderr = w
	c.Stderr = w
	c.Stdin = os.Stdin
	//go func() {
	//	fmt.Println("start read ")
	//	all, err2 := ioutil.ReadAll(r)
	//	if err2 != nil {
	//		panic(err2.Error())
	//	}
	//	fmt.Println(all)
	//	fmt.Println("end read ")
	//
	//}()
	err = c.Run()
	if err != nil {
		print(err.Error())
	}
	//err = c.Wait()
	//if err != nil {
	//	panic(err.Error())
	//}

	time.Sleep(time.Hour)
}
