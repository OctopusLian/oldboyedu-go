package main

import "fmt"

// 关闭通道

func main() {
	ch1 := make(chan bool, 2)
	ch1 <- true
	ch1 <- true
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)
}
