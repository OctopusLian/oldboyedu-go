/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-04 18:35:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-04 18:41:05
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

// 读操作
func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 2)
	rwLock.RUnlock()
}

// 写操作
func write() {
	defer wg.Done()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
