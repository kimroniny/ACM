package main

import (
	"container/heap"
	"sort"
)

type Info struct {
	idx int
	val int
}

type InfoHeap []*Info

func (a InfoHeap) Len() int      { return len(a) }
func (a InfoHeap) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a InfoHeap) Less(i, j int) bool {
	if a[i].val == a[j].val {
		return a[i].idx < a[j].idx
	}
	return a[i].val < a[j].val
}
func (h *InfoHeap) Push(x interface{}) {
	*h = append(*h, x.(*Info))
}

func (h *InfoHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kWeakestRows(mat [][]int, k int) []int {
	lenrow := len(mat[0])
	infoHeap := make(InfoHeap, 0, len(mat))
	heap.Init(&infoHeap)
	for idx, row := range mat {
		pos := sort.Search(lenrow, func(i int) bool { return row[i] == 0 })
		heap.Push(&infoHeap, &Info{idx: idx, val: pos})
	}
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		x := heap.Pop(&infoHeap).(*Info)
		result = append(result, x.idx)
	}
	return result
}

func main() {

}
