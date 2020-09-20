package main

import "fmt"

func main() {
	a, b, c := multipleVal()
	fmt.Println(a, b, c)
}

func multipleVal() (int, int, int) {
	return 3, 5, 7
}
