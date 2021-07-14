package main

import (
	"fmt"
	"math"
	"sort"
)

const MOD int = 1000000007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	maxx, sum, pos := 0, 0, -1
	for k, v := range nums1 {
		x := int(math.Abs(float64(v - nums2[k])))
		sum = (sum + x) % MOD
	}
	sort.Ints(nums1)
	for k, v := range nums2 {
		l := 0
	}

}

func main() {
	fmt.Println(
		minAbsoluteSumDiff(
			[]int{1, 10, 4, 4, 2, 7},
			[]int{9, 3, 5, 1, 7, 4},
		),
	)
}
