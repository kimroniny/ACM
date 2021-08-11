package main

import (
	"fmt"
	"math"
)

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	calc := func(x, step int) int {
		return ((x+step)%n + n) % n

	}
	flag := -1000
	for i := 0; i < n; i++ {
		k := i
		flag -= 1
		for {
			fmt.Println(i, k, nums[k])
			p := calc(k, nums[k])
			if p == k {
				break
			} else {
				if int(math.Abs(float64(nums[k]))) < 1001 {
					if nums[k] > 0 {
						nums[k] = -flag
					} else {
						nums[k] = flag
					}
				} else if int(math.Abs(float64(nums[k]))) == -flag {
					fmt.Println("==")
					kk := calc(k, nums[k])
					fmt.Println(kk)
					sign := true
					for kk != k {
						if float32(nums[k])/float32(nums[kk]) < 0 {
							fmt.Println("--")
							sign = false
							break
						}
						kk = calc(kk, nums[kk])
						fmt.Println(kk)
					}
					if sign {
						fmt.Println(k)
						return true
					} else {
						break
					}

				} else {
					break
				}
			}
			k = p
		}
	}

	return false
}
func main() {
	fmt.Println(circularArrayLoop(
		[]int{2, -1, 1, 2, 2},
	))
}
