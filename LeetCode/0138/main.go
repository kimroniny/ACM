package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// var newhead *Node
	if head == nil {
		return nil
	}
	p := head.Next
	newhead := &Node{Val: head.Val, Next: nil, Random: nil}
	newp := newhead
	for p != nil {
		temp := &Node{Val: p.Val, Next: nil, Random: nil}
		newp.Next = temp
		newp = newp.Next
		p = p.Next
	}
	p = head
	newp = newhead
	for p != nil {
		if p.Random != nil {
			pr := p
			newpr := newp
			for pr != p.Random {
				if pr == nil {
					pr = head
					newpr = newhead
					continue
				}
				pr = pr.Next
				newpr = newpr.Next
			}
			newp.Random = newpr
		}
		p = p.Next
		newp = newp.Next
	}
	return newhead
}

func main() {

}
