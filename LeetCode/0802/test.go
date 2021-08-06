package main

import "fmt"

func main() {
	a := make([]int, 0, 10)
	for i := 0; i < cap(a); i++ {
		a = append(a, i)
	}
	b := a[2:4]
	fmt.Println(&a[0], &a[2], &b[0])
}
