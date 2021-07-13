package ct4

import (
	"fmt"
)

const MAXN int = 1001

var d [][][]int

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(edges)
	d := make([][][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = make([]int, 0, MAXN)
		}
	}
	for i, v := range edges {
		d[v[0]][v[1]] = append(d[v[0]][v[1]], v[2])
		d[v[1]][v[0]] = append(d[v[1]][v[0]], v[2])
	}
	start, end := 0, n-1
	q := make([]int, 0, MAXN*10)
	h := 0
	q = append(q, start)
	for h < len(q) {
		v := q[h]
		for i := 0; i < n; i++ {
			if len(d[v][i]) == 0 {
				continue
			}
			
		}
		h += 1
	}

}
func Main() {
	fmt.Println()
}
