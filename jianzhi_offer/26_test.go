package jianzhi_offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	二叉树的中序遍历就是一个单调递增的元素序列，因此使用中序遍历
	相邻节点的关系，pre, cur pre.right = cur, cur.left = pre
	同时，头尾的 head,tail head.left = tail, tail.right = head

	这里先中序遍历，最后头尾处理
*/
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
