# day02 课上笔记

# 【免费学习go语言资料www.5lmh.com】

# 内容回顾

## GO安装

`$GOPATH`: 你写Go代码的工作区，保存你的Go代码的。

`go env`

![1562377877799](assets/1562377877799.png)

`GOPATH/bin`添加到环境变量：`go install` 命名会把生成的二进制可执行文件拷贝到`GOPATH/bin`

## Go 命令

`go build` :编译Go程序

`go build -o "xx.exe"` ：编译成xx.exe文件

`go run main.go`: 像执行脚本一样执行main.go文件

`go install`: 先编译后拷贝

## GO语言文件基础语法

存放Go源代码的文件后缀名是`.go`

文件第一行：`package`关键字声明包名

如果要编译可执行文件，必须要有main包和main函数（入口函数）

```go
// 单行注释

/*
多行注释
*/
```

Go语言函数外的语句必须以关键字开

函数内部定义的变量必须使用

## 变量

3种声明方式：

1. `var name1 string`
2. `var name2 = "沙河娜扎"`
3.  函数内部专属：`name3:="沙河小王子"`

匿名变量（哑元变量）：

当有些数据必须用变量接收但是又不使用它时，就可以用`_`来接收这个值。

## 常量

`const PI = 3.1415926`

`const UserNotExistErr = 10000`



iota:实现枚举

两个要点：

1. `iota`在const关键字出现时将被重置为0
2. const中每新增一行常量声明，iota累加1

## 流程控制

#### if 

```go
var age = 19
	if age > 18 {
		fmt.Println("成年了")
	} else if age > 7 {
		fmt.Println("上小学")
	} else {
		fmt.Println("最快乐的时光！")
	}
```

### for循环

标准for循环

```go
for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
```

## 基本数据类型

### 整型

​	无符号整型：`uint8`、`uint16`、`uint32`、`uint64`

​	带符号整型:`int8`、`int16`、`int32`、`int64`

​	`uint`和`int`:具体是32位还是64位看操作系统

​	`uintptr`:表示指针



### 其他进制数

Go语言中没办法直接定义二进制数

```go
// 八进制数
	var n1 = 0777
	// 十六进制数
	var n2 = 0xff
	fmt.Println(n1, n2)
	fmt.Printf("%o\n", n1)
	fmt.Printf("%x\n", n2)
```



### 浮点型

`float64`和`float32`

Go语言中浮点数默认是`float64`

### 复数

`complex128`和`complex64`

### 布尔值

`true`和`false`

不能和其他的类型做转换

### 字符串

常用方法

字符串不能修改

### byte和rune类型

都属于类型别名



### 字符串、字符、字节 都是什么？

字符串：双引号包裹的是字符串

字符：单引号包裹的是字符，单个字母、单个符号、单个文字

字节：1byte=8bit

go语言中字符串都是UTF8编码，UTF8编码中一个常用汉字一般占用3个字节。



# 今日内容

流程控制

## 运算符

[https://www.liwenzhou.com/posts/Go/03_operators/](https://www.liwenzhou.com/posts/Go/03_operators/)

详见课上代码`04/main.go`

## 复合数据类型

## 数组

### 数组的声明



```go
// 数组是存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
var a1 [3]bool // [true false true]
var a2 [4]bool // [true true false false]

fmt.Printf("a1:%T a2:%T\n", a1, a2)
```

### 数组的初始化

```go
// 数组的初始化
// 如果不初始化：默认元素都是零值（布尔值：false, 整型和浮点型都是0, 字符串：""）
fmt.Println(a1, a2)
// 1. 初始化方式1
a1 = [3]bool{true, true, true}
fmt.Println(a1)
// 2. 初始化方式2：根据初始值自动推断数组的长度是多少
// a10 := [9]int{0, 1, 2, 3, 4, 4, 5, 6, 7}
a10 := [...]int{0, 1, 2, 3, 4, 4, 5, 6, 7}
fmt.Println(a10)
// 3. 初始化方式3：根据索引来初始化
a3 := [5]int{0: 1, 4: 2}
fmt.Println(a3)
```

### 数组的遍历

```go
// 数组的遍历
citys := [...]string{"北京", "上海", "深圳"} // 索引：0~2 citys[0],citys[1],citys[2]
// 1. 根据索引遍历
for i := 0; i < len(citys); i++ {
	fmt.Println(citys[i])
}
// 2. for range遍历
for i, v := range citys {
	fmt.Println(i, v)
}
```

### 二维数组

```go
// 多维数组
// [[1 2] [3 4] [5 6]]
var a11 [3][2]int
a11 = [3][2]int{
	[2]int{1, 2},
	[2]int{3, 4},
	[2]int{5, 6},
}
fmt.Println(a11)

// 多维数组的遍历
for _, v1 := range a11 {
	fmt.Println(v1)
	for _, v2 := range v1 {
		fmt.Println(v2)
	}
}
```

### 数组是值类型

```go
// 数组是值类型
b1 := [3]int{1, 2, 3} // [1 2 3]
b2 := b1              // [1 2 3] Ctrl+C Ctrl+V => 把world文档从文件夹A拷贝到文件夹B
b2[0] = 100           // b2:[100 2 3]
fmt.Println(b1, b2)   // b1:[1 2 
```

## 切片（slice）



切片指向了一个底层的数组。

切片的长度就是它元素的个数。

切片的容量是底层数组从切片的第一个元素到最后一个元素的数量。

### 切片的定义

```go
// 切片的定义
var s1 []int    // 定义一个存放int类型元素的切片
var s2 []string // 定义一个存放string类型元素的切片
fmt.Println(s1, s2)
fmt.Println(s1 == nil) // true
fmt.Println(s2 == nil) // true
```

### 切片的初始化

```go
// 初始化
s1 = []int{1, 2, 3}
s2 = []string{"沙河", "张江", "平山村"}
fmt.Println(s1, s2)
fmt.Println(s1 == nil) // false
fmt.Println(s2 == nil) // false
```

### 切片的长度和容量

```go
// 长度和容量
fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2)
```



### make

make()函数用于创建指定长度和容量的切片。

```go
s1 := make([]int, 5, 10)
fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

s2 := make([]int, 0, 10)
fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s2, len(s2), cap(s2))
```

### 切片的本质

切片就是一个框，框住了一块连续的内存。

切片属于引用类型，真正的数据都是保存在底层数组里的。



判断一个切片是否是空的，要是用`len(s) == 0`来判断

### append

```go
// 调用append函数必须用原来的切片变量接收返回值
// append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
// 必须用变量接收append的返回值
s1 = append(s1, "广州")
fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
s1 = append(s1, "杭州", "成都")
fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
ss := []string{"武汉", "西安", "苏州"}
s1 = append(s1, ss...) // ...表示拆开
fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
```

### copy

```go
a1 := []int{1, 3, 5}
a2 := a1 // 赋值
var a3 = make([]int, 3, 3)
copy(a3, a1) // copy
fmt.Println(a1, a2, a3)
a1[0] = 100
fmt.Println(a1, a2, a3)
```

## 指针

Go语言中不存在指针操作，只需要记住两个符号：

1. `&`:取地址
2. `*`:根据地址取值

## make和new的区别

1. make和new都是用来申请内存的
2. new很少用，一般用来给基本数据类型申请内存，`string`、`int`,返回的是对应类型的指针(\*string、\*int)。
3. make是用来给`slice`、`map`、`chan`申请内存的，make函数返回的的是对应的这三个类型本身

## map

map也是引用类型，必须初始化之后才能使用。

```go
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["jiwuming"] = 35

	fmt.Println(m1)
	fmt.Println(m1["理想"])
	// 约定成俗用ok接收返回的布尔值
	fmt.Println(m1["娜扎"]) // 如果不存在这个key拿到对应值类型的零值
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "jiwuming")
	fmt.Println(m1)
	delete(m1, "沙河") // 删除不存在的key
}
```

## 函数

```go
// 函数

// 函数存在的意义？
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用就可以了
// 使用函数能够让代码结构更清晰、更简洁。

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值的
func f3() int {
	ret := 3
	return ret
}

// 返回值可以命名也可以不命名

// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值可以return后省略
}

// 多个返回值
func f5() (int, string) {
	return 1, "沙河"
}

// 参数的类型简写:
// 当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是切片 []int
}

// Go语言中函数没有默认参数这个概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f7("下雨了")
	f7("下雨了", 1, 2, 3, 4, 5, 6, 7)
}
```





匿名函数和闭包



# 作业

1. 把运算符的所有例子自己写一下（详见课上代码 `04op/main.go`）
2. 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
3. 把函数部分7、8个示例自己签一遍。
4. 路飞上的作业

