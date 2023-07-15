package stack

import (
	"container/heap"
)

/*
	给定网络的 n 个节点，标号从 1 到 n 同时给定次数 times 数组，times[i] = (ui, vi, wi) 表示从 ui 到 vi 节点需要的时间是 wi。
	我们会从给定节点 k 发送一个信号，返回所有的 n 个节点收到信号的所需的最小时间。如果不可能全都收到就返回 -1
	这里其实是计算消耗时间最长的最短路径
*/

func networkDelayTimeOpt(times [][]int, n int, k int) int {
	// 需要构造 graph
	graph := make(map[int][]*gn)
	for _, v := range times {
		start := v[0]
		next := v[1]
		cost := v[2]
		if _, ok := graph[start]; ok {
			graph[start] = append(graph[start], &gn{
				pos:  next,
				cost: cost,
			})
		} else {
			graph[start] = []*gn{
				{next, cost},
			}
		}
	}
	// 这里只需要执行一次，遇到不可达的节点这里也会继续执行，所以和原来的算法对比需要看不可达节点的情况
	dist := dijkstraOpt(graph, k, n)
	// 表示有节点不能访问
	if len(dist) < n {
		return -1
	}

	maxMinWeight := -1
	for _, v := range dist {
		if v > maxMinWeight {
			maxMinWeight = v
		}
	}
	return maxMinWeight
}

// 这里就是直接从开始节点将整个图进行遍历，将所有节点如果能通的就记录下最小 cost
func dijkstraOpt(graph map[int][]*gn, start, end int) map[int]int {
	// 先将开始节点放入
	h := &graphNodes{gnList: []*gn{}}
	// 从 start 到 start cost 0
	h.Push(&gn{
		pos:  start,
		cost: 0,
	})
	dist := make(map[int]int, 0)
	// 记录已经访问的节点
	for h.Len() > 0 {
		topNodes := heap.Pop(h).(*gn)
		pos := topNodes.pos
		cost := topNodes.cost
		// 已经遍历过就跳过，因为 pop 出来的是最小的
		if _, ok := dist[pos]; ok {
			continue
		}
		dist[pos] = cost
		// 如果已经遇到目标的节点了，直接返回
		for _, next := range graph[pos] {
			if _, ok := dist[next.pos]; ok {
				continue
			}
			heap.Push(h, &gn{
				pos:  next.pos,
				cost: cost + next.cost,
			})
		}
	}
	return dist
}

func networkDelayTime(times [][]int, n int, k int) int {
	// 需要构造 graph
	graph := make(map[int][]*gn)
	for _, v := range times {
		start := v[0]
		next := v[1]
		cost := v[2]
		if _, ok := graph[start]; ok {
			graph[start] = append(graph[start], &gn{
				pos:  next,
				cost: cost,
			})
		} else {
			graph[start] = []*gn{
				{next, cost},
			}
		}
	}

	maxMinWeight := -1
	for i := 1; i <= n; i++ {
		if i == k {
			continue
		}
		weight := dijkstra(graph, k, i)
		if weight == -1 {
			return -1
		}
		if weight > maxMinWeight {
			maxMinWeight = weight
		}
	}
	return maxMinWeight
	// 对应 dijkstra 开始是 k 结束是所有的节点，然后取其中的最大值
}

func dijkstra(graph map[int][]*gn, start, end int) int {
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
		for _, next := range graph[pos] {
			if _, ok := visitedMap[next.pos]; ok {
				continue
			}
			heap.Push(h, &gn{
				pos:  next.pos,
				cost: cost + next.cost,
			})
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
