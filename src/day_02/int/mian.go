package main

/**
数据类型：整数型
	1.int： 有符号
	2.uint： 无符号
*/
import "fmt"

func main() {
	// go不能直接定义二进制数据，但是可以直接定义八进制和十六进制的变量
	// 8进制
	num1 := 0101
	// 10进制
	num2 := 101
	// 16进制
	num3 := 0x101

	fmt.Println(num1, num2, num3)

	// go还可以通过Printf转换数字格式
	fmt.Printf("%b\n", num2) // 二进制
	fmt.Printf("%o\n", num2) // 八进制
	fmt.Printf("%d\n", num2) // 十进制
	fmt.Printf("%x\n", num2) // 十六进制
	fmt.Printf("%T\n", num2) // 查看数据类型

}
