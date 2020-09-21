package main

import "fmt"

type obj struct {
	name string
	age  int
}

// 可以在结构体中定义方法

func (o obj) instruct() {
	fmt.Printf("我是%s,我%d岁了", o.name, o.age)
}

func main() {
	person := obj{"胖胖", 3}
	person.name = "肉肉"
	person.age = 10
	person.instruct()
	// fmt.Println(person)
}
