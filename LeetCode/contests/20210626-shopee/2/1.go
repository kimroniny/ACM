package contest2

import (
	"fmt"	
	"strings"
)


func removeOccurrences(s string, part string) string {
	for {
		if x := strings.Index(s, part); x > -1 {
			s = s[:x] + s[x+len(part):]
			continue
		}
		break
	}
	return s
}

func Main() {
	fmt.Println(
		removeOccurrences(
			"daabcbaabcbc",
			"abc",
		),
	)
	
}
