/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-04 18:36:12
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-04 19:10:10
 */
package main

// 如何判断一个链表有没有闭环

type a struct {
	val  int
	next *a
}

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}

func main() {

}
