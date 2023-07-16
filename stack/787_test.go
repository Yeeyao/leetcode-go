package stack

import (
	"container/heap"
	"fmt"
)

/*
有 n 个城市通过 m 个航班连接。每个航班 flights 从城市 from 开始，以价格 price 抵达 to。
现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到从 src 到 dst 最多经过 k 站中转的最便宜的价格。 如果没有这样的路线，则输出 -1。
不超过 k 站

这里的问题在于 cost 小的，站数不一定小，需要判断如何选择。这里就是如果站数超过 k 就认为不能到达
- 这里就是取出一个节点的图之后，判断它到相邻节点的 step 如果超过 k 就不需要放入堆，这样保证堆里面的都是在 k steps 的

TLE 了

*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// 构造 map
	graph := make(map[int][]*gn)
	for _, v := range flights {
		from := v[0]
		to := v[1]
		cost := v[2]
		if _, ok := graph[from]; ok {
			graph[from] = append(graph[from], &gn{
				pos:  to,
				cost: cost,
				step: 0,
			})
		} else {
			graph[from] = []*gn{{to, cost, 0}}
		}
	}
	return dijkstra(graph, src, dst, k+1)
}

type toAndStep struct {
	to   int
	step int
}

func dijkstra(graph map[int][]*gn, start, end, k int) int {
	// 先将开始节点放入
	h := &graphNodes{gnList: []*gn{}}
	// 从 start 到 start cost 0
	h.Push(&gn{
		pos:  start,
		cost: 0,
		step: 0,
	})
	// 记录已经访问的节点，这里需要记录 step，因为相同的 to 节点，不同的 step 路径是不同的
	visitedMap := make(map[string]interface{})
	for h.Len() > 0 {
		topNodes := heap.Pop(h).(*gn)
		pos := topNodes.pos
		cost := topNodes.cost
		step := topNodes.step
		// 已经遍历过就跳过
		ts := fmt.Sprintf("%d,%d", pos, step)
		if _, ok := visitedMap[ts]; ok {
			continue
		}
		visitedMap[ts] = struct{}{}
		// 大于 k 就表示超过 step 了，继续判断
		if step > k {
			continue
		}
		// 如果已经遇到目标的节点了，直接返回
		if pos == end {
			return cost
		}
		for _, next := range graph[pos] {
			ts = fmt.Sprintf("%d,%d", next.pos, step+1)
			if _, ok := visitedMap[ts]; ok {
				continue
			}
			heap.Push(h, &gn{
				pos:  next.pos,
				cost: cost + next.cost,
				step: step + 1,
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
	step int
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
