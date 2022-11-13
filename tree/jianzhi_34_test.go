package tree

import (
	"fmt"
	"reflect"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	[ref](https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/)

	给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
	叶子节点 是指没有子节点的节点。

	这个也是 DFS 了，然后需要直接到叶子节点，可以提前终止，如果当前的路径下的总和大于目标和
*/
func TestPro(t *testing.T) {
	t.Run("jianzhi 33 叉树中和为某一值的路径 ", func(t *testing.T) {
		root := TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		}
		resPath := pathSumOfficial(&root, 22)
		fmt.Println(resPath)
		want := [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}
		if !reflect.DeepEqual(resPath, want) {
			t.Errorf("got: %v, want: %v", resPath, want)
		}
	})
}

func pathSumJianzhi34(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	pathSumHelper(root, target, []int{}, &res)
	return res
}

// 前序处理问题
func pathSumHelper(root *TreeNode, sum int, tempPath []int, res *[][]int) {
	if root == nil {
		return
	}
	sum -= root.Val
	tempPath = append(tempPath, root.Val)
	// 到达叶子节点
	if root.Left == nil && root.Right == nil && sum == 0 {
		// go slice 的坑，这里的意思是强制开辟新的存储空间保存当前的结果？
		*res = append(*res, append([]int(nil), tempPath...))
		return
	}
	pathSumHelper(root.Left, sum, tempPath, res)
	pathSumHelper(root.Right, sum, tempPath, res)
}

// 这里的 defer
func pathSumOfficial(root *TreeNode, target int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(rootNode *TreeNode, sum int) {
		if rootNode == nil {
			return
		}
		sum -= rootNode.Val
		path = append(path, rootNode.Val)
		// 不同结果需要清空
		defer func() { path = path[:len(path)-1] }()
		if rootNode.Left == nil && rootNode.Right == nil && sum == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(rootNode.Left, sum)
		dfs(rootNode.Right, sum)
	}
	dfs(root, target)
	return
}
