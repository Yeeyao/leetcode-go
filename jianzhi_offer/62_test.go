package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("62 二叉搜索树的第k个结点", func(t *testing.T) {
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

var count int
var res int

/*
	给定一棵二叉搜索树，请找出其中的第k小的结点。例如， （5，3，7，2，4，6，8）  中，按结点数值大小顺序第三小结点的值为4。
	第 k 个，其左子树有 k - 1 个元素
	中序遍历，左子树，根，右子树，中序遍历的倒序是递减的序列，求第 k 大节点可以转换为中序倒序遍历的第 k 个节点
*/
func solution(root *TreeNode, k int) int {
	count = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		count--
		if count == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
}
