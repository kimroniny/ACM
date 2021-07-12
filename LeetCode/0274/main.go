package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	sum := make([]int, len(citations)+1)
	a := 0
	l := len(citations)
	j := l - 1
	for i := citations[l-1]; i >= 0; i-- {
		// fmt.Println(j)
		for j >= 0 && i == citations[j] {
			a += 1
			j -= 1
		}
		if i <= len(citations) {
			sum[i] = a
		}
	}
	// fmt.Println(sum)
	for i := len(sum) - 1; i >= 0; i-- {
		if sum[i] >= i {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex(
		[]int{1,2,2},
	))
}
