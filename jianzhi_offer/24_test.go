package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("24 二叉树中和为某一值的路径", func(t *testing.T) {
		A := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
		get := solution(A, 3)
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
	需要找到树种和为 sum 的所有节点的路径
	直接用 DFS 但是这样，如果树很大，就不太行
	需要注意这里是从根节点开始的，一直到叶子节点结束
*/
func solution(root *TreeNode, sum int) []int {
	var res [][]int
	solutionHelper(root, sum, []int{}, &res)
	return res
}

func solutionHelper(root *TreeNode, sum int, tempPath []int, res *[][]int) {
	// 遍历完一条路径，后面就不需要额外判断当前 root 的左右子树
	if root == nil {
		return
	}
	val := root.Val
	tempPath = append(tempPath, val)
	sum -= val
	if sum == 0 && root.Left == nil && root.Right == nil {
		*res = append(*res, tempPath)
		return
	}
	solutionHelper(root.Left, sum, tempPath, res)
	solutionHelper(root.Right, sum, tempPath, res)
}
