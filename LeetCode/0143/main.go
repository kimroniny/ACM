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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	p, cnt := head, 0
	for p != nil {
		cnt += 1
		p = p.Next
	}
	// for p = head; p != nil && cnt != 0; p, cnt = p.Next, cnt-1 {
	// 	fmt.Println(p.Val)
	// }
	// fmt.Println("------")
	pre := cnt / 2
	p = head
	var newp *ListNode
	var dfs func(p *ListNode, x int)
	dfs = func(p *ListNode, x int) {
		if x == pre+1 {
			if cnt-pre != pre {
				newp = p.Next
				p.Next = nil
			} else {
				newp = p
			}
			return
		}
		dfs(p.Next, x+1)

		temp := p.Next
		p.Next = newp
		temp2 := newp.Next
		if newp == temp {
			newp.Next = nil
		} else {
			newp.Next = temp
		}
		newp = temp2

	}
	dfs(p, 1)
	// for p = head; p != nil && cnt != 0; p, cnt = p.Next, cnt-1 {
	// 	fmt.Println(p.Val)
	// }
}

func main() {
	infos := []int{0}
	head := &ListNode{Val: -1, Next: nil}
	p := head
	for _, v := range infos {
		temp := &ListNode{Val: v, Next: nil}
		p.Next = temp
		p = temp
	}
	reorderList(head.Next)
	p = head.Next

}
