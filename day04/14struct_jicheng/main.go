package main

import "fmt"

// 结构体模拟实现其他语言中的"继承"

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

// 狗类
type dog struct {
	feet   uint8
	animal // animal拥有的方法,dog此时也有了【www.hlwnjl.com(互联网农家乐)下载免费程序员教程】
}

// 给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在叫:汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "周林"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move() // ?
}
