/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	minival := root.Val
	secondval := math.MaxInt64
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val < secondval && root.Val > minival {
			secondval = root.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	if secondval == math.MaxInt64 {
		return -1
	}
	return secondval
}

func buildTree(node *TreeNode, p int, vals []int) {
	if p >= len(vals) {
		return
	}
	if vals[p] == -1 {
		node.Left = nil
	} else {
		node.Left = &TreeNode{Val: vals[p], Left: nil, Right: nil}
		buildTree(node.Left, (p+1)*2-1, vals)
	}
	if vals[p+1] == -1 {
		node.Right = nil
	} else {
		node.Right = &TreeNode{Val: vals[p+1], Left: nil, Right: nil}
		buildTree(node.Right, (p+2)*2-1, vals)
	}
}

func printTree(head *TreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	printTree(head.Left)
	printTree(head.Right)
}

func main() {
	vals := []int{2, 2, 2}
	head := &TreeNode{Val: vals[0], Left: nil, Right: nil}
	buildTree(head, 1, vals)
	printTree(head)
	fmt.Println("------------")
	fmt.Println(findSecondMinimumValue(head))
}
