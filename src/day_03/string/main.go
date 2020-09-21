package main

import (
	"container/list"
	"fmt"
	"unicode"

)

/**
string 字符串
1.go语言中，字符串用双引号，字符用单引号
*/

func main() {
	// str := `acafacadaeagaf`
	// str2 := strings.Split(str, "a")
	// fmt.Println(str2)
	// fmt.Println(strings.Contains(str, "g"))
	// fmt.Println(strings.Join(str2, "-"))

	listData := list.New()

	listData.PushBack(1)
	listData.PushBack(1.2)
	listData.PushBack(false)
	listData.PushBack("aabc")

	for e := listData.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, "\t")
		fmt.Printf("%T\n", e.Value)
	}

	var str = "hello 啦啦啦 i am 德玛西亚"
	strArr := []rune(str)
	counte := 0
	for i := 0; i < len(strArr); i++ {
		if unicode.Is(unicode.Han, strArr[i]) {
			counte++
			fmt.Printf("%c\n", strArr[i])
		}
	}

}
