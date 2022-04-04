/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-04 18:36:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-04 19:00:29
 */
package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数
func main() {
	// 创建一个标志位参数
	name := flag.String("name", "王冶", "请输入名字")
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了吗")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")

	flag.Parse() // 先解析再使用

	fmt.Println(*name)  //王冶
	fmt.Println(*age)  //9000
	fmt.Println(*married)  //false
	fmt.Println(*cTime)
	fmt.Printf("%T\n", *cTime)

	// var name string
	// flag.StringVar(&name, "name", "王冶", "请输入名字")
	// 使用flag
	// flag.Parse()
	// fmt.Println(name)

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
