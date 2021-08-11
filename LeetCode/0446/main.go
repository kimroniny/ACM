package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			dp[i][d] += dp[j][d] + 1
			ans += dp[j][d]
		}
	}
	return ans
}

func main() {
	fmt.Println(
		numberOfArithmeticSlices(
			// []int{2, 4, 6, 8, 10},
			[]int{7, 7, 7, 7, 7},
		),
	)
}
