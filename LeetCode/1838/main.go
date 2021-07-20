package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, ak int) int {
	sort.Ints(nums)
	ans, l, p, rest := 1, 0, 0, ak
	for p < len(nums)-1 {
		diff := nums[p+1] - nums[p]
		for true {
			if rest >= (p-l+1)*diff {
				ans = max(ans, p-l+1+1)
				rest -= (p - l + 1) * diff
				break
			}
			d := nums[p] - nums[l] + rest
			k := p - l
			if p == l {
				k = 1
			}
			if d >= diff*k {
				ans = max(ans, k+1)
				if k != 1 {
					l += 1
				}
				rest = d - k*diff
				break
			}
			rest += nums[p] - nums[l]
			l += 1
		}
		p += 1
	}
	return ans
}

func maxFrequency2(nums []int, ak int) int {
	sort.Ints(nums)
	ans, l, p, rest := 1, 0, 0, ak
	for p < len(nums)-1 {
		diff := nums[p+1] - nums[p]
		for {
			if rest >= (p-l+1)*diff {
				ans = max(ans, p-l+1+1)
				rest -= (p - l + 1) * diff
				p += 1
				break
			} else {
				d := nums[p] - nums[l] + rest
				k := p - l
				if k == 0 {
					k = 1
				}
				if d >= diff*k {
					ans = max(ans, k+1)
					rest = d - k*diff
					p += 1
					l += 1
					break
				} else {
					rest += nums[p] - nums[l]
					l += 1
				}
			}
		}
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
	fmt.Println(
		maxFrequency(
			[]int{9930,9923,9983,9997,9934,9952,9945,9914,9985,9982,9970,9932,9985,9902,9975,9990,9922,9990,9994,9937,9996,9964,9943,9963,9911,9925,9935,9945,9933,9916,9930,9938,10000,9916,9911,9959,9957,9907,9913,9916,9993,9930,9975,9924,9988,9923,9910,9925,9977,9981,9927,9930,9927,9925,9923,9904,9928,9928,9986,9903,9985,9954,9938,9911,9952,9974,9926,9920,9972,9983,9973,9917,9995,9973,9977,9947,9936,9975,9954,9932,9964,9972,9935,9946,9966},
			3056,
			// []int{1, 3, 5, 7, 100, 301, 303, 305, 307, 309},
			// []int{1, 3, 5, 7, 100, 301},
			// 8,
		),
	)
}

// [9930,9923,9983,9997,9934,9952,9945,9914,9985,9982,9970,9932,9985,9902,9975,9990,9922,9990,9994,9937,9996,9964,9943,9963,9911,9925,9935,9945,9933,9916,9930,9938,10000,9916,9911,9959,9957,9907,9913,9916,9993,9930,9975,9924,9988,9923,9910,9925,9977,9981,9927,9930,9927,9925,9923,9904,9928,9928,9986,9903,9985,9954,9938,9911,9952,9974,9926,9920,9972,9983,9973,9917,9995,9973,9977,9947,9936,9975,9954,9932,9964,9972,9935,9946,9966]
// 3056
