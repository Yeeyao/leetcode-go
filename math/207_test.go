package math

/*
207. Course Schedule 类似 210
[ref](https://leetcode-cn.com/problems/course-schedule/solution/ke-cheng-biao-by-leetcode-solution/)
TODO: 还需要仔细思考
There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
Some courses may have prerequisites, for example to take course 0 you have to first take course 1,
which is expressed as a pair: [0,1]
Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

给定从 0 到 numCourses - 1 的一共 numCourses 个课程，一些课程有提前的课程需要先修
使用序对表示，比如 [0,1] 表示修 1 前需要先修 0，给定课程和需求，判断可能修完全部课程
要求的课程中不存在重复的边

这里是 0 到 nc - 1 直接先判断 pre 数组 或者说只需要判断 pre 数组，只要出现相互依赖就返回 false
将相同开头的使用 map 保存，key 是 前面的，val 是后面的列表，然后每次遍历到一对就判断对应的后面的是否有相同的开头
但这样太暴力解了

DFS BFS 拓扑排序 对于图 G 中任意一条有向边 e(u,v) u 在排列中在 v 的前面
如果存在环，就不存在拓扑排序 有向无环图则拓扑排序不止一种
将每门课程看作一个节点，先修条件看作边

由于求出一种拓扑排序方法的最优时间复杂度为 O(n+m)，其中 n 和 m 分别是有向图 G 的节点数和边数
*/

/*
	DFS DFS 与拓扑排序联系，用一个栈存储所有已经搜索完成的节点
	当前搜索到节点 u，若它的相邻节点都已经搜索完了，则这些节点都已经在栈中，此时可以将
	u 入栈，u 的相邻节点都在栈内，从栈顶到栈底，u 满足拓扑排序要求
	对图进行一次深度优先搜索，当每个节点需要回溯时，将该节点放入栈中，
	最终从栈顶到栈底序列是拓扑排序

	将当前搜索的节点 u 标记为搜索中，遍历该节点每个相邻节点 v
		如果 v 为未搜索，开始搜索 v ，待搜索完成回溯到 u
		如果 v 为搜索中，则找到了图中的一个环，因此不存在拓扑排序
		如果 v 为已完成，说明 v 已经在栈中，u 不在栈中，u 无论何时入栈都不会影响到 (u,v) 间拓扑关系，不用操作
	当 u 所有相邻节点都为已完成，将 u 放入栈中，标记为已完成

	创建 edges 二维数组保存每条边中后继节点的前面的节点数组，visited 标记每个节点的状态，result 结果数组
	初始是将 edges 根据输入的边进行初始化，遍历每个节点，这里也要判断 valid 然后调用  dfs
	因为这里 edges 是以结束节点为开头保存的，所以 result 最终是拓扑排序的结果
	dfs
		当前节点 u 的 visited = 1 表示当前正在搜索，遍历当前节点的 edges 二维数组
			判断当前节点的前驱节点的 visited 状态
				如果是 0，表示没有遍历过，递归调用 dfs
				如果 valid 为 false 直接返回
				如果是 1，表示出现了环，直接 valid = false 并返回
			最后将 visited[u] = 2 表示已经搜索完，然后将 u 存放到 result


*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		// 当前节点标记为搜索中
		visited[u] = 1
		// 对每个相邻节点，如果搜索中就存在环，直接返回
		for _, v := range edges[u] {
			if visited[v] == 0 {
				// 这里对相邻节点递归调用即是 dfs
				dfs(v)
				if !valid {
					return
				} else if visited[v] == 1 {
					valid = false
					return
				}
			}
			// 搜索完成 这里先将 u 保存，即原来的边当中的后一个节点保存
			visited[u] = 2
			result = append(result, u)
		}
		// 这里将边中，比如 [u,v] 将所有 v 结尾的边的开始节点保存到 edges
		for _, info := range prerequisites {
			edges[info[1]] = append(edges[info[1]], info[0])
		}
		// 对每个节点进行 dfs
		for i := 0; i < numCourses && valid; i++ {
			if visited[i] == 0 {
				dfs(i)
			}
		}
	}
	return valid
}

/*
	广度优先 BFS
	考虑拓扑排序中最前面的节点，它一定不会有任何入边
	用一个队列进行广度优先搜索，初始时，所有入度为 0 的节点被放入队列中，
	它们可以作为拓扑排序最前面的节点，且它们相对顺序无关紧要

	广度优先搜索每一步，取出队首节点 u
		我们将 u 放入答案中
		移除 u 的所有出边，即将 u 的所有相邻系欸但入度减少 1，如果某个相邻节点入度变 0 则将 v 放入队列
	广度优先搜索结束后，如果答案包含了这 n 个节点，则找到了一种拓扑排序，否则说明图中存在环
	就不存在拓扑排序
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		inEdge = make([]int, numCourses)
		result []int
	)

	// edges 同样保存每个出节点的入节点队列
	// inedge 保存所有入节点
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		inEdge[info[0]]++
	}

	// 先将非入节点保存到 q
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if inEdge[i] == 0 {
			q = append(q, i)
		}
	}

	// 取队列首，先保存到结果，然后将队首节点作为出节点的入节点列表中节点的入度 - 1
	// 当入度为 0 就保存到 q
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			inEdge[v]--
			if inEdge[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
