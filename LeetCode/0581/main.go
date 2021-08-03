package main

import (
	"fmt"
	"math"
	"sort"
)

func findUnsortedSubarray1(nums []int) int {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	// fmt.Println(nums, nums2)
	sort.Ints(nums2)
	// fmt.Println(nums, nums2)
	l, r := 0, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			l = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != nums2[i] {
			r = i
			break
		}
	}
	return r - l + 1
}

func findUnsortedSubarray(nums []int) int {
	flagl, flagr := -1,-1
	minx, maxx := math.MaxInt32, math.MinInt32
	for i := 1; i < len(nums); i++ {

		if flagl == -1 && nums[i] < nums[i-1] {
			flagl = i
		}
		if flagl != -1 {
			minx = min(minx, nums[i])
		}
		
	}
	if flagl == -1 {return 0}
	for i := len(nums)-2; i >= 0; i-- {
		if flagr == -1 && nums[i] > nums[i+1] {
			flagr = i
		}
		if flagr != -1 {
			maxx = max(maxx, nums[i])
		}
	}
	flagl, flagr = flagl-1, flagr+1
	for flagl >= 0 && minx < nums[flagl] {
		flagl -= 1
	}
	for flagr < len(nums) && maxx > nums[flagr] {
		flagr += 1
	}
	// fmt.Println(flagl, flagr, flagr-flagr-1)
	return flagr-flagl-1
}
func min(x,y int) int {
	if x < y {return x}
	return y
}
func max(x,y int) int {
	if x < y {return y}
	return x
}
func main() {
	fmt.Println(
		findUnsortedSubarray(
			[]int{2,3,4},
		),
	)
}
