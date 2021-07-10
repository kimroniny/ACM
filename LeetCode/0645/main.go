package main

import "fmt"

var allnum []bool

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	allnum = make([]bool, len(nums)+1)
	for _, v := range nums {
		if allnum[v] {
			result[0] = v
		}
		allnum[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		if !allnum[i] {
			result[1] = i
			break
		}
	}
	return result
}

func main() {
	fmt.Println(
		findErrorNums(
			[]int{1,1},
		),
	)
}
