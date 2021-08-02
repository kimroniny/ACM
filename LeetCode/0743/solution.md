三种方法来练手

### BFS
使用队列保存每一个将要访问的节点 `i`，并将节点 `i` 的标志位 `flag[i]` 置为 `false`, 入队时置为 `true`
```golang
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
```

### 朴素 Dijsktra
每次搜索 `cost[]` 的最小值及对应的节点，用该节点更新所连接的边的 `cost`，使用 `flag[]` 标记是否需要更新（'false' 不需要更新）
```golang
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
```

### 最小堆 Dijsktra
搜索 `cost[]` 最小值的过程用堆来实现。
```golang
for len(queue) > 0 {
    info := queue.Pop().(*Info)
    p := info.v
    if info.time > cost[p] {
        continue // 某个节点会存在多个值，因为push的时候并没有把原来的值删掉，所以从堆里面取出的时候判断一下是否是该节点 p 的最小值（即cost[p]）
    }
    for _, edge := range edges[p] {
        if cost[edge.v] > cost[p]+edge.time {
            cost[edge.v] = cost[p] + edge.time
            heap.Push(&queue, &Info{v: edge.v, time: cost[edge.v]})
        }
    }
}
```