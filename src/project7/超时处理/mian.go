package main

import (
	"fmt"
	"time"
)

func main() {
	// 超时处理，定义一个操作的处理时间，如果超过了指定时间，则走超时后定义的逻辑

	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "我是你爸爸我真伟大，养你这么大"
	}()

	select {
	// 如果未超时，则会打印res
	case res := <-c1:
		fmt.Println(res)
	// 超时则会执行超时后的处理
	case <-time.After(time.Second * 1):
		fmt.Println("哦豁，超时了")
	}

}
