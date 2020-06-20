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
func soution(root *TreeNode) *TreeNode {
	solutionHelper(root.Left)
	solutionHelper(root.Right)
	return root
}

func solutionHelper(root *TreeNode) {
	if root == nil {
		return
	}
	solutionHelper(root.Left)
	solutionHelper(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}
