package main

import "fmt"

func insert1(intervals [][]int, newInterval []int) [][]int {
	x, y := newInterval[0], newInterval[1]
	var p, q, posp, posq int = -1, -1, -1, -1
	for k, v := range intervals {
		if v[0] <= x && v[1] >= x {
			p = v[0]
			posp = k
		}
		if v[0] > x && k == 0 {
			p = x
			posp = -1
		}
		if v[1] < x {
			if k == len(intervals)-1 || intervals[k+1][0] > x {
				p = x
				posp = k + 1
			}
		}
		if v[0] <= y && v[1] >= y {
			q = v[1]
			posq = k
		}
		if v[0] > y {
			if k == 0 || intervals[k-1][1] < y {
				q = y
				posq = k - 1
			}
		}
		if v[1] < y && k == len(intervals)-1 {
			q = y
			posq = len(intervals)
		}
	}
	fmt.Println(p, q, posp, posq)
	if posq == -1 {
		result := make([][]int, 0, len(intervals)+1)
		result = append(result, []int{x, y})
		for _, v := range intervals {
			result = append(result, v)
		}
		return result
	} else if posp == -1 {
		result := make([][]int, 0, len(intervals)+1)
		for _, v := range intervals {
			result = append(result, v)
		}
		result = append(result, []int{x, y})
		return result
	} else {
		result := make([][]int, 0, len(intervals)-(posq-posp))
		for k, v := range intervals {
			if k < posp || k > posq {
				result = append(result, v)
			} else {
				result = append(result, []int{p, q})
			}
		}
		return result
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	x, y := newInterval[0], newInterval[1]
	var p, q int = -1, -1
	intervals = append([][]int{{-1, -1}}, intervals...)
	intervals = append(intervals, []int{1000000, 1000000})
	result := make([][]int, 0, len(intervals))

	for i := 1; i < len(intervals); i++ {
		v := intervals[i]
		flag := false
		if p == -1 && v[0] <= x && v[1] >= x {
			p = v[0]
		}
		if q == -1 && v[0] <= y && v[1] >= y {
			q = v[1]
		}
		if v[0] > y && intervals[i-1][1] < x {
			p, q = x, y
			flag = true && (i != len(intervals)-1)
		}
		if p == -1 && v[0] > x && intervals[i-1][1] < x && y >= v[0] {
			p = x
		}
		if q == -1 && v[1] < y && x <= v[1] && (i < len(intervals)-1 && intervals[i+1][0] > y) {
			q = y
		}
		if p == -1 && q == -1 && i != len(intervals)-1 {
			result = append(result, v)
		}

		if p != -1 && q != -1 {
			result = append(result, []int{p, q})
			p, q = -1, -1
			if flag {
				result = append(result, v)
			}
		}

	}
	return result
}

func main() {
	fmt.Println(
		insert(
			[][]int{{1, 3}, {6, 9}, {10, 12}},
			[]int{7, 8},
		),
	)
}
