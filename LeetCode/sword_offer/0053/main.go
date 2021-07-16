package main

import "fmt"

func search(nums []int, target int) int {
	ans := 0
	for _,v := range nums {
if v == target {
	ans += 1
}
	}
	return ans
}
func main() {
	fmt.Println(
		search(
			[]int{5, 7, 7, 8, 8, 10},
			8,
		),
	)
}
