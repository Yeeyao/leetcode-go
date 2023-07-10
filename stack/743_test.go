package stack

import "container/heap"

/*
	给定网络的 n 个节点，标号从 1 到 n 同时给定次数 times 数组，times[i] = (ui, vi, wi) 表示从 ui 到 vi 节点需要的时间是 wi。
	我们会从给定节点 k 发送一个信号，返回所有的 n 个节点收到信号的所需的最小时间。如果不可能全都收到就返回 -1
	这里其实是计算消耗时间最长的路径
*/
func networkDelayTime(times [][]int, n int, k int) int {
	// 对应 dijkstra 开始是 k 结束是所有的节点，然后取其中的最大值
}

// graph[i][j] = c 表示从 i 到 j 需要 cost c
func dijkstra(graph [][]int, start, end int) int {
	// 先将开始节点放入
	h := &graphNodes{gnList: []*gn{}}
	// 从 start 到 start cost 0
	h.Push(&gn{
		pos:  start,
		cost: 0,
	})
	// 记录已经访问的节点
	visitedMap := make(map[int]interface{})
	for h.Len() > 0 {
		topNodes := heap.Pop(h).(*gn)
		pos := topNodes.pos
		cost := topNodes.cost
		// 已经遍历过就跳过
		if _, ok := visitedMap[pos]; ok {
			continue
		}
		visitedMap[pos] = struct{}{}
		// 如果已经遇到目标的节点了，直接返回
		if pos == end {
			return cost
		}
		for i, v := range graph[pos] {
			if _, ok := visitedMap[i]; ok {
				continue
			}
			heap.Push(h, &gn{
				pos:  i,
				cost: cost + v,
			})
		}
	}
	return -1
}

func dijkstraBackup(graph [][]int, start, end int) int {
	// 先将开始节点放入
	h := &graphNodes{gnList: []*gn{}}
	// 从 start 到 start cost 0
	h.Push(&gn{
		pos:  start,
		cost: 0,
	})
	// 记录已经访问的节点
	visitedMap := make(map[int]interface{})
	for h.Len() > 0 {
		topNodes := heap.Pop(h).(*gn)
		pos := topNodes.pos
		cost := topNodes.cost
		// 已经遍历过就跳过
		if _, ok := visitedMap[pos]; ok {
			continue
		}
		visitedMap[pos] = struct{}{}
		// 如果已经遇到最后的节点了，直接返回
		if pos == end {
			return cost
		}
		for _, v := range graph {
			// 找到当前节点开始的图
			if v[0] == pos {
				// 如果已经遍历过就跳过
				if _, ok := visitedMap[v[1]]; ok {
					continue
				}
				// 没有遍历过，开始到下一个节点的 cost 等于开始到当前节点的 cost 加上 当前节点到下一个节点的 cost
				heap.Push(h, &gn{
					pos:  v[1],
					cost: cost + v[2],
				})
			}
		}
	}
	return -1
}

type graphNodes struct {
	gnList []*gn
}

// 从 start 到 post 权重是 cost
type gn struct {
	pos  int
	cost int
}

func (g *graphNodes) Len() int {
	return len(g.gnList)
}

func (g *graphNodes) Less(i, j int) bool {
	return g.gnList[i].cost < g.gnList[j].cost
}

func (g *graphNodes) Swap(i, j int) {
	temp := g.gnList[i]
	g.gnList[i] = g.gnList[j]
	g.gnList[j] = temp
}

func (g *graphNodes) Push(v interface{}) {
	g.gnList = append(g.gnList, v.(*gn))
}

func (g *graphNodes) Pop() interface{} {
	top := g.gnList[len(g.gnList)-1]
	g.gnList = g.gnList[:len(g.gnList)-1]
	return top
}
