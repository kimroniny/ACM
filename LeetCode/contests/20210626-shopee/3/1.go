package contest3

import "fmt"

func maxAlternatingSum(nums []int) int64 {
	var dp [100050][2][2]int
	// dp[0][0] = 0
	dp[0][1][0] = 0
	dp[0][1][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0][0] = max(
			dp[i-1][1][0], dp[i-1][0][0],
		)
		dp[i][0][1] = max(
			dp[i-1][1][1], dp[i-1][0][1],
		)
		dp[i][1][0] = max(
			dp[i-1][1][1]-nums[i], dp[i-1][0][1]-nums[i],
		)
		dp[i][1][1] = max(
			dp[i-1][1][0]+nums[i], dp[i-1][0][0]+nums[i],
		)
	}
	// fmt.Println(dp[:len(nums)])
	ans := 0
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			ans = max(ans, dp[len(nums)-1][i][j])
		}
	}
	return int64(ans)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Main() {
	fmt.Println(
		maxAlternatingSum(
			[]int{1},
		),
	)
}
