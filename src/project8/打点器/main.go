package main

import (
	"fmt"
	"time"
)

func main() {\
	// 打点器，即类似于js中的setInterval，每隔固定时间执行一次回调函数
	timeTricker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range timeTricker.C {
			fmt.Println("lalala", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	timeTricker.Stop()
	fmt.Println("打点器结束")
}
