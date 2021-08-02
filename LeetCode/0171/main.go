package main

import "fmt"

func titleToNumber(columnTitle string) int {
	p := 1
	ans := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		ans += (int(columnTitle[i]-'A') + 1) * p
		p *= 26
	}
	return ans
}

func main() {
	fmt.Println(
		titleToNumber("ZY"),
	)
}
