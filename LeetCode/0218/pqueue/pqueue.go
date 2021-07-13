package pqueue

import (
	heap "container/heap"
	"fmt"
)

var mapx map[int]int

type IntQueue []int

func (h IntQueue) Len() int           { return len(h) }
func (h IntQueue) Less(i, j int) bool { return h[i] > h[j] }
func (h IntQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	mapx[h[i]] = j
	mapx[h[j]] = i
}

func (h *IntQueue) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestPQueue() {
	mapx = make(map[int]int)
	elems := []int{2, 4, 3, 6, 1, 4}
	h := &IntQueue{}
	heap.Init(h)
	for k, v := range elems {
		mapx[v] = k
	}
	for _, v := range elems {
		h.Push(v)
	}
	fmt.Println(h)
	fmt.Println(mapx)
	// for h.Len() > 0 {
	// 	fmt.Println(heap.Pop(h))

	// }
}
