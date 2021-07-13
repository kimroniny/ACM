package ct2

import "fmt"

func getP(l, a, b int) (int, int) {
	if l == 0 {
		return 0, 1
	}
	if l <= a {
		return l - 1, 0
	}
	if l > a && l <= a+b-1 {
		return a - 1, l - a
	}
	if l > a+b-1 && l <= a+b-1+a-1 {
		return a - 1 - (l - (a + b - 1)), b - 1
	}
	return 0, b - 1 - (l - (a + b - 1 + a - 1))
}

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	result := make([][]int, m)
	for i:=0; i < m; i++ {
		result[i] = make([]int, n)
	}
	// result := [m][n]int
	for p := 0; (p < m/2) && (p < n/2); p++ {
		a, b := m-p*2, n-p*2
		l := a*2 + b*2 - 4
		epoch := k % l
		
		for i := 0; i < a; i++ {
			x, y := getP((i+1+epoch)%(l), a, b)
			result[x+p][y+p] = grid[i+p][p]
			x, y = getP((a-i+a+b-2+epoch)%(l), a, b)
			result[x+p][y+p] = grid[i+p][b-1+p]
		}
		for i := 0; i < b; i++ {
			x, y := getP((a+b-1+a-1+b-1-i+epoch)%(l), a, b)
			result[x+p][y+p] = grid[p][i+p]
			x, y = getP((a+i+epoch)%(l), a, b)
			result[x+p][y+p] = grid[a-1+p][i+p]
		}

	}
	return result
}
func Main() {
	fmt.Println(
		rotateGrid(
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			2,
		),
	)
}
