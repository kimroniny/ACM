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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, ca := headA, 0
	hb, cb := headB, 0
	for ha != nil {
		ha = ha.Next
		ca += 1
	}
	for hb != nil {
		hb = hb.Next
		cb += 1
	}
	d := cb - ca
	ha = headA
	hb = headB
	if ca > cb {
		ha = headB
		hb = headA
		d = ca - cb
	}
	fmt.Println(ca, cb, d)
	fmt.Println("before: ", ha.Val, hb.Val)
	for d > 0 {
		hb = hb.Next
		d -= 1
	}
	for ha != nil && hb != nil {
		fmt.Println(ha.Val, hb.Val)
		if ha == hb {
			return ha
		}
		ha = ha.Next
		hb = hb.Next
	}
	return nil
}

func main() {
	a := []int{4, 1, 8, 4, 5}
	b := []int{5, 0, 1, 8, 4, 5}
	heada := &ListNode{Val: 0, Next: nil}
	headb := &ListNode{Val: 0, Next: nil}
	pa := heada
	pb := headb
	for _, v := range a {
		temp := &ListNode{Val: v, Next: nil}
		pa.Next = temp
		pa = temp
	}
	for _, v := range b {
		temp := &ListNode{Val: v, Next: nil}
		pb.Next = temp
		pb = temp
	}
	getIntersectionNode(heada.Next, headb.Next)
}
