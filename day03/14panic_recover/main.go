package main

import "fmt"

// panic 和 recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}()
	panic("出现了严重的错误！！！") // 程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}
func main() {
	funcA()
	funcB()
	funcC()
}
