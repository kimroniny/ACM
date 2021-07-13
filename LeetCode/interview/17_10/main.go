package main

import "fmt"

func majorityElement1(nums []int) int {
	ms := make(map[int]int)
	for _, v := range nums {
		if _, ok := ms[v]; !ok {
			ms[v] = 0
		}
		ms[v] += 1
	}
	for k, v := range ms {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

func majorityElement(nums []int) int {
	x := nums[0]
	cnt := 0
	for _, v := range nums {
		if x == v {
			cnt += 1
		} else {
			if cnt == 0 {
				x = v
			}else {
				cnt -=1 
			}
		}
	}
	// fmt.Println(x)
	cnt = 0
	for _, v := range nums {
		if v == x {
			cnt += 1
		}
	}
	if cnt > len(nums)/2 {return x}
	return -1
}

func main() {
	fmt.Println(
		majorityElement(
			[]int{1,2,5,9,5,9,5,5,5},
		),
	)
}
