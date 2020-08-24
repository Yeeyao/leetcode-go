package tree

/*
   Symmetric Tree
   Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
   For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

  给定一个二叉树，判断它是否对称
  递归处理，如果当前节点 nil 返回 true
  否则，调用递归辅助函数
     需要当前节点左右子树数值相同，或者两个节点为 nil
     如果其中一个 nil 或者两个的数值不同
     然后递归调用子树的左右部分
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return helper(left.Right, right.Left) && helper(left.Left, right.Right)
}
