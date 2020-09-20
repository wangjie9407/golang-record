package main

import "fmt"

func main() {
	c1 := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			val, flag := <-c1
			if flag {
				fmt.Printf("这是收到的第%d个值\n", val)
			} else {
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		c1 <- i
		fmt.Printf("这是传出的的第%d个值\n", i)
	}
	close(c1)
	fmt.Println("已处理完所有通道传参")

	<-done
}
