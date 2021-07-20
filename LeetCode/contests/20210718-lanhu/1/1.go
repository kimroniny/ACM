package lanhu1

import "fmt"

func canBeTypedWords(text string, brokenLetters string) int {
	b := make([]bool, 26)
	for _, v := range brokenLetters {
		// fmt.Println(int(v)-int('a'))
		b[int(v)-int('a')] = true
	}
	flag := true
	ans := 0
	text += " "
	for _, v := range text {
		if v == ' ' {
			if flag {
				ans += 1
			}
			flag = true
		} else {
			if b[int(v)-int('a')] {
				flag = false
			}
		}
	}
	return ans
}
func Main() {
	fmt.Println(
		canBeTypedWords(
			"hello world",
			"ad",
		),
	)
}
