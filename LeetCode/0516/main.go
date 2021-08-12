package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	n := len(s)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i == 0 || j == n-1 {
					d[i][j] = 1
				} else {
					d[i][j] = d[i-1][j+1] + 1
				}
			} else {
				if i == 0 && j != n-1 {
					d[i][j] = d[i][j+1]
				} else if i != 0 && j == n-1 {
					d[i][j] = d[i-1][j]
				} else if i == 0 && j == n-1 {
					d[i][j] = 0
				} else {
					d[i][j] = max(d[i-1][j], d[i][j+1])
				}
			}
		}
	}
	// fmt.Println(d)
	return d[n-1][0]
}

func main() {
	fmt.Println(longestPalindromeSubseq(
		"cbbd",
	))
}
