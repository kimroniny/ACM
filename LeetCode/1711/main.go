package main

import "fmt"

const MOD int = 1000000007

type Info struct {
	x   int
	cnt int
}

var mapi map[int]*Info

func countPairs(deliciousness []int) int {
	mapi = make(map[int]*Info)
	for _, v := range deliciousness {
		if _, ok := mapi[v]; !ok {
			mapi[v] = &Info{x: 1, cnt: 1}
		} else {
			mapi[v].x += 1
			mapi[v].cnt += 1
		}
	}
	ans := 0
	for _, v := range deliciousness {
		for i := 0; i <= 21; i++ {
			x := (1 << i) - v
			if v > x {
				continue
			}
			if x == v {
				if mapi[x].cnt > 0 {
					mapi[x].cnt -= 1
				}
				ans = (ans + mapi[x].cnt) % MOD
			} else {
				if _, ok := mapi[x]; ok {
					ans = (ans + mapi[x].x) % MOD
				}
			}

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
