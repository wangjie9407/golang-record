package main

import "fmt"

func main() {
	// 数组内元素求和
	arr1 := [...]int{1, 3, 5, 7, 8}
	res := 0
	for _, val := range arr1 {
		res += val
	}
	fmt.Println(res)
	// 找出和为指定值的两个元素的下标
	res = 8
	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr1); j++ {
			num := arr1[i] + arr1[j]
			if num == res {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
