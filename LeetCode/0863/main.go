package main

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

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	var dfs func(node *TreeNode, cnt int) (uint8, int)
	var dfsK func(node *TreeNode, cnt int, k int)
	dfs = func(node *TreeNode, cnt int) (uint8, int) {
		if node == nil {
			return 3, -1
		}
		if node == target {
			return 0, cnt
		}
		flag := dfs(node.Left, cnt+1)
		if flag > -1 {
			return 1, flag
		}
		flag = dfs(node.Right, cnt+1)
		return 2, flag
	}
	dfsK = func(node *TreeNode, cnt int, k int) {
		if cnt > k {
			return
		}
		if cnt == k {
			result = append(result, node.Val)
			return
		}
	}
	flag, dist := dfs(root, 0)
	if flag == 3 {
		return []int{}
	}
	if flag == 1 {
		dfsK(target, 0, k)
		if dist == k {
			result = append(result, root.Val)
		}else {
				
		}
		
	}

}

func main() {

}
