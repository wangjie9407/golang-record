package main

import (
	"fmt"
)

func main() {
	done := make(chan string, 2)
	go worker(done)
	go worker(done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}

func worker(done chan string) {
	fmt.Println("i am working....")
	fmt.Println("i was done")
	// done用于通知其它通道该通道已执行完成
	done <- "i am done"
}
