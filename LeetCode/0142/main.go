package main

import (
	"fmt"
	"unsafe"
)

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

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	c1, c2 := 0, 0
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
		c1 += 1
		c2 += 2
		if p1 == p2 {
			break
		}
	}
	diff := 1
	p1, p2 = head, head
	for {
		if diff == c2-c1 {
			if p1 == p2 {
				return p1
			}
			p2 = p2.Next
		} else {
			diff += 1
		}
		p1 = p1.Next
	}
}

func main() {
	a := new(int)
	*a = 10
	fmt.Printf("%p, %T, %d\n", a, a, unsafe.Pointer(a))
	var temp *int
	temp = a
	fmt.Println(*a, *temp)
}
