package main

import "fmt"

func maximumTime(time string) string {
	a, b, c, d := time[0], time[1], time[3], time[4]
	if a == '?' {
		a = '2'
		if b >= '4' && b <= '9' {
			a = '1'
		}
	}
	if b == '?' {
		if a == '1' || a == '0' {
			b = '9'
		} else {
			b = '3'
		}
	}
	if c == '?' {
		c = '5'
	}
	if d == '?' {
		d = '9'
	}
	return string([]byte{a, b, ':',c, d})
}

func main() {
	fmt.Println(maximumTime("?6:??"))
}
