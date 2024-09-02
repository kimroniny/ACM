package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ans := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ans++
		}
	}
	return ans
}

func main() {
	startTime := []int{1, 2}
	endTime := []int{2, 3}
	fmt.Println(busyStudent(startTime, endTime, 2))
}
