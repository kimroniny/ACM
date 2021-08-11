package main

import "fmt"

func tribonacci(n int) int {
	calc := func(x int) int {
		i := 3
		a, b, c := 0, 1, 1
		for i <= x {
			a, b, c = b, c, a+b+c
			i++
		}
		return c
	}
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return calc(n)
	}
}
func main() {
	fmt.Println(tribonacci(25))
}
