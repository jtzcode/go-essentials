package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func otherTask() {
	fmt.Println("working on sth else...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Other task is done.")
}
func TestAysncService(t *testing.T) {
	channel := AsyncService()
	otherTask()
	fmt.Println(<-channel)
}
