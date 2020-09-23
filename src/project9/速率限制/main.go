package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		results <- i
	}
	close(results)

	limiter := time.Tick(time.Millisecond * 200)

	for res := range results {
		// 速率限制，每隔200ms打印一次通道值
		<-limiter
		fmt.Println(res)
	}
}
