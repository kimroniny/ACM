package main

import (
	"fmt"
)

func eventualSafeNodes1(graph [][]int) []int {
	var n int
	var in []int
	var result []int
	var edges [][]int
	var flag []bool
	n = len(graph)
	in = make([]int, n)
	edges = make([][]int, n)
	flag = make([]bool, n)
	for i := 0; i < n; i++ {
		k := len(graph[i])
		for j := 0; j < k; j++ {
			edges[graph[i][j]] = append(edges[graph[i][j]], i)

		}
		in[i] = k
	}
	for {
		f := false
		for i := 0; i < n; i++ {
			if in[i] == 0 && !flag[i] {
				for _, edge := range edges[i] {
					in[edge]++
				}
				flag[i] = true
				f = true
			}
		}
		if !f {
			break
		}
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			result = append(result, i)
		}
	}

	// sort.Ints(result)
	return result
}

func eventualSafeNodes(graph [][]int) []int {
	var n int
	var in []int
	var result []int
	var edges [][]int
	n = len(graph)
	in = make([]int, n)
	edges = make([][]int, n)
	for i := 0; i < n; i++ {
		k := len(graph[i])
		for j := 0; j < k; j++ {
			edges[graph[i][j]] = append(edges[graph[i][j]], i)

		}
		in[i] = k
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	h := 0
	for h < len(q) {
		for _, edge := range edges[q[h]] {
			in[edge]--
			if in[edge] == 0 {
				q = append(q, edge)
			}
		}
		h += 1
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			result = append(result, i)
		}
	}

	// sort.Ints(result)
	return result
}

func main() {
	fmt.Println(
		eventualSafeNodes(
			[][]int{
				{1, 2},
				{2, 3},
				{5},
				{0},
				{5},
				{},
				{},
			},
		),
	)
}
