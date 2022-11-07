package tree

/*
	[ref](https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/)
	给定一棵二叉树，求树的深度，从根节点到叶节点依次经过的节点形成树的一条路径，最长路径称为树的深度
	看起来是 dfs
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
