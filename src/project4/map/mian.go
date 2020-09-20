package main

import "fmt"

/**
map:
	3.使用fmt.Println打印map对象时，会显示全部的键值对
	5.map可使用len方法来获取键值对数目len(map1)

*/

func main() {
	//1.创建方式： var map1 = make(map[key: type]value: type)
	map1 := make(map[string]int)

	// 2.设置键值对 map1[key: type] = value
	map1["k1"] = 7
	map1["k2"] = 8
	map1["k3"] = 7
	//6.可使用delete方法删除指定键值对delete(map, key)
	delete(map1, "k2")

	//4.可以使用map1[key]来获取对应的键值对
	//7.返回的第二个值用域判断该键是否存在于该map中
	k3, hasOwnProperty := map1["k3"]

	// 初始化map的方法
	map2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(map1)
	fmt.Println(k3, hasOwnProperty)
	fmt.Println(map2)
}
