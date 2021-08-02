package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	x   int
	val int
}

type infos []*Info

func (a infos) Len() int      { return len(a) }
func (a infos) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a infos) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].val < a[j].val
	}
	return a[i].x < a[j].x
}

func verticalTraversal(root *TreeNode) [][]int {

	var data infos
	data = make(infos, 0, 1050)
	var dfs func(node *TreeNode, x, y int)
	dfs = func(node *TreeNode, x, y int) {
		if node == nil {
			return
		}

		data = append(data, &Info{x: x, val: node.Val})
		dfs(node.Left, x-1, y+1)
		dfs(node.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
	sort.Sort(data)
	result := make([][]int, 0, 1050)
	l := len(result)
	temp := make([]int, 0, 1050)
	result = append(result, temp)
	result[l] = append(result[l], data[0].val)
	for i := 1; i < data.Len(); i++ {
		if data[i].x == data[i-1].x {
			result[l] = append(result[l], data[i].val)
		} else {
			l = len(result)
			last := result[l-1]
			sort.Ints(last)
			result[l-1] = last
			temp := make([]int, 0, 1050)
			result = append(result, temp)
			result[l] = append(result[l], data[i].val)
		}
	}
	return result
}

func main() {

}
