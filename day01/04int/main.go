package main

import "fmt"

// 整型

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 把十进制数转换成二进制
	fmt.Printf("%o\n", i1) // 把十进制数转换成八进制
	fmt.Printf("%x\n", i1) // 把十进制数转换成十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9) // 明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}
