package main

import "fmt"

var maps map[int][]int
var num int
var ans int

func bfs(v, k int) {
	if k == 0 {
		if v == num {
			ans += 1
		}
		return
	}
	for _, nextv := range maps[v] {
		bfs(nextv, k-1)
	}
}

func numWays(n int, relation [][]int, k int) int {
	num = n - 1
	ans = 0
	maps = make(map[int][]int)
	for _, v := range relation {
		if len(maps[v[0]]) == 0 {
			maps[v[0]] = make([]int, 0, n)
		}
		maps[v[0]] = append(maps[v[0]], v[1])
	}
	bfs(0, k)
	return ans
}

func main() {
	fmt.Println(
		numWays(
			3,
			[][]int{
				{0, 2},
				{2, 1},
			},
			2,
			// 5,
			// [][]int{
			// 	{0, 2},
			// 	{2, 1},
			// 	{3, 4},
			// 	{2, 3},
			// 	{1, 4},
			// 	{2, 0},
			// 	{0, 4},
			// },
			// 3,
		),
	)
}
