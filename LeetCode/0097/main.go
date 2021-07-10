package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	f := make([]bool, len(s2)+1)
	f[0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				f[j] = f[j] && (s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				f[j] = f[j] || (f[j-1] && (s2[j-1] == s3[i+j-1]))
			}
		}
	}
	// fmt.Println(f)
	return f[len(f)-1]
}

func main() {
	fmt.Println(
		isInterleave(
			"db",
			"b",
			"cbb",
		),
	)
}
