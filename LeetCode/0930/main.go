package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
	l, r, ans := 0, 0, 0
	sum := nums[0]
	for l < len(nums) && r < len(nums) {

		if sum == goal {
			ans += 1
			p := r + 1
			s2 := sum
			for p < len(nums) {
				s2 += nums[p]
				if s2 == goal {
					ans += 1
					p += 1
				} else if s2 < goal {
					p += 1
				} else {
					break
				}
			}
			sum -= nums[l]
			l += 1
			if l > r {
				r += 1
				if r < len(nums) {
					sum += nums[r]
				}
			}
		} else if sum > goal {
			sum -= nums[l]
			l += 1
			if l > r {
				r += 1
				if r < len(nums) {
					sum += nums[r]
				}
			}
		} else {
			r += 1
			if r < len(nums) {
				sum += nums[r]
			}
		}

	}
	return ans
}
func main() {
	fmt.Println(
		numSubarraysWithSum(
			[]int{1, 0, 1, 0, 1},
			1,
		),
	)
}
