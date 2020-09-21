package main

import (
	"fmt"
	"time"

)

/**
1.相对于进程和线程，所有的协程在用户态完成，创建和切换时消耗的性能更低
2.一个进程内可运行多个线程，一个线程内可运行多个协程，同一个线程内，
最多只会有一个协程运行
3.
*/

func main() {
	str1 := "我是你爸爸我真伟大，养你这么大"
	str2 := "我不是你爸爸也伟大，老子最伟大"
	str3 := "我不管，反正我最伟大"

	printStr(str1, "str1")
	go printStr(str2, "str2")
	printStr(str3, "str3")
	time.Sleep(time.Second * 1)
	fmt.Println("done")

}

func printStr(str string, strName string) {
	for _, val := range str {
		fmt.Printf("%s:%c\n", strName, val)
	}
}
