package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("58 对称二叉树", func(t *testing.T) {
		root := Node{1, nil}
		get := solution(&root)
		want := nil
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
	对称二叉树 当前节点的左右节点需要和相邻的节点的左右节点相同
	同 leetcode 101
	L.val == R.val L.Left.Val == R.Right.Val L.Right.Val == R.Left.Val
*/
func solution(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return solutionHelper(root.Left, root.Right)
}

func solutionHelper(Left, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}
	// 这里比较数值
	if Left == nil || Right == nil || Left.Val != Right.Val {
		return false
	}
	// 递归判断的是
	// 当前节点的左子树和相邻节点的右子树 当前节点的右子树和相邻节点的左子树
	return solutionHelper(Left.Left, Right.Right) && solutionHelper(Left.Right, Right.Left)
}
