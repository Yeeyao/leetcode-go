package tree

/*
	递归处理
	如果当前节点是 nil 返回 nil
	先交换，然后递归调用左右子树
	最后返回 root

*/
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}

/*
	这里可以先递归，在最后进行交换
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}