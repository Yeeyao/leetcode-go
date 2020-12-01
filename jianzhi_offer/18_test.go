package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("18 二叉树镜像", func(t *testing.T) {
		A := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
		B := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
		get := solution(A, B)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	二叉树镜像 递归
*/
func solution(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	solution(root.Left)
	solution(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func soution2(root *TreeNode) *TreeNode {
	solution2Helper(root.Left)
	solution2Helper(root.Right)
	return root
}

func solution2Helper(root *TreeNode) {
	if root == nil {
		return
	}
	solution2Helper(root.Left)
	solution2Helper(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}

func solution3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = solution3(root.Left), solution3(root.Right)
	return root
}
