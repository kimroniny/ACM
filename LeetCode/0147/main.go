package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	newhead := &ListNode{Val: -10000, Next: head}
	for p, lastp := head, newhead; p != nil; {
		// fmt.Println(p.Val)
		k := newhead.Next
		for lastk := newhead; k != p; k, lastk = k.Next, k {
			if k.Val > p.Val {
				lastp.Next = p.Next
				lastk.Next = p
				p.Next = k
				p = lastp.Next
				break
			}
		}
		if k == p {
			p, lastp = p.Next, p
		}
	}
	// for p := newhead.Next; p != nil; p = p.Next {
	// 	fmt.Println(p.Val)
	// }
	return newhead.Next
}

func main() {
	infos := []int{4, 1,  3, 2}
	head := &ListNode{Val: -1, Next: nil}
	p := head
	for _, v := range infos {
		temp := &ListNode{Val: v, Next: nil}
		p.Next = temp
		p = temp
	}
	insertionSortList(head.Next)

}
