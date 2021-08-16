package main

import "fmt"

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	MOD := 1000000007
	d := make([][][]int, m)
	for idx := range d {
		d[idx] = make([][]int, n)
		for id2 := range d[idx] {
			d[idx][id2] = make([]int, maxMove+1)
			for i:= 0; i < len(d[idx][id2]); i++ {
				d[idx][id2][i] = -1
			}
		}
	}
	var dfs func(x, y, rest int) int
	dfs = func(x, y, rest int) int {
		if x >= m || x < 0 || y >= n || y < 0 {
			return 1
		}
		if rest == 0 {
			return 0
		}
		if d[x][y][rest] > -1 {
			return d[x][y][rest]
		}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i+j == 1 || i+j == -1 {
					k := dfs(x+i, y+j, rest-1)
					if d[x][y][rest] == -1 { // 如果不单独判断的话，那么对于无法在rest步骤内走到边界外的(x,y)，其值总是-1
						d[x][y][rest] = k % MOD
					}else {
						d[x][y][rest] = (d[x][y][rest]+k)%MOD
					}
				}
			}
		}
		return d[x][y][rest]
	}
	return dfs(startRow, startColumn, maxMove)
}

func main() {
	fmt.Println(
		findPaths(45, 35, 47, 20, 31),
	)
}
