package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)

	msg1 := <-time1.C
	fmt.Println(msg1, "time1已定时完成")

	time2 := time.NewTimer(time.Second * 2)
	go func() {
		<-time2.C
		fmt.Println("time2已定时完成")
	}()

	// stop := time2.Stop()
	// if stop {
	// 	fmt.Println("time2已被取消")
	// }
}
