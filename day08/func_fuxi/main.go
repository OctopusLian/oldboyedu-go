package main

import "fmt"

// 函数可变参数

func f1(a ...interface{}) {
	fmt.Printf("type:%T value:%#v\n", a, a)
}

func main() {
	// f1()
	// f1(1)
	// f1(1, false, "a", struct{}{}, []int{1, 2}, [...]int{1, 2, 3}, map[string]int{"zhoulin": 9000})

	var s = []interface{}{1, 3, 5, 7, 9}
	f1(s)
	f1(s...)
}
