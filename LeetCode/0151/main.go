package main

import "strings"

func reverseWords(s string) string {
	strs := []string{}
	for _, v := range strings.Split(s, " ") {
		if v == " " || v == "" {
			continue
		}
		strs = append(strs, v)
	}
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	
	return strings.Join(strs, " ")
}

func main() {

}
