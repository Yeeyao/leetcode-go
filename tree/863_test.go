package tree

/*
给定二叉树的 root 以及一个 target 节点和整型 k，返回距离 target 节点为 k 的所有节点的列表，可以按照任何顺序返回
所有节点的数值都是唯一的
距离的计算，有点类似 987 的构建坐标然后才能计算距离，同时应该可以剪枝
要先找到 target，然后途中记录经过的节点的坐标吗？然后继续向下？不行，因为这里是计算经过的节点而不是位置

哈夫曼编码？所有节点都有一个 parent，然后大家都只计算到 parent 的距离，到某个节点的距离计算就将从 parent 向下计算
每个节点都记录下到自己为 1 的其他节点，然后计算到目标节点的距离的时候就可以直接找，出现环的就跳过
比如题目第一个例子中，统计每个节点距离为 1 的其他节点

3 [5, 1]
5 [3, 6, 2]
6 [5]
2 [5, 7, 4]
7 [2]
4 [2]
1 [3, 0, 8]
0 [1]
8 [1]

然后从 target 的距离为 1 的节点找，找到下一个节点的时候，从它的距离为 1 的节点中排除掉到达它的节点，比如 5 找到 3 然后将 3 的距离为 1 的节点中的
5 排除，避免出现环，因此 5，3，1 一直下去
*/

var distOneMap map[int][]int
var retList []int

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	retList = make([]int, 0)
	distOneMap = make(map[int][]int)
	genDistOneMap(root)
	getTarget(-1, target.Val, k)
	return retList
}

func getTarget(prev, cur, k int) {
	if k == 0 {
		retList = append(retList, cur)
		return
	}
	k--
	nodeList, ok := distOneMap[cur]
	if ok {
		for _, v := range nodeList {
			if v != prev {
				getTarget(cur, v, k)
			}
		}
	}
}

func genDistOneMap(root *TreeNode) {
	if root == nil {
		return
	}
	rootVal := root.Val
	if root.Left != nil {
		leftVal := root.Left.Val
		if _, ok := distOneMap[rootVal]; ok {
			distOneMap[rootVal] = append(distOneMap[rootVal], leftVal)
		} else {
			distOneMap[rootVal] = []int{leftVal}
		}
		if _, ok := distOneMap[leftVal]; ok {
			distOneMap[leftVal] = append(distOneMap[leftVal], rootVal)
		} else {
			distOneMap[leftVal] = []int{rootVal}
		}
	}
	if root.Right != nil {
		rightVal := root.Right.Val
		if _, ok := distOneMap[rootVal]; ok {
			distOneMap[rootVal] = append(distOneMap[rootVal], rightVal)
		} else {
			distOneMap[rootVal] = []int{rightVal}
		}
		if _, ok := distOneMap[rightVal]; ok {
			distOneMap[rightVal] = append(distOneMap[rightVal], rootVal)
		} else {
			distOneMap[rightVal] = []int{rootVal}
		}
	}
	genDistOneMap(root.Left)
	genDistOneMap(root.Right)
}

// 最快的答案，感觉思路是类似的，但是我感觉自己的容易理解一点
type Entry struct {
	Node *TreeNode
	Dist int
}

type nodeSet map[int]*TreeNode

func (s nodeSet) add(node *TreeNode) {
	s[node.Val] = node
}

func (s nodeSet) remove(node *TreeNode) {
	delete(s, node.Val)
}

// 这里找到每个节点的与之连接的节点，这里是 map slice 元素是 map
func dfs(root *TreeNode, graph map[int]nodeSet) {
	if _, ok := graph[root.Val]; !ok {
		graph[root.Val] = make(nodeSet)
	}
	if root.Left != nil {
		graph[root.Val].add(root.Left)
		if _, ok := graph[root.Left.Val]; !ok {
			graph[root.Left.Val] = make(nodeSet)
		}
		graph[root.Left.Val].add(root)
		dfs(root.Left, graph)
	}
	if root.Right != nil {
		graph[root.Val].add(root.Right)
		if _, ok := graph[root.Right.Val]; !ok {
			graph[root.Right.Val] = make(nodeSet)
		}
		graph[root.Right.Val].add(root)
		dfs(root.Right, graph)
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 1. Build a graph by traversing the tree
	graph := make(map[int]nodeSet)
	dfs(root, graph)

	// 2. Run DFS w/ disconnecting edges to the parent after adding to the queue
	ret := make([]int, 0)
	q := make([]Entry, 0)
	// 从 target 出发
	q = append(q, Entry{target, 0})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		// 当前的节点到 target 距离等于 k 就保存
		if curr.Dist == k {
			ret = append(ret, curr.Node.Val)
		} else if curr.Dist > k {
			// 如果大于就跳过
			break
		} else {
			// 添加当前节点的相邻的节点，然后 dist 在当前的节点的基础上 + 1
			for nbrKey, nbrNode := range graph[curr.Node.Val] {
				q = append(q, Entry{nbrNode, curr.Dist + 1})
				// Disconnect the edge to the parent 如注释，这里断开和父母节点的连接
				graph[nbrKey].remove(curr.Node)
			}
		}
	}
	return ret
}
