package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next // 把下一个节点缓存起来
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Printf("%#v\n", head)
	ret := reverseList(head)
	fmt.Printf("%#v\n", ret)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
