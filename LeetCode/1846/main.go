package main

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if abs(arr[i]-arr[i-1]) <= 1 {
			continue
		}
		arr[i] = arr[i-1] + 1
	}
	return arr[len(arr)-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(
		maximumElementAfterDecrementingAndRearranging(
			[]int{100,1,1000},
		),
	)
}
