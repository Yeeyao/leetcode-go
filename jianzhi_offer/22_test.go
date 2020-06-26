package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("22 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。", func(t *testing.T) {
		A := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
		get := solution(A)
		want := []int{3, 9, 20, 15, 7}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	直接递归打印 TLE
*/
func solution(root *TreeNode) []int {
	var res []int
	solutionHelper(root, &res)
	return res
}

func solutionHelper(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		solutionHelper(root.Left, res)
		solutionHelper(root.Right, res)
	}
}

/*
	使用队列处理
	初始化列表以及队列，先将 root 保存到队列中
	队列为空则跳出循环。
		将队首元素出队并将数值保存到列表
		如果左右两个子节点非空，则左右子节点入队列

			1
		2		3
	 4    5   6   7
*/
func solution2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	queueHead := 0
	queue := []*TreeNode{root}
	queueHead++
	// 非空
	for queueHead != 0 {
		head := queue[0]
		res = append(res, head.Val)
		queueHead--
		if head.Left != nil {
			queue = append(queue, head.Left)
			queueHead++
		}
		if head.Right != nil {
			queue = append(queue, head.Left)
			queueHead++
		}
	}
	return res
}
