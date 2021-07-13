package ct2

import (
	"fmt"
	"math"
)

const MAXINT int = math.MaxInt16

type Point struct {
	x, y int
}

var queue []*Point
var d [][]int
var inqueue [][]bool

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	// fmt.Println(m, n)
	queue = make([]*Point, 0, m*n*3)
	d = make([][]int, m)
	inqueue = make([][]bool, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		inqueue[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			d[i][j] = MAXINT
		}

	}
	x, y := entrance[0], entrance[1]
	d[x][y] = 0
	queue = append(queue, &Point{x: x, y: y})
	h := 0
	for len(queue) > h {
		x, y = queue[h].x, queue[h].y
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if int(math.Abs(float64(j))) == int(math.Abs(float64(i))) {
					continue
				}
				// fmt.Println(i, j)
				p, q := x+i, y+j
				if (0 <= p && p < m && q >= 0 && q < n) && maze[p][q] == '.' {
					if d[p][q] > d[x][y]+1 {
						d[p][q] = d[x][y] + 1
						if !inqueue[p][q] {
							queue = append(queue, &Point{x: p, y: q})
							inqueue[p][q] = true
						}
					}
				}
			}
		}
		h += 1
		inqueue[x][y] = false
	}
	// fmt.Println(d)
	ans := MAXINT
	for i := 0; i < m; i++ {
		if maze[i][0] == '.' && !(i == entrance[0] && entrance[1] == 0) {
			ans = min(ans, d[i][0])
		}
		if maze[i][n-1] == '.' && !(i == entrance[0] && n-1 == entrance[1]) {
			ans = min(ans, d[i][n-1])
		}
	}
	for i := 0; i < n; i++ {
		if maze[0][i] == '.' && !(i == entrance[1] && entrance[0] == 0) {
			ans = min(ans, d[0][i])
		}
		if maze[m-1][i] == '.' && !(i == entrance[1] && entrance[0] == m-1) {
			ans = min(ans, d[m-1][i])
		}
	}
	if ans == MAXINT {
		return -1
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Main() {
	fmt.Println(
		nearestExit(
			[][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			[]int{1, 0},
		),
	)
}
