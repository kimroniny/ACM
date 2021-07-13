package ct1

import (
	"fmt"
)

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for v, k := range nums {
		ans[v] = nums[k]
	}
	return ans	
}

func Main() {
	fmt.Println(
		buildArray(
			[]int{5,0,1,2,3,4},
		),
	)
}
