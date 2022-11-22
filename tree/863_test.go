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
5 排除，避免出现环
*/

var distOneMap map[int][]int

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	distOneMap = make(map[int][]int)
	genDistOneMap(root)
	return nil
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
