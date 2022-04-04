package main

import (
	"fmt"
	"os"
)

// 1. 文件对象的类型
// 2. 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("D:/test.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	// 1. 文件对象的类型
	fmt.Printf("%T\n", fileObj)
	// 2. 获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return
	}
	fmt.Printf("文件大小是:%dB\n", fileInfo.Size())
	fmt.Printf("文件名是:%s\n", fileInfo.Name())
}
