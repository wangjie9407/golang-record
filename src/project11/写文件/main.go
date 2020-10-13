package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// var f, err = os.Create("../file/test2.txt")
	var f, err = os.OpenFile("../file/test2.txt", os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	writer := bufio.NewWriter(f)
	// writer.WriteString("没人见过你咬牙坚持的样子，但你深知无力的瞬间")
	writer.WriteString("记住永远努力，记住永远年轻")
	writer.Flush()
}
