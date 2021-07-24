package main

import (
	"fmt"
)

func restoreArray(adjacentPairs [][]int) []int {
	maps := make(map[int][]int)
	for _, v := range adjacentPairs {
		maps[v[0]] = append(maps[v[0]], v[1])
		maps[v[1]] = append(maps[v[1]], v[0])
	}
	// fmt.Println(maps)
	p := -100050
	for k, v := range maps {
		if len(v) == 1 {
			p = k
			break
		}
	}
	flag := make(map[int]bool)
	flag[p] = true
	ans := make([]int, 0, len(adjacentPairs)+1)
	for {
		ans = append(ans, p)
		f := true
		for _, v := range maps[p] {
			if !flag[v] {
				p = v
				f = false
				flag[p] = true
				continue
			}
		}
		if f {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(
		restoreArray(
			[][]int{
				{2, 1},
				{3, 4},
				{3, 2},
				// {},
			},
		),
	)
}
