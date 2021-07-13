package contest1

import (
	"fmt"
)


func canBeIncreasing(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	// flag := false
	for i := 0; i < len(nums); i++ {
		var nums1 []int
		nums1 = append(nums1, nums[:i]...)
		nums1 = append(nums1, nums[i+1:]...)
		// fmt.Println(nums1)
		var j int
		for j = 1; j < len(nums1); j++ {
			if nums1[j] <= nums1[j-1] {
				break
			}
		}
		if j == len(nums1) {
			return true
		}

	}

	return false
}

func Main() {
	// var arr []int = {1,2,3}
	fmt.Println(
		canBeIncreasing([]int{1, 2, 3}),
	)
}
