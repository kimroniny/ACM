package ct3

import "fmt"

func colorTheGrid(m int, n int) int {
	d := make([][][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = make([]int, 3)
		}
	}
	for i := 0; i < 3; i++ {
		d[0][0][i] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for kp := 0; kp < 3; kp++ {
				flag := false
				for ki := 0; ki < 3; ki++ {
					if kp != ki {
						if i != 0 {
							d[i][j][kp] += d[i-1][j][ki]
						}
						if i == 0 && flag {
							continue
						}
						if j == 0 {
							continue
						}
						for kj := 0; kj < 3; kj++ {
							flag = true
							if kp != kj {
								d[i][j][kp] += d[i][j-1][kj]
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(d)
	return d[m-1][n-1][0] + d[m-1][n-1][1] + d[m-1][n-1][2]
}

func Main() {
	fmt.Println(colorTheGrid(2, 2))
}
