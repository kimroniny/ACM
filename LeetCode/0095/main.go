package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []*TreeNode
var head *TreeNode
var allnum int

func getMidOrder(node *TreeNode, p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}
	leftNode := new(TreeNode)
	rightNode := new(TreeNode)
	node.Val = p.Val
	node.Left = getMidOrder(leftNode, p.Left)
	node.Right = getMidOrder(rightNode, p.Right)
	return node
}

func dfs(node *TreeNode, l, r int) {
	if l > r {
		return
	}
	leftnode := new(TreeNode)
	rightnode := new(TreeNode)
	for i := l; i <= r; i++ {
		allnum -= 1
		// node = &TreeNode{Val: i, Left: leftnode, Right: rightnode}
		node.Val = i
		node.Left = leftnode
		node.Right = rightnode
		dfs(leftnode, l, i-1)
		dfs(rightnode, i+1, r)
		if allnum == 0 {
			
			newHead := new(TreeNode)
			result = append(result, getMidOrder(newHead, head))
		}
		allnum += 1
	}
}

func generateTrees(n int) []*TreeNode {
	leftnode := new(TreeNode)
	rightnode := new(TreeNode)
	for i := 1; i <= n; i++ {
		allnum = n - 1
		head = &TreeNode{Val: i, Left: leftnode, Right: rightnode}
		dfs(leftnode, 1, i-1)
		dfs(rightnode, i+1, n)
		if allnum == 0 {
			newHead := new(TreeNode)
			result = append(result, getMidOrder(newHead, head))
		}
		allnum += 1
	}
	return result
}

func prindfs(node *TreeNode) {
	if node != nil {
		fmt.Println(node.Val)
		prindfs(node.Left)
		prindfs(node.Right)
	}
}

func prin() {
	for _, v := range result {
		fmt.Println("=========================")
		prindfs(v)
	}
}

func main() {
	generateTrees(3)
	prin()
}
