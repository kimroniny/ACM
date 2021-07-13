package ct1

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			sum := i*i + j*j
			if sum > n*n {
				break
			}
			a := int(math.Sqrt(float64(sum)))
			if a*a == sum {
				if i == j {
					ans += 1
				} else {
					ans += 2
				}
			}
		}
	}
	return ans
}

func Main() {
	fmt.Println(
		countTriples(1),
	)
}
