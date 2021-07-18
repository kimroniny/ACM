package lanhu3

import "fmt"

func maxPoints(points [][]int) int64 {
	l := len(points)
	n := len(points[0])
	if l == 0 {
		return 0
	}
	dp := make([]int64, n)
	pre := make([]int64, n)
	last := make([]int64, n)
	for k := 0; k < l; k++ {
		if k == 0 {
			for j := 0; j < n; j++ {
				dp[j] = int64(points[k][j])
			}
			continue
		}

		temp := make([]int64, n)
		for j := 0; j < n; j++ {
			p := j
			temp[j] = dp[j] - int64(p)
		}
		last[n-1] = temp[n-1]
		for j := n - 2; j >= 0; j-- {
			last[j] = max(last[j+1], temp[j])
		}
		for j := 0; j < n; j++ {
			p := n - 1 - j
			temp[j] = dp[j] - int64(p)
		}
		pre[0] = temp[0]
		for j := 1; j < n; j++ {
			pre[j] = max(pre[j-1], temp[j])
		}

		for j := 0; j < n; j++ {
			dp[j] = max(
				pre[j]+int64(n-1-j),
				last[j]+int64(j),
			) + int64(points[k][j])
		}
	}
	for j := 0; j < n; j++ {
		dp[0] = max(dp[0], dp[j])
	}
	return dp[0]
}

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func Main() {
	fmt.Println(
		maxPoints(
			[][]int{
				{0, 3, 0, 4, 2},
				{5, 4, 2, 4, 1},
				{5, 0, 0, 5, 1},
				{2, 0, 1, 0, 3},
			},
		),
	)
}

// [[0,3,0,4,2],[5,4,2,4,1],[5,0,0,5,1],[2,0,1,0,3]]
