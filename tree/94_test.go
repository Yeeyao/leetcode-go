package tree

/*
94. Binary Tree Inorder Traversal
二叉树中序遍历，直接递归
需要注意输入 nil 情况，然后左右子树需要先判断 nil 再递归调用
*/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	helper(root, &res)
	return res
}

// 最外层已经判断了 root 为 nil 的情况，因此里层函数不需要判断了
func helper(root *TreeNode, res *[]int) {
	if root.Left != nil {
		helper(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		helper(root.Right, res)
	}
}
