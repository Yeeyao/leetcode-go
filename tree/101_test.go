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
	return isSymmetricHelper(root.Left, root.Right)
}

/*
这里其实就是从 root 最开始的两个子树开始，判断它们数值，然后继续判断子树的对称的子树。2 个向下递归判断 4 个，最终到叶子节点
还是要想到这里递归的处理是 2 个目标子树进行分裂递归判断，每一对分裂成对称的两对，最终比较所有
*/
func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Right, right.Left) && isSymmetricHelper(left.Left, right.Right)
}

/*
	迭代的方式，其实就是将递归的需要判断的子树放到一个自定义的栈里面
*/
func isSymmetricIt(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 最开始的左右子树
	if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	nodeList := []*TreeNode{root.Left, root.Right}
	for len(nodeList) > 0 {
		// 先取出当前的两个对称的节点
		leftNode, rightNode := nodeList[0], nodeList[1]
		if leftNode.Val != rightNode.Val {
			return false
		}
		nodeList = nodeList[2:]
		if leftNode.Left != nil && rightNode.Right != nil {
			nodeList = append(nodeList, leftNode.Left, rightNode.Right)
		} else if leftNode.Left == nil && rightNode.Right != nil ||
			leftNode.Left != nil && rightNode.Right == nil {
			return false
		}
		if leftNode.Right != nil && rightNode.Left != nil {
			nodeList = append(nodeList, leftNode.Right, rightNode.Left)
		} else if leftNode.Right == nil && rightNode.Left != nil ||
			leftNode.Right != nil && rightNode.Left == nil {
			return false
		}
	}
	return true
}
