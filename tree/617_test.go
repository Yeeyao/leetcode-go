package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given two binary trees and imagine that when you put one of them to cover the other,
some nodes of the two trees are overlapped while the others are not.
You need to merge them into a new binary tree. The merge rule is that if two nodes overlap,
then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

给定两个二叉树，将两个二叉树合并，相同位置的节点就将数值相加
直接用递归
当左右两个节点其中一个是 nil 就返回另外一个
先计算当前节点的数值，然后当前的左右节点是递归调用的结果
*/
func mergeTree(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	root := &TreeNode{0, nil, nil}
	root.Val = t1.Val + t2.Val
	root.Left = mergeTree(t1.Left, t2.Left)
	root.Right = mergeTree(t1.Right, t2.Right)
	return root
}

func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// 直接使用第一个树保存结果
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
