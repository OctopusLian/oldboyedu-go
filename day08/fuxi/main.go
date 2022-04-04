package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel

// 往通道中发送值
func sendNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(5 * time.Second)
	}
}
func main() {
	ch := make(chan int, 1)
	// ch <- 100 // 把一个值发送到通道中
	// ch <- 200
	// <-ch          // 把通道中把100取出来
	go sendNum(ch)
	for {
		x, ok := <-ch // 阻塞等4秒钟
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}
}
