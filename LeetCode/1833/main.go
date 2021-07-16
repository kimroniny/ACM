package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for k, v := range costs {
		coins -= v
		if coins < 0 {
			return k
		}
	}
	return len(costs)
}
func main() {
	fmt.Println(
		maxIceCream(
			[]int{1, 6, 3, 1, 2, 5},
			20,
		),
	)
}
