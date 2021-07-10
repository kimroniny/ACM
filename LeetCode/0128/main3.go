package main

import "fmt"

var ms map[int]int

func longestConsecutive(nums []int) int {
	ms = make(map[int]int)
	for _, v := range nums {
		ms[v] = 0
	}

	maxv := 0
	for _, v := range nums {
		if ms[v] > 0 {
			continue
		}
		x := v - 1
		cnt := 1
		for {
			if _, ok := ms[x]; ok {

				if ms[x] > 1 {
					cnt += ms[x]
					break
				} else {
					ms[x] = 1
					cnt += ms[x]
				}
				x -= 1
			} else {
				break
			}
		}
		ms[v] = cnt
		if maxv < ms[v] {
			maxv = ms[v]
		}
	}
	return maxv

}

func main() {
	fmt.Println(
		longestConsecutive(
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		),
	)
}
