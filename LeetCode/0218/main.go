package main

import (
	heap "container/heap"
	"fmt"
	"math"
	"sort"
)

var mapx map[int]int

const MININT int = math.MinInt16

type Info struct {
	idx, height, x int
}

type InfoSort []Info

func (a InfoSort) Len() int      { return len(a) }
func (a InfoSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a InfoSort) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].height > a[j].height
	}
	return a[i].x < a[j].x
}

type InfoHeap []Info

func (a InfoHeap) Len() int { return len(a) }
func (a InfoHeap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
	mapx[a[i].idx] = i
	mapx[a[j].idx] = j
}
func (a InfoHeap) Less(i, j int) bool {
	if a[i].height == a[j].height {
		return a[i].x < a[j].x
	}
	return a[i].height > a[j].height
}

func (h *InfoHeap) Push(x interface{}) {
	info := x.(Info)
	*h = append(*h, info)

	mapx[info.idx] = h.Len() - 1
}

func (h *InfoHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	mapx = make(map[int]int)
	infoList := InfoSort{}
	infoHeap := &InfoHeap{}
	heap.Init(infoHeap)
	for k, v := range buildings {
		infoList = append(infoList, Info{idx: k, x: v[0], height: v[2]})
		if v[2] == 0 {
			infoList = append(infoList, Info{idx: k, x: v[1], height: MININT})
		} else {
			infoList = append(infoList, Info{idx: k, x: v[1], height: -v[2]})
		}

	}
	sort.Sort(infoList)
	var ans [][]int
	for _, info := range infoList {
		if info.height >= 0 {
			heap.Push(infoHeap, info)
			if (*infoHeap)[0].idx == info.idx {
				tmp := []int{info.x, info.height}
				ans = append(ans, tmp)
			}
		} else {
			headId := mapx[info.idx]
			lastHeapId := (*infoHeap)[0].idx
			heap.Remove(infoHeap, headId)
			if (*infoHeap).Len() > 0 {
				idx := (*infoHeap)[0].idx
				if idx != lastHeapId && buildings[idx][2] != buildings[lastHeapId][2] {
					tmp := []int{info.x, (*infoHeap)[0].height}
					ans = append(ans, tmp)
				}
			}else{
				tmp := []int{info.x,0}
				ans = append(ans, tmp)
			}
			
		}
	}
	return ans
}

func main() {
	fmt.Println(
		getSkyline(
			[][]int{
				{2, 9, 10},
				{3, 7, 15},
				{5, 12, 12},
				{15, 20, 10},
				{19, 24, 8},
			},
		),
	)
}
