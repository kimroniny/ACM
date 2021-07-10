package main

import (
	"fmt"
)

func jump1(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for k, _ := range nums {
		for i := k - 1; i >= 0; i-- {
			if nums[i] >= k-i && dp[i] >= 0 {
				x := dp[i] + 1
				if dp[k] == -1 {
					dp[k] = x
				} else if dp[k] > x {
					dp[k] = x
				}
			}
		}
	}
	return dp[len(dp)-1]
}

func jump2(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for k, _ := range nums {
		for i := 0; i < k; i++ {
			if nums[i] >= k-i && dp[i] >= 0 {
				dp[k] = dp[i] + 1
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func jump(nums []int) int {
	dp := make([]int, len(nums))
	q := make([]int, 0, len(nums))
	dp[0] = 0
	q = append(q, 0)
	h := 0
	for k, v := range nums {
		max := v + k
		if k > q[h] {
			h += 1
		}
		if q[len(q)-1] < max {
			q = append(q, max)
			if max >= len(nums)-1 {
				if len(nums) == 1 {
					return 0
				}
				return dp[q[h]] + 1
			}
			dp[max] = min(dp[q[h]]+1, dp[max])
		}
	}

	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(
		jump(
			[]int{2, 9, 6, 5, 7, 0, 7, 2, 7, 9, 3, 2, 2, 5, 7, 8, 1, 6, 6, 6, 3, 5, 2, 2, 6, 3},
		),
	)
}

// 0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5
// 2,9,6,5,7,0,7,2,7,9,3,2,2,5,7,8,1,6,6,6,3,5,2,2,6,3
// 0,  1
//     1,              2
//     1           2
// 				2
// 		               3

//                           3
// 			                  3
// 			                        3
// 				                        4
// 							                  4
// 0 0 1 0 0 0 0 0 0 0 2 3 0 3 0 3 0 0 3 0 0 4 0 4 4 0
