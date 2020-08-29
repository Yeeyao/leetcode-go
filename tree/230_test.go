package tree

import "testing"

func TestPro(t *testing.T) {
	t.Run("230. Kth Smallest Element in a BST", func(t *testing.T) {
		input := "UD"
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
	What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently?
	How would you optimize the kthSmallest routine?

	给定一个二叉搜索树，找到最小第 k 个元素
	直接中序遍历来查找，如果当前遍历数量等于 k 就直接返回，否则继续遍历，最后返回 -1 表示不足

	使用两个全局变量 先遍历左子树，然后判断数值，最后遍历右子树
*/
var count, val int

func kthSmallest(root *TreeNode, k int) int {
	count = k
	helper(root)
	return val
}

func helper(root *TreeNode) {
	if root.Left != nil {
		helper(root.Left)
	}
	// 注意这里是先减 1 再判断
	count--
	if count == 0 {
		val = root.Val
		return
	}
	if root.Right != nil {
		helper(root.Right)
	}
}

/*
	iteration
	判断 k 以及 root 是否需要返回 0
	先将所有左子树根节点都入队列
	循环判断队列非空
		判断 k 是否 1 如果是，直接返回数值
		不是 1 将 k 递减 队列元素弹出
		将队首元素左子树不断入队列

*/
func kthSmallestIter(root *TreeNode, k int) int {
	if k <= 0 || root == nil {
		return 0
	}
	var queue []*TreeNode
	// 先将所有左子树根节点入队列
	for root != nil {
		queue = append(queue, root)
		root = root.Left
	}
	for len(queue) > 0 {
		// 取队首元素
		top := queue[len(queue)-1]
		if k == 1 {
			return top.Val
		}
		k--
		queue = queue[:len(queue)-1]
		// 需要将当前节点右子树的左子树入队列
		right := top.Right
		for right != nil {
			queue = append(queue, right)
			right = right.Left
		}
	}
	return 0
}
