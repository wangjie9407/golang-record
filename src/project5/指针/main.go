package main

import "fmt"

func main() {
	var num = 2
	// 将num的指针赋值给num2, &符是根据变量获取地址
	var num2 = &num
	//  *符是根据地址获取变量, 直接操作地址，会影响赋值的变量
	*num2 = 3

	fmt.Println(num, *num2)
}
