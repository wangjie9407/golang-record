package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p := fmt.Println

	// 读文件方法1
	f1, err1 := ioutil.ReadFile("../file/test.txt")
	check(err1)
	p(string(f1))

	// 读文件方法2
	f2, err2 := os.Open("../file/test.txt")
	check(err2)
	b2 := make([]byte, 5)
	str2, err3 := f2.Read(b2)
	check(err3)
	fmt.Printf("%d bytes: %s\n", str2, string(b2))
}
