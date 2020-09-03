package tree

/*
114. Flatten Binary Tree to Linked List
Given a binary tree, flatten it to a linked list in-place.
给定一个二叉树，将它转换为链表，原地处理
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	[ref](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/solution/er-cha-shu-zhan-kai-wei-lian-biao-by-leetcode-solu/)
	直接中序遍历二叉树来更改 怎么一边遍历一边修改元素的指针
	这里只是一个二叉树，元素之间的大小没有保证。。。同时 in place 这里理解成使用 O(1) 空间？
	emm，被例子误导了，不需要进行大小的处理，只需要处理指向的关系

	前序遍历 使用列表保存前序遍历的结果，这里用递归进行遍历
	然后将每个节点进行处理，这里也叫原地。。。
*/
func flatten(root *TreeNode) {
	nodeList := preorderTraversal(root)
	for i := 1; i < len(nodeList); i++ {
		pre, cur := nodeList[i-1], nodeList[i]
		pre.Right, cur.Left = cur, nil
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
}

/*
	找前驱节点的方法，对于当前节点，如果左子节点非空，在左子树找到最右节点，作为前驱节点。
	将当前节点的右子节点赋给前驱节点，然后将当前节点的左子节点赋值给当前节点的右子节点，
	将当前节点的左子节点设置为空，对当前节点处理完成，继续处理链表下一个节点，这里是右边的节点

	循环条件是 cur != nil
		判断 cur.Left != nil
		找到 left 的最右节点，将 cur.Right 赋值给该节点
		cur.Left 赋值为 nil， cur.Right 赋值为 cur.Left
		然后 cur = cur.Right 前进
*/
func flatten2(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}
