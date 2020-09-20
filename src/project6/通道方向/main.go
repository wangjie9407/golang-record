package main

import "fmt"

func main() {
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, "i am your father i am so grate, brown you so much")
	go pong(pings, pongs)
	fmt.Println(<-pongs)

}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
