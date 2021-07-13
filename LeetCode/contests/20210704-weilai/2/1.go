package ct2

import (
	"fmt"
	"math"
	"sort"
)

func calc(x, y int) int {
	return int(math.Ceil(float64(x) / float64(y)))
}

type info struct {
	dist  int
	speed int
	t int
}
type newlist []*info

func (I newlist) Len() int {
	return len(I)
}
func (I newlist) Less(i, j int) bool {
	return I[i].t < I[j].t
}
func (I newlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func eliminateMaximum(dist []int, speed []int) int {
	var infos newlist
	// mint := math.MaxInt32
	// mints := make([]int, len(dist))
	for v, k := range dist {
		infos = append(infos, &info{dist: k, speed: speed[v], t: calc(k, speed[v])})
	}
	sort.Sort(infos)
	ans := 0
	for k, v:=range infos {
		if k < v.t {
			ans += 1
		}else {
			break
		}
	}
	return ans
}

func Main() {
	fmt.Println(
		eliminateMaximum(
			[]int{1,1,2,3},
			[]int{1,1,1,1},
		),
	)
}
