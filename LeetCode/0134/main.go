package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	ans := 0
	p := 0
	l := len(gas)
	for j := 0; j < l*2; j++ {
		i := j % l
		x := gas[i] - cost[i]
		ans += x
		if ans < 0 {
			if j >= l {
				break
			}
			ans = 0
			p = (j + 1) % l
		} else {
			if p > 0 && i == p-1 || p == 0 && i == l-1 {
				return p
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(
		canCompleteCircuit(
			[]int{2, 3, 4},
			[]int{3, 4, 3},
		),
	)
}
