package tree

/*
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。

直接中序遍历，可以这样做，但是没有利用二叉搜索树的特性
*/

var nextNode *TreeNode
var isFound bool

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	nextNode = nil
	isFound = false
	inorderSuccessorHelper(root, p)
	return nextNode
}

func inorderSuccessorHelper(root *TreeNode, p *TreeNode) {
	if root == nil {
		return
	}
	inorderSuccessorHelper(root.Left, p)
	if isFound {
		nextNode = root
		isFound = false
		return
	}
	if root == p {
		isFound = true
	}
	inorderSuccessorHelper(root.Right, p)
}
