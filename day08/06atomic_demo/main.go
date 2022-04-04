package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	// x++
	// lock.Lock()
	// x = x + 1
	// lock.Unlock()

	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	// 比较并交换
	// x = 200
	// ok := atomic.CompareAndSwapInt64(&x, 200, 300)
	// fmt.Println(ok, x)
}
