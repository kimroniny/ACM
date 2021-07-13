package ct3

import "fmt"

const MOD int = 1000000007

func mi(x int, y int64) int {
	var a int
	if y == 0 {
		return 1
	}
	if y&1 == 1 {
		a = mi(x, y/2)
		a *= a * x
	}else {
		a = mi(x, y/2)
		a *= a
	}
	return  a % MOD
}

func countGoodNumbers(n int64) int {
	return mi(5, (n+1)/2) * mi(4, n/2) % MOD
}

func Main() {
	fmt.Println(
		countGoodNumbers(50),
		// mi(4, 2),
	)
}
