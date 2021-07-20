package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	i, j := 0, l-1
	ans := nums[i] + nums[j]
	for i < j {
		ans = max(ans, nums[i+1]+nums[j-1])
		i += 1
		j -= 1
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(minPairSum(
		[]int{3,5,4,2,4,6},
	))
}
