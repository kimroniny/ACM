package main

func shortestPathLength(graph [][]int) int {
	a := make([][]int, len(graph))
	flag := make([]bool, len(graph))
	dist := 0
	var dfs func(x int) int
	dfs = func(x int) int {
		flag[x] = true
		for _, v := range graph[x] {
			if flag[v] {
				continue
			}
			dist += 1
			dist += dfs(v)
			dist += 1
		}
	}
}

func main() {

}
