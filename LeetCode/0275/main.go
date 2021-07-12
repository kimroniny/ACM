package main

import "fmt"

func hIndex1(citations []int) int {
	l, r := 0, len(citations)-1
	for l <= r {
		mid := (l + r) >> 1
		if len(citations)-mid < citations[mid] {
			r = mid - 1
		} else if len(citations)-mid == citations[mid] {
			return len(citations) - mid
		} else {
			l = mid + 1
		}
	}
	return len(citations) - r - 1
}

func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) >> 1
		if n-mid < citations[mid] {
			r = mid - 1
		} else if n-mid == citations[mid] {
			return n - mid
		} else {
			l = mid + 1
		}
	}
	return n - r - 1
}

func main() {
	fmt.Println(hIndex([]int{0}))
}
