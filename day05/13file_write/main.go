/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-02 13:37:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-02 18:09:54
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
// 0100 0000

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// write
	fileObj.Write([]byte("zhoulin mengbi le!\n"))
	// writeString
	fileObj.WriteString("周林解释不了!")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello沙河\n") // 写到缓存中
	wr.Flush()                  // 将缓存中的内容写入文件
}

func writeDemo3() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()

}
