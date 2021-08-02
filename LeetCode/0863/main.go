package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	node   *TreeNode
	dist   int
	direct bool
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	route := make([]*Info, 0)
	cnt := 0
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if node == target {
			route = append(route, &Info{node: node, dist: cnt, direct: true})
			cnt += 1
			return true
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left {
			route = append(route, &Info{node: node, dist: cnt, direct: true})
			cnt += 1
			return true
		}
		if right {
			route = append(route, &Info{node: node, dist: cnt, direct: false})
			cnt += 1
			return true
		}
		return false
	}
	dfs(root)
	
	// for _, v := range route {
	// 	fmt.Println(*v)
	// }

	result := make([]int, 0)
	var dfsK func(node *TreeNode, cnt int)
	dfsK = func(node *TreeNode, cnt int) {
		if node == nil {
			return
		}
		if cnt == 0 {
			result = append(result, node.Val)
			return
		}
		dfsK(node.Left, cnt-1)
		dfsK(node.Right, cnt-1)
	}
	for _, info := range route {
		dist := k - info.dist
		if info.dist == 0 {
			if dist == 0 {
				result = append(result, info.node.Val)
			} else {
				dfsK(info.node.Left, dist-1)
				dfsK(info.node.Right, dist-1)
			}
		} else {
			if dist == 0 {
				result = append(result, info.node.Val)
			} else {
				if info.direct {
					dfsK(info.node.Right, dist-1)
				} else {
					dfsK(info.node.Left, dist-1)
				}
			}
		}
	}
	return result
}

func buildTree(p int, arrs []int) *TreeNode {
	if p >= len(arrs) || arrs[p] == -1 {
		return nil
	}
	node := &TreeNode{}
	node.Val = arrs[p]
	pl := (p+1)*2 - 1
	pr := pl + 1
	node.Left = buildTree(pl, arrs)
	node.Right = buildTree(pr, arrs)
	return node
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

func main() {
	// arrs := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	arrs := []int{3, 5}
	node := buildTree(0, arrs)
	printTree(node)
	target := node.Left
	fmt.Println(
		distanceK(node, target, 0),
	)
}
