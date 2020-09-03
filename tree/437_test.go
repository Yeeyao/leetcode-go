package tree

/*
437. Path Sum III
You are given a binary tree in which each node contains an integer value.
Find the number of paths that sum to a given value.
The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).
The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

给定二叉树，仅包含整型节点，找到树中的路径使路径中节点的和等于给定的和的路径数量
路径不需要从根节点开始且终止于叶子节点

对比根节点到叶子节点的回溯法，那边需要从根节点开始，然后某个路径大于和就剪枝，等于和还要判断终止节点是否是叶子节点
这里的回溯需要以每个节点为起点，上一次就需要递归，同时注意这里是统计路径数量 所以每个路径也只需要统计和就行
需要注意这里的元素值有正负，所以不能因为临时和大于数值就直接剪枝，直接递归，满足要求就保存
*/

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	helper1(root, sum, &count)
	return count
}

func helper1(root *TreeNode, sum int, count *int) {
	helper2(root, count, 0, sum)
	if root.Left != nil {
		helper1(root.Left, sum, count)
	}
	if root.Right != nil {
		helper1(root.Right, sum, count)
	}
}

func helper2(node *TreeNode, count *int, temp, sum int) {
	temp += node.Val
	// 注意这里不能 return 因为可能还有路径
	if temp == sum {
		*count++
	}
	if node.Left != nil {
		helper2(node.Left, count, temp, sum)
	}
	if node.Right != nil {
		helper2(node.Right, count, temp, sum)
	}
}

/*
	非尾递归
*/
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := countPath(root, sum)
	left := pathSum(root.Left, sum)
	right := pathSum(root.Right, sum)
	return res + left + right
}

func countPath(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum -= node.Val
	res := 0
	if sum == 0 {
		res = 1
	}
	return res + countPath(node.Left, sum) + countPath(node.Right, sum)
}
