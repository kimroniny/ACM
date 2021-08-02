// 理解 heap
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] < a[j] }
func (a *IntHeap) Pop() interface{} {
	l := a.Len()
	old := *a
	x := old[l-1]
	*a = old[0 : l-1]
	return x
}
func (a *IntHeap) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func main() {
	n := 10
	infos := make(IntHeap, 0, n)
	heap.Init(&infos)
	for i := 0; i < n; i++ {
		heap.Push(&infos, i)
		fmt.Println(i, ": ", infos)
	}
	for i := 0; i < n; i++ {
		_ = ((heap.Pop(&infos))).(int)
		fmt.Println(infos)
	}

}
