package main

import (
	"fmt"

)

// 如果一个变量生命在if后面，则该变量只有在if语句中有效
/**
循环：
	1.常规循环
	for i := 0; i < 10; i++{}
	2. 省略循环
	var i = 0;
	for ;i < 10;{
		i++
	}
	3.死循环
	for {}
	4. range 循环（key:索引， val：值）
	for key,val := range s{}

*/
func main() {
	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
