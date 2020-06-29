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
*/
func solution(root *TreeNode, sum int) []int {
	var res [][]int
	solutionHelper(root, sum, 0, []int{}, &res)
	return res
}

func solutionHelper(root *TreeNode, sum, tempSum int, tempPath []int, res *[][]int) {
	if sum == tempSum {
		*res = append(*res, append([]int{}, tempPath...))
	} else if sum > tempSum {
		return
	} else {
		tempPath = append(tempPath, root.Val)
		tempSum += root.Val
		solutionHelper(root, sum, tempSum, tempPath, res)
		tempPath = tempPath[:len(tempPath)-1]
		tempSum -= root.Val
		if root.Left != nil {
			solutionHelper(root.Left, sum, tempSum, tempPath, res)
		}
		if root.Right != nil {
			solutionHelper(root.Right, sum, tempSum, tempPath, res)
		}
	}
}
