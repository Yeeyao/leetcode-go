package math

/*
337 House Robber III
The thief has found himself a new place for his thievery again.
There is only one entrance to this area, called the "root."
Besides the root, each house has one and only one parent house.
After a tour, the smart thief realized that "all houses in this place forms a binary tree".
It will automatically contact the police if two directly-linked houses were broken into on the same night.
Determine the maximum amount of money the thief can rob tonight without alerting the police.

二叉树结构的房子，当两个相连的房子不能连续被抢
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	f(o) 表示选择 o 节点情况下，o 节点的子树上被选择的节点的最大权值和
	g(o) 表示不选择 o 节点情况下，o 节点的子树上被选择的节点的最大权值和
	l, r 表示 o 的左右子树

	o 被选中，o 的左右子树不能被选中，则 f(o) = g(l) + g(r)
	o 不被选中，o 左右子树可以被选中，也可以不被选中 g(o) = max(f(l), g(l)) + max(f(r), g(r))

	使用哈希表来映射 f 和 g 的函数值，用 DFS 后序遍历树
	初始化 f, g 两个 map
	主函数调用 dfs(root) 然后返回 f(root) g(root) 较大值
dfs
	如果 root 等于 nil 直接返回
	递归对左右子树调用后，更新 f(root) g(root)

[ref](https://leetcode-cn.com/problems/house-robber-iii/solution/da-jia-jie-she-iii-by-leetcode-solution/)
*/
var f, g map[*TreeNode]int

func rob(root *TreeNode) int {
	f, g = make(map[*TreeNode]int), make(map[*TreeNode]int)
	dfs(root)
	return max(f[root], g[root])
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
	f[root] = root.Val + g[root.Left] + g[root.Right]
	g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])
}

// 这里使用二维数组保存结果
func rob2(root *TreeNode) int {
	val := dfs2(root)
	return max(val[0], val[1])
}

func dfs2(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l, r := dfs2(root.Left), dfs2(root.Right)
	selected := root.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	那直接将每一层的总和都求出来，这里可以直接求，不对，因为可以选上一层的节点，然后下一层的节点只要不和上层的连接就行
	计算最大和 dp[i] = max((dp[i - 2] + levelSum[i]), dp[i-1]) 这样思路不对
*/
//func rob(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	levelSum := make([]int, 0)
//	var temp []*TreeNode
//	var tempSum int
//	queue := []*TreeNode{root}
//	for len(queue) > 0 {
//		tempSum = 0
//		temp = []*TreeNode{}
//		queueLen := len(queue)
//		for i := 0; i < queueLen; i++ {
//			tempSum += queue[i].Val
//			if queue[i].Left != nil {
//				temp = append(temp, queue[i].Left)
//			}
//			if queue[i].Right != nil {
//				temp = append(temp, queue[i].Right)
//			}
//		}
//		levelSum = append(levelSum, tempSum)
//		queue = temp
//	}
//	// 对每一行进行 dp 计算最大和 类似第一种
//	level := len(levelSum)
//	dp := make([]int, level)
//	if level == 1 {
//		return levelSum[0]
//	}
//	if level == 2 {
//		return max(levelSum[0], levelSum[1])
//	}
//	dp[0], dp[1] = levelSum[0], max(levelSum[0], levelSum[1])
//	for i := 2; i < level; i++ {
//		dp[i] = max(dp[i-2]+levelSum[i], dp[i-1])
//	}
//	return dp[level-1]
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
