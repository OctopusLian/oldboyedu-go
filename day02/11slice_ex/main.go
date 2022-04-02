package main

import (
	"fmt"
	"sort"
)

// 切片的练习题

func main() {
	var a = make([]int, 5, 10) // 创建切片，长度为5，容量为10
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a) // ? [1 2 3 4 5 6 7 8 9]  [0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a))

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)
}
