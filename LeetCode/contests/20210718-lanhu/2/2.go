package lanhu2

import "fmt"

func addRungs(rungs []int, dist int) int {
	ans := 0
	rungs = append([]int{0}, rungs...)
	for k, v := range rungs[1:] {
		x := v - rungs[k]
		ans += x/dist
		if x%dist == 0 {
			ans -= 1
		}
	}
	return ans
}

func Main() {
	fmt.Println(
		addRungs(
			[]int{},
			1,
		),
	)
}
