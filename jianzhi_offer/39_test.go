package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("39 平衡二叉树 ", func(t *testing.T) {
		root := TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}
		get := solution(&root)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
	直接递归判断每个节点的左右子树深度
	停止 root 为 nil 表示到叶子节点了，返回 0 z左右子树高度有一个是 -1 直接返回 -1
	后序遍历  + 剪枝
	左右子树深度差 <= 1 则返回当前子树深度，root 的左右子树的最大深度 + 1
	如果 > 2 则直接返回 -1
			5
		3		9
	 1     4  6   10
5 9 10 6 3 4 1
p(4)
p(3) p(3)
p(2) p(2) p(2) p(2)
p(1) p(1) p(1) p(1) p(1)p(1)
4 3 2 1 1 2 1 1 3 2 1 1 2 1 1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	叶子节点左右深度都是 0
q	中间节点深度
*/
func solution(root *TreeNode) bool {
	return solutionHelper(root) != -1
}

func solutionHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := solutionHelper(root.Left)
	if left == -1 {
		return -1
	}
	right := solutionHelper(root.Right)
	if right == -1 {
		return -1
	}
	if left-right > 0 {
		if left-right < 2 {
			return left + 1
		} else {
			return -1
		}
	} else {
		if right-left < 2 {
			return right + 1
		} else {
			return -1
		}
	}
}
