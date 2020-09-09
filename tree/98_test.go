package tree

import "math"

/*
98. Validate Binary Search Tree
Given a binary tree, determine if it is a valid binary search tree (BST).
Assume a BST is defined as follows:
    The left subtree of a node contains only nodes with keys less than the node's key.
    The right subtree of a node contains only nodes with keys greater than the node's key.
    Both the left and right subtrees must also be binary search trees.

[ref]() https://leetcode-cn.com/problems/validate-binary-search-tree/solution/yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution/
给定一个二叉树，判断它是否是一个合法的二叉搜索树，这里要需要判断上层根节点。。。
所以递归中需要传递最大最小值
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)
}
