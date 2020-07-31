package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("60 把二叉树打印成多行  ", func(t *testing.T) {
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
	二叉树打印成多行
	从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。
	比上一题简单，直接用队列然后打印就行，先将根节点入队列，然后每次都遍历队列里面的元素保存起来，
	同时将队列的左右子树入队列
*/
func solution(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	var res [][]int
	for len(queue) != 0 {
		qLen := len(queue)
		var list []int
		// 遍历当前队列里面的元素，保存并将左右子树入队列
		for i := 0; i < qLen; i++ {
			list = append(list, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// 保存当前结果
		res = append(res, list)
		//过滤遍历过的元素
		queue = queue[qLen:]
		list = []int{}
	}
	return res
}
