package main

import "fmt"

func countArrangement(n int) int {
	flag := make([]bool, n)
	var dfs func(x int) int
	dfs = func(x int) int {
		if x == 0 {
			return 1
		}
		k := 0
		for i := 1; i <= n; i++ {
			if !flag[i-1] && (x%i == 0 || i%x == 0) {
				flag[i-1] = true
				k += dfs(x - 1)
				flag[i-1] = false
			}
		}
		return k
	}
	return dfs(n)
}

func main() {
	fmt.Println(countArrangement(15))
}
