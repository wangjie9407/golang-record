package main

import "fmt"

/**
通道缓冲：
	1.默认通道无缓冲
	2.通道准备好接收时才允许发送
	3.可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值
*/

func main() {
	// 定义可缓存数量的值
	message := make(chan string, 2)

	go func() {
		message <- "this is first"
		message <- "and i am the second one"
	}()

	fmt.Println(<-message)
	fmt.Println(<-message)
}
