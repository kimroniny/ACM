package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle[len(triangle)-1]))
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == len(triangle)-1 {
				dp[j] = triangle[i][j]
			} else {
				dp[j] = Min(
					dp[j],
					dp[j+1],
				) + triangle[i][j]
			}
		}
	}
	return dp[0]
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(
		minimumTotal(
			[][]int{{-10},{2,3}},
		),
	)
}
