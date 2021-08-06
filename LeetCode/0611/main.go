package main

import (
	"fmt"
	"sort"
)

func triangleNumber1(nums []int) int {
	l := len(nums)
	ans := 0
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			p := sort.Search(
				l-i-1,
				func(x int) bool {
					return nums[x+i+1] >= nums[i]+nums[j]
				},
			)
			ans += p
		}
	}
	return ans
}

func triangleNumber(nums []int) int {
	l := len(nums)
	ans := 0
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		j, k := i+1, i+2
		for k < l {
			if nums[i]+nums[j] > nums[k] {
				ans += k - j
				k += 1
			} else {
				j += 1
				if j == k {
					k += 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(
		triangleNumber(
			[]int{48,66,61,46,94,75},
		),
	)
}
