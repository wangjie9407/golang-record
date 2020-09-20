package main

import "fmt"

// 1.gofmt -w xxx.go指令可格式化文档
// 2.大括号的起始括号不可换行
// 3.go语言一行最好不要超过八十个子符，可以使用逗号换行

func main() {
	fmt.Println("暗裔剑魔\t麦林炮手\t远古巫灵\n亚托克斯\t小炮\t泽拉斯")
	fmt.Println("暗裔剑魔\t麦林炮手\t远古巫灵\n亚托克斯\t小炮\t泽拉斯")
	fmt.Println("暗裔剑魔\t`	麦林炮手\t远古巫灵\n亚托克斯\t小炮\t泽拉斯")
	fmt.Println("暗裔剑魔\t麦林炮手\t远古巫灵\n亚托克斯\t小炮\t泽拉斯")
}
