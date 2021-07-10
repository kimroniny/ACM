package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	sign := true
	if n < 0 {
		sign = false
		n = -n
	}
	if n == 0 {
		return 1
	}
	mid := n / 2
	y := myPow(x, mid)
	y *= y
	if n&1 == 1 {
		y *= x
	}
	if !sign {
		return 1 / y
	} else {
		return y
	}
}

func main() {
	fmt.Println(myPow(1, -2))
	// fmt.Println(-1/2)
}
