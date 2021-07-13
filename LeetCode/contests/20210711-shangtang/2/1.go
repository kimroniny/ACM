package ct2

import "fmt"

func countPalindromicSubsequence(s string) int {
	s = "." + s + "$"
	prev := make([][]int, len(s))
	suf := make([][]int, len(s))
	alp := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		alp[i] = make([]bool, 26)
	}
	for i := 0; i < len(s); i++ {
		prev[i] = make([]int, 26)
		suf[i] = make([]int, 26)
	}
	for k, v := range s {
		if k != 0 && k != len(s)-1 {
			for j := 0; j < 26; j++ {
				prev[k+1][j] = prev[k][j]
			}
			prev[k+1][v-'a'] += 1
		}
	}
	for i := len(s) - 2; i > 0; i-- {
		for j := 0; j < 26; j++ {
			suf[i-1][j] = suf[i][j]
		}
		suf[i-1][s[i]-'a'] += 1
	}
	ans := 0
	for i := 1; i < len(s)-1; i++ {
		for j := 0; j < 26; j++ {
			if prev[i][j] > 0 && suf[i][j] > 0 && !alp[s[i]-'a'][j] {
				ans += 1
				alp[s[i]-'a'][j] = true
				// fmt.Println(s[i]-'a', j)
			}
		}
	}
	return ans
}

func Main() {
	fmt.Println(countPalindromicSubsequence("aabca"))
}
