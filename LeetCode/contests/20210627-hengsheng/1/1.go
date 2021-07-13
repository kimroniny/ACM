package ct1

import (
	"fmt"
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}

func Main() {
	fmt.Println(
		maxProductDifference(
			[]int{4, 2, 5, 9, 7, 4, 8},
		),
	)
}
