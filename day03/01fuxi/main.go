package main

import "fmt"

// day03复习

func main() {
	// 	var name string
	// 	name = "理想"
	// 	fmt.Println(name)
	// 	var ages [30]int // 声明了一个变量ages,它是[30]int类型
	// 	ages = [30]int{1, 2, 3, 4, 6}
	// 	fmt.Println(ages)
	// 	var ages2 = [...]int{1, 2, 3, 4, 5, 7, 8, 90}
	// 	fmt.Println(ages2)
	// 	var ages3 = [...]int{1: 100, 99: 200}
	// 	fmt.Println(ages3)
	// 	// 二维数组
	// 	var a1 = [...][2]int{
	// 		[2]int{1, 2},
	// 		[2]int{3, 4},
	// 		[2]int{5, 6},
	// 	}
	// 	fmt.Println(a1)
	// 	// 多维数组只有最外层可以使用...

	// 	// 数组是值类型

	// 	x := [3]int{1, 2, 3}
	// 	y := x     // ? 把x的值拷贝了一份给了y
	// 	y[1] = 200 // 修改的是副本y,并不影响x

	// 	fmt.Println(x) // [1 2 3]
	// 	f1(x)
	// 	fmt.Println(x) // ?[1 2 3]

	// 切片(Slice):
	// var s1 []int // 没有分配内存, == nil
	// fmt.Println(s1)
	// fmt.Println(s1 == nil)
	// s1 = []int{1, 2, 3}
	// fmt.Println(s1)
	// // make初始化 分配内存
	// s2 := make([]bool, 2, 4)
	// fmt.Println(s2) // [false false]
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3 == nil)

	// s1 := []int{1, 2, 3} // [1 2 3]
	// s2 := s1
	// var s3 = make([]int, 3, 3)
	// copy(s3, s1)
	// fmt.Println(s2) // [1 2 3]
	// s2[1] = 200
	// fmt.Println(s2) // [1 200 3]
	// fmt.Println(s1) // [1 200 3]
	// fmt.Println(s3) // ?[1 2 3]

	// var s1 []int // nil
	// // s1 = make([]int, 1)
	// // s1[0] = 100
	// // fmt.Println(s1)
	// s1 = append(s1, 1) // 自动初始化切片
	// fmt.Println(s1)

	// 指针
	// Go里面的指针只能读不能修改，不能修改指针变量指向的地址
	// addr := "沙河"
	// addrP := &addr
	// fmt.Println(addrP) // 内存地址
	// fmt.Printf("%T\n", addrP)
	// addrV := *addrP // 根据内存地址找值
	// fmt.Println(addrV)

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["理想"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"]) // 如果key不存在返回的是value对应类型的零值
	// 如果返回值是布尔型，我们通常用ok去接收它
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数是", score)
	}
	delete(m1, "lixiang") // 删除的key不存在什么都不干
	delete(m1, "理想")
	fmt.Println(m1)
	fmt.Println(m1 == nil) // 已经开辟了内存
}

func f1(a [3]int) {
	// Go语言中的函数传递的都是值（Ctrl+C Ctrl+V）
	a[1] = 100 // 此处修改的是副本的值
}
