package main

import "fmt"

func checkRecord(n int) int {
	MOD := 1000000007
	f := make([][4][2]int, n)
	f[0][0][0] = 1
	f[0][1][0] = 1
	f[0][2][0] = 1
	for i := 1; i < n; i++ {
		// 'P'
		f[i][0][0] += (f[i-1][0][0] + f[i-1][2][0] + f[i-1][3][0]) % MOD
		f[i][0][1] += (f[i-1][1][0] + f[i-1][0][1] + f[i-1][2][1] + f[i-1][3][1]) % MOD
		// 'A'
		f[i][1][0] += (f[i-1][0][0] + f[i-1][2][0] + f[i-1][3][0]) % MOD
		f[i][1][1] = 0
		// 'XL'
		f[i][2][0] += f[i-1][0][0] % MOD
		f[i][2][1] += (f[i-1][0][1] + f[i-1][1][0]) % MOD
		// 'XLL'
		if i > 1 {
			f[i][3][0] += f[i-2][0][0] % MOD
			f[i][3][1] += (f[i-2][0][1] + f[i-2][1][0]) % MOD
		} else {
			f[i][3][0], f[i][3][1] = 1, 0
		}
	}
	ans := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			ans += f[n-1][i][j]
			ans %= MOD
		}
	}
	return ans
}

func main() {
	fmt.Println(
		checkRecord(100),
	)
}
