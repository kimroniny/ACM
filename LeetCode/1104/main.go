package main

import (
	"fmt"
	"math"
)

func pathInZigZagTree(label int) []int {

	pows := make([]int, int(math.Log2(1000000))+2)
	pows[0] = 1
	for i := 1; i < len(pows); i++ {
		pows[i] = pows[i-1] << 1
	}
	// fmt.Println("ok")
	result := make([]int, 0, int(math.Log2(float64(label)))+2)
	var dfs func(k int)
	dfs = func(k int) {
		if k == 1 {
			result = append(result, k)
			return
		}
		row := int(math.Log2(float64(k))) + 1
		// fmt.Println(row)
		if row&1 == 1 {
			realk := pows[row] + pows[row-1] - 1 - k
			// fmt.Println("1: ", realk)
			dfs(realk / 2)
		} else {
			nextk := pows[row-1] + pows[row-2] - 1 - k/2
			// fmt.Println("0: ", nextk)
			dfs(nextk)
		}
		result = append(result, k)
	}
	dfs(label)
	return result
}

func main() {
	fmt.Println(pathInZigZagTree(26))
}
