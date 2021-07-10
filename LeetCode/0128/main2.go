package main

import "fmt"

var ms map[int]int

func dfs(x int) int {
	if _, ok := ms[x]; !ok {
		return 0
	}
	if ms[x] == 0 {
		ms[x] = dfs(x-1) + 1
	}
	return ms[x]
}

func longestConsecutive(nums []int) int {
	ms = make(map[int]int)
	for _, v := range nums {
		ms[v] = 0
	}

	maxv := 0
	for _, v := range nums {
		if ms[v] == 0 {
			ms[v] = dfs(v-1) + 1
		}
		if maxv < ms[v] {
			maxv = ms[v]
		}
	}
	return maxv

}

func main() {
	fmt.Println(
		longestConsecutive(
			[]int{100, 4, 200, 1, 3, 2},
		),
	)
}
