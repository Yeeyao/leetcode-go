package tree

/*
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。

直接中序遍历，可以这样做，但是没有利用二叉搜索树的特性
二叉搜索树的中序遍历是单调递增的
*/

/*
[ref](https://leetcode.cn/problems/successor-lcci/solution/hou-ji-zhe-by-leetcode-solution-6hgc/)
使用二叉搜索树的性质，这里中序遍历二叉搜索树的性质，则后继节点大于 P 的值，后继节点是 p 之后所有节点的数值最小节点
如果 p 右子树非空，则 p 后继节点在右子树中，且在其最左的节点。如果 p 右子树为空，则后继节点是 p 的根
遍历节点，ret 表示 p 后继节点，cur 表示当前遍历的节点。
	如果 cur > p 则 ret = cur。然后继续向 cur 的左子树遍历
	如果 cur < p 则继续向 cur 的右子树遍历
*/

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ret *TreeNode
	// 如果 p 的右子树非空，则可以直接查找，不需要从 root 遍历了
	if p.Right != nil {
		ret = p.Right
		// 一直找到最左边的子树
		for ret.Left != nil {
			ret = ret.Left
		}
		return ret
	}
	cur := root
	for cur != nil {
		if cur.Val > p.Val {
			ret = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ret
}

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
