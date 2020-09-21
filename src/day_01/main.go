package main

import "fmt"

// 声明方式：
// 1.单独声明
var age int

// 2.批量声明
var (
	name string
	flag bool
)

// 3.go语言中声明变量后必须使用，否则会报错
// 4 job := "前端", 简短声明，只能用于函数内
// 5 const用于声明常量，批量声明变量时，如果后一行没有赋值，则默认与上一行赋相同的值

// iota
// 1.const会使iota重置为零，但是批量赋值时只有一个const，故iota只会被重置一次
// 2.const批量赋值中，每新增一行赋值，iota都会被加1，注：与是否使用iota无关，只要新增一行赋值，iota就会被加1

func main() {
	name = "wj"
	age = 26
	job := "前端"
	/**
	fmt可以打印，一共有三种打印方式，其中
		Printf是格式化打印，可以使用%s来替换打印内容
	*/
	fmt.Printf("我是%s", name)
	fmt.Println(age, job)
}
