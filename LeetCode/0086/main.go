package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var littleNodeHead, littleNodeTail, bigNodeHead, bigNodeTail *ListNode = nil, nil, nil, nil
	for head != nil {
		if head.Val < x {
			if littleNodeHead == nil {
				littleNodeHead = head
				littleNodeTail = littleNodeHead
			} else {
				littleNodeTail.Next = head
				littleNodeTail = littleNodeTail.Next
			}
		} else {
			if bigNodeHead == nil {
				bigNodeHead = head
				bigNodeTail = bigNodeHead
			} else {
				bigNodeTail.Next = head
				bigNodeTail = bigNodeTail.Next
			}
		}
		head = head.Next
	}
	if littleNodeTail != nil {
		littleNodeTail.Next = nil
	}
	if bigNodeTail != nil {
		bigNodeTail.Next = nil
	}

	if littleNodeTail != nil {
		littleNodeTail.Next = bigNodeHead
	} else {
		littleNodeHead = bigNodeHead
	}
	return littleNodeHead
}

func main() {
	headList := []int{1, 4, 3, 2, 5, 2}
	head := &ListNode{Val: 1, Next: nil}
	tail := head
	for _, v := range headList[1:] {
		tail.Next = &ListNode{Val: v, Next: nil}
		tail = tail.Next
	}
	new := partition(head, 0)
	for new != nil {
		fmt.Println(new.Val)
		new = new.Next
	}

}
