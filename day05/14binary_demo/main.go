package main

import "fmt"

// 二进制实用途

const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

// 111
// 左边的1表示吃饭    100
// 中间的1表示睡觉    010
// 右边的1表示打豆豆  001

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | da)         // 101
	f(eat | sleep | da) // 111
}
