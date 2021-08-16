package main

import "fmt"

func checkRecord(s string) bool {
	cntA, cntL := 0, 0
	for _, c := range s {
		switch c {
		case 'A':
			cntA++
			if cntA >= 2 {
				return false
			}
			cntL = 0
		case 'P':
			cntL = 0
		case 'L':
			cntL++
			if cntL >= 3 {
				return false
			}
		default:
			cntL = 0
		}
	}
	return true
}

func main() {
	fmt.Println(
		checkRecord("PPALLP"),
	)
}
