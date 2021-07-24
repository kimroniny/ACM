package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	if len(nums) == 2 {
		if nums[0] >= nums[1] {return false}
		return true
	}
	nums = append([]int{0}, nums...)
	for i := 1; i < len(nums); i++ {
		last := nums[0]
		flag := true
		for j := 1; j < len(nums); {
			if j == i {
				j += 1
				continue
			}
			if nums[j] <= last {
				flag = false
				break
			}
			last, j = nums[j], j+1
		}
		if flag {
			// fmt.Println(i)
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(
		canBeIncreasing(
			[]int{4, 3},
		),
	)
}
