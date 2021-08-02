package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	maxv, cnt := dp[0], 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxv {
			maxv = dp[i]
			cnt = 1
		} else if dp[i] == maxv {
			cnt += 1
		}
	}
	fmt.Println(dp)
	return cnt
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(
		findNumberOfLIS(
			[]int{1, 3, 5, 4, 7},
		),
	)
}
