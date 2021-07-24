package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	zheng := math.MaxInt32
	fu := math.MinInt32
	ans := 0
	pro := 1
	flag := false
	for i := 0; i < len(nums); i++ {
		pro *= nums[i]
		if pro > 0 {
			ans = max(ans, pro)

		} else if pro < 0 {
			flag = !flag
			if !flag {}

		} else {
			zheng = 0
			fu = 0
			ans = max(ans, 0)
		}
		fmt.Println(ans)
	}
	return ans
}

func main() {
	fmt.Println(
		maxProduct(
			[]int{2, 3, -2, 4},
		),
	)
}
