package main

import (
	"fmt"
)

func main() {
	// 非阻塞通道，即select中使用default处理未被监听的通道
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c2 <- "哈撒ki"
		c1 <- "面对疾风咯"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c2:
			fmt.Println(msg)
		default:
			fmt.Println("哈哈哈")
		}
	}
}
