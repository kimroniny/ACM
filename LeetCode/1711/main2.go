package main

import "fmt"

const MOD int = 1000000007

var mapi map[int]int

func countPairs(deliciousness []int) int {
	mapi = make(map[int]int)
	ans := 0
	for _, v := range deliciousness {
		for i := 0; i <= 21; i++ {
			x := (1 << i) - v
			if _, ok := mapi[x]; ok {
				ans = (ans + mapi[x]) % MOD
			}
		}
		if _, ok := mapi[v]; !ok {
			mapi[v] = 1
		} else {
			mapi[v] += 1
		}
	}

	return ans
}

func main() {
	fmt.Println(
		countPairs(
			[]int{4, 4},
		),
	)
}
