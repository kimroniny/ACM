package main

import "math"

func maxSubArray(nums []int) int {
	ans := math.MinInt16
	sum := 0
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		ans = max(sum, ans)
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

}
