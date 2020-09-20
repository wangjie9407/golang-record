package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// select选择器类似switch语句，用于监听channel的I/O操作，当有一个chan执行I/O完成时，则会执行对应的case
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
