package main

import "fmt"

func partition(s string) [][]string {
	l := len(s)
	sr := make([]rune, l)
	for i := 0; i < l; i++ {
		sr[l-1-i] = rune(s[i])
	}
	p := make([]string, 0)
	ans := make([][]string, 0)
	var dfs func(x int)
	dfs = func(x int) {
		if x == l {
			temp := make([]string, len(p))
			copy(temp, p)
			ans = append(ans, temp)
			return
		}
		for i := x; i < l; i++ {
			k := s[x : i+1]
			k1 := string(sr[l-i-1 : l-x])
			if k == k1 {
				p = append(p, k)
				dfs(i + 1)
				p = p[:len(p)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(
		partition("aab"),
	)
}
