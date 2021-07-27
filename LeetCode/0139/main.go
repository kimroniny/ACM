package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	d := make([]bool, l)
	maps := make(map[string]bool)
	for _, v := range wordDict {
		maps[v] = true
	}
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			if maps[s[j:i+1]] {
				if j == 0 {
					d[i] = true
				} else {
					d[i] = d[i] || d[j-1]
				}

			}
		}
	}
	return d[l-1]
}

func main() {
	fmt.Println(
		wordBreak(
			"a",
			[]string{"a"},
		),
	)
}
