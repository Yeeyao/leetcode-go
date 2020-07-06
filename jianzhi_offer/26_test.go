package jianzhi_offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var pre *TreeNode
	helper(root, &pre)
	// 头尾两个的处理
	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left = tail
	tail.Right = head
	return head
}

func helper(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left, pre)
	if *pre != nil {
		root.Left = *pre
		(*pre).Right = root
	}
	*pre = root
	helper(root.Right, pre)
}
