package main

import (
	"fmt"
	"math"
)

func countDigitOne(n int) int {
	d := make([]int, 10)
	d[0] = 0
	for i := 1; i < len(d); i++ {
		d[i] = 10*d[i-1] + int(math.Pow(10, float64(i-1)))
	}
	// fmt.Println("d: ", d)
	a := make([]int, 0, 20)
	for n > 0 {
		a = append(a, n%10)
		n /= 10
	}
	// fmt.Println("a: ", a)
	calc := func(x int) int {
		ans, p := 0, 1
		for i := 0; i <= x; i++ {
			ans += p * a[i]
			p *= 10
		}
		return ans
	}
	var dfs func(x int) int
	dfs = func(x int) int {
		if x < 0 {
			return 0
		}
		y := a[x]
		cnt := 0
		if y > 1 {
			cnt = int(math.Pow(10, float64(x)))
		} else if y == 1 {
			cnt = calc(x - 1)+1
		} else {
			cnt = 0
		}
		return a[x]*d[x] + cnt + dfs(x-1)
	}
	return dfs(len(a) - 1)
}

func main() {
	fmt.Println(countDigitOne(11))
}
