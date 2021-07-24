package main

import (
	"fmt"
	"sort"
)

type Info struct {
	left, right int
}

type Infos []*Info

func (a Infos) Len() int      { return len(a) }
func (a Infos) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Infos) Less(i, j int) bool {
	if a[i].left == a[j].left {
		return a[i].right < a[j].right
	}
	return a[i].left < a[j].left
}

func isCovered(ranges [][]int, left int, right int) bool {
	infos := make(Infos, 0, len(ranges))
	for _, v := range ranges {
		infos = append(infos, &Info{left: v[0], right: v[1]})
	}
	sort.Sort(infos)
	// fmt.Println(infos)
	infoShort := make(Infos, 0, len(ranges))
	leftv, rightv := infos[0].left, infos[0].right
	for _, v := range infos {
		// fmt.Println(*v)
		if v.left >= leftv && v.left <= rightv {
			if rightv < v.right {
				rightv = v.right
			}
		} else {
			infoShort = append(infoShort, &Info{left: leftv, right: rightv})
			leftv = v.left
			rightv = v.right
		}
	}
	infoShort = append(infoShort, &Info{left: leftv, right: rightv})
	// fmt.Println(infoShort)
	leftp, rightp := -1, -1
	for k, v := range infoShort {
		// fmt.Println(*v)
		if left >= v.left && left <= v.right && leftp == -1 {
			leftp = k
		}
		if right >= v.left && right <= v.right && rightp == -1 {
			rightp = k
		}
	}
	// fmt.Println(leftp, rightp)
	if leftp == -1 || rightp == -1 {
		return false
	}
	for i := leftp; i < rightp; i++ {
		// fmt.Println(infoShort[i].right, infoShort[i+1].left)
		if infoShort[i].right < infoShort[i+1].left-1 {
			return false
		}
	}
	return true
}

func isCovered1(ranges [][]int, left int, right int) bool {
	maps := make(map[int]bool)
	for _, v := range ranges {
		for i := v[0]; i <= v[1]; i++ {
			maps[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !maps[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(
		isCovered(
			[][]int{
				{37, 49},
				{5, 17},
				{8, 32},
				// {25, 42},
				// {7, 14},
				// {2, 32},
				// {25, 28},
				// {39, 49},
				// {1, 50},
				// {29, 45},
				// {18, 47},
			},
			29,
			49,
		),
	)
}

// [[1,2],[3,4],[5,6]]
// 2
// 5
// [[25,42],[7,14],[2,32],[25,28],[39,49],[1,50],[29,45],[18,47]]
// 15
// 38
// [[37,49],[5,17],[8,32]]
// 29
// 49
