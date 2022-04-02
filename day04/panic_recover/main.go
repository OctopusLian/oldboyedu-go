package main

import "fmt"

// panic和recover

func f1() {
	defer func() {
		err := recover() // 收集当前的错误
		fmt.Println("松手去爱...")
		fmt.Println(err)
	}()
	panic("犯了不可原谅的错误")
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
