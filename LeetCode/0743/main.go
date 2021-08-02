package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Info struct {
	v, time int
}

type Infos []*Info

func (a Infos) Len() int           { return len(a) }
func (a Infos) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Infos) Less(i, j int) bool { return a[i].time < a[j].time }
func (a *Infos) Pop() interface{} {
	l := a.Len()
	x := (*a)[l-1]
	old := *a
	*a = old[:l-1]
	return x

}
func (a *Infos) Push(x interface{}) {
	*a = append(*a, x.(*Info))
}

func networkDelayTime1(times [][]int, n int, k int) int {
	n += 1
	edges := make([][]Info, n)
	for _, v := range times {
		edges[v[0]] = append(edges[v[0]], Info{v: v[1], time: v[2]})
	}
	flag := make([]bool, n)
	queue := make([]int, 0, n*3)
	queue = append(queue, k)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = -1
	}
	h := 0
	flag[queue[h]] = true
	cost[queue[h]] = 0

	for h < len(queue) {
		v := queue[h]
		for _, edge := range edges[v] {
			if cost[edge.v] > cost[v]+edge.time || (cost[edge.v] == -1) {
				cost[edge.v] = cost[v] + edge.time
				queue = append(queue, edge.v)
				flag[edge.v] = true
			}
		}
		h += 1
		flag[v] = false
	}
	ans := 0
	for i := 1; i < n; i++ {
		if ans < cost[i] {
			ans = cost[i]
		} else if cost[i] == -1 {
			return -1
		}

	}
	return ans
}

func networkDelayTime2(times [][]int, n int, k int) int {
	n += 1
	edges := make([][]Info, n)
	for _, v := range times {
		edges[v[0]] = append(edges[v[0]], Info{v: v[1], time: v[2]})
	}
	flag := make([]bool, n)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	flag[k] = true
	cost[k] = 0

	for {
		p, minx := -1, math.MaxInt32
		for i := 1; i < n; i++ {
			if flag[i] && cost[i] < minx {
				p, minx = i, cost[i]
			}
		}
		if p == -1 {
			break
		}
		flag[p] = false
		for _, edge := range edges[p] {
			if cost[edge.v] > cost[p]+edge.time {
				cost[edge.v] = cost[p] + edge.time
				flag[edge.v] = true
			}
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if cost[i] == math.MaxInt32 {
			return -1
		}
		if ans < cost[i] {
			ans = cost[i]
		}
	}
	return ans
}

func networkDelayTime(times [][]int, n int, k int) int {
	n += 1
	edges := make([][]Info, n)
	for _, v := range times {
		edges[v[0]] = append(edges[v[0]], Info{v: v[1], time: v[2]})
	}
	cost := make([]int, n)
	queue := make(Infos, 0, n*3)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[k] = 0
	heap.Init(&queue)
	heap.Push(&queue, &Info{v: k, time: 0})

	for len(queue) > 0 {
		info := queue.Pop().(*Info)
		p := info.v
		if info.time > cost[p] {
			continue
		}
		for _, edge := range edges[p] {
			if cost[edge.v] > cost[p]+edge.time {
				cost[edge.v] = cost[p] + edge.time
				heap.Push(&queue, &Info{v: edge.v, time: cost[edge.v]})
			}
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if cost[i] == math.MaxInt32 {
			return -1
		}
		if ans < cost[i] {
			ans = cost[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(
		networkDelayTime(
			[][]int{
				{1, 2, 10},
				{1, 3, 2},
				{3, 2, 3},
				{3, 4, 1},
				{4, 2, 1},
			},
			4,
			1,
		),
	)
}
