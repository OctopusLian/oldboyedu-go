package main

import (
	"encoding/json"
	"fmt"
)

// 复习结构体

type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

func sum(x int, y int) (ret int) {
	ret = x + y
	return ret
}

// 构造(结构体变量的)函数,返回值是对应的结构体类型
func newPerson(n string, i int) (p person) {
	p = person{
		name: n,
		age:  i,
	}
	return p
}

// 方法
// 接收者使用对应类型的首字母小写
// 指定了接收者之后,只有接收者这个类型的变量才能调用这个方法
func (p *person) dream(str string) {
	fmt.Printf("%s的梦想是%s.\n", p.name, str)
}

// func (p person) guonian() {
// 	p.age++ // 此处的p是p1的副本,改的是副本
// }

// 指针接收者
// 1. 需要修改结构体变量的值时要使用指针接收者
// 2. 结构体本身比较大,拷贝的内存开销比较大时也要使用指针接收者
// 3. 保持一致性:如果有一个方法使用了指针接收者,其他的方法为了统一也要使用指针接收者
func (p *person) guonian() {
	p.age++ // 此处的p是p1的副本,改的是副本
}

func main() {
	var b = tmp{
		10,
		20,
	}
	fmt.Println(b)

	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	var x int
	y := int8(10)
	fmt.Println(x, y)

	var p1 person // 结构体实例化
	p1.name = "周林"
	p1.age = 9000

	p2 := person{"保德路", 18} // 结构体实例化
	p3 := person{"马笑", 20}

	// 调用构造函数生成person类型变量
	p4 := newPerson("nazha", 18)

	fmt.Println(p1, p2, p3, p4)
	p1.dream("做个咸鱼")
	p2.dream("学好Go语言")

	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)

	// 结构体嵌套
	type addr struct {
		province string
		city     string
	}
	type student struct {
		name string
		addr // 匿名嵌套别的结构体,就使用类型名做名称
	}

	type point struct {
		X int `json:"zhoulin"`
		Y int `json:"baodelu"`
	}

	po1 := point{100, 200}
	// 序列化
	b1, err := json.Marshal(po1)
	// 如果出错了
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Println(string(b1))

	// 反序列化:由字符串 --> Go中的结构体变量
	str1 := `{"zhoulin":10010,"baodelu":10086}`
	var po2 point // 造一个结构体变量,准备接收反序列化的值
	err = json.Unmarshal([]byte(str1), &po2)
	if err != nil {
		fmt.Printf("unmarshal failed,err:%v\n", err)
	}
	fmt.Println(po2)

	// map

	m1 := map[int64]string{
		10081: "哈哈哈",
		10010: "嘿嘿嘿",
		10000: "呵呵呵",
	}

	name1 := m1[20000]
	fmt.Println(name1) // 取不到就返回value类型的零值

	name2, ok := m1[2000] // ok=true表示有这个key,ok=false表示没有这key
	fmt.Println(name2, ok)

	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }

	// for k := range m1 {
	// 	fmt.Println(k)
	// }
}
