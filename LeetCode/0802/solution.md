这道题主要是理解题意. 无论走哪条路都能到达终点(没有出度的点), 所以如果某个点处于环中肯定是不可以的. 解决有向环的话, 肯定使用拓扑排序了.

终点即没有出度的点, 而拓扑排序面向入度为0的点, 所以先建立反向边. 

求解拓扑排序用队列, 把每个入度为零的入队, 不要遍历每个顶点, 然后判断入读是否为0
```golang
func eventualSafeNodes(graph [][]int) []int {
	var n int
	var in []int
	var result []int
	var edges [][]int
	n = len(graph)
	in = make([]int, n)
	edges = make([][]int, n)
	for i := 0; i < n; i++ {
		k := len(graph[i])
		for j := 0; j < k; j++ {
			edges[graph[i][j]] = append(edges[graph[i][j]], i)

		}
		in[i] = k
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	h := 0
	for h < len(q) {
		for _, edge := range edges[q[h]] {
			in[edge]--
			if in[edge] == 0 {
				q = append(q, edge)
			}
		}
		h += 1
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}
```