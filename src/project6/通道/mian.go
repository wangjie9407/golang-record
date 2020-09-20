package main

import (
	"fmt"
)

/**
通道：
	1.用于链接多个协程的管道
	2.可以从一个go协程中将值发送到通道中，然后在另一个协程中接收
*/

func main() {
	// 可以使用make创建go通道，声明通道时需要定义类型
	message := make(chan string)

	// 在协程中对通道赋值，使用 "channel<-" 语法将值送入通道中
	go func() { message <- "我是协程中被赋值的方法" }()

	// 使用 "<- channel "语法将通道中的值赋予变量
	msg := <-message

	fmt.Print(msg)
}
