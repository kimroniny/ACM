package main

import "fmt"

type Info struct {
	flag bool
	cnt  int
}

var ms map[int]*Info

func dfs(x int) int {
	if _, ok := ms[x]; !ok {
		return 0
	}
	if ms[x].cnt == 1 {
		ms[x].cnt += dfs(x - 1)
	}
	return ms[x].cnt
}

func longestConsecutive(nums []int) int {
	ms = make(map[int]*Info)
	for _, v := range nums {
		if _, ok := ms[v]; !ok {
			ms[v] = &Info{flag: true, cnt: 1}
		}
	}
	for _, v := range nums {
		if ms[v].cnt == 1 {
			ms[v].cnt += dfs(v - 1)
		}
	}
	maxv := 0
	for _, v := range ms {
		if v.cnt > maxv {
			maxv = v.cnt
		}
	}
	// for k, i := range ms {
	// 	fmt.Println(k, i.cnt)
	// }
	return maxv

}

func main() {
	fmt.Println(
		longestConsecutive(
			[]int{100, 4, 200, 1, 3, 2},
		),
	)
}
