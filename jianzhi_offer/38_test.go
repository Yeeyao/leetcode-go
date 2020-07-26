package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("38 二叉树的深度 ", func(t *testing.T) {
		root := TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}
		get := solution(&root)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("38 二叉树的深度2 ", func(t *testing.T) {
		root := TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}
		get := solution2(&root)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度
	注意这里会给定 null [3,9,20,null,null,15,7] 3
	同 leetcode 104
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	可以用递归，将左右两边的深度相加
	如果是 null 则返回深度是 0，上一层就 + 1
	左右两个子树取较大值来加 DFS
*/
func solution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if solution(root.Left) > solution(root.Right) {
		return 1 + solution(root.Left)
	} else {
		return 1 + solution(root.Right)
	}
}

/*
	BFS 使用队列处理，将跟节点入队列，然后每次从队列取一个元素，将左右节点放入队列
	先将 root 放入队列中，然后判断队列非空，将队列中的所有元素的左右节点入队列
*/
func solution2(root *TreeNode) int {
	cnt := 0
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var tmp []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		// 前面的元素不需要处理，已经遍历了
		queue = tmp
		cnt++
	}
	return cnt
}
