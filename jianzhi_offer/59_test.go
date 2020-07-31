package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("59 按之字形顺序打印二叉树 ", func(t *testing.T) {
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
请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，
其他行以此类推。
类似层次打印，只是需要记录打印方向，根据方向，保存在临时列表的左边或者右边
其实只需要队列也能打印
*/
func solution(root *TreeNode) [][]int {
	var res [][]int
	var isReverse bool
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		qLen := len(queue)
		// 用列表保存，根据方向保存到列表的开头和尾部
		list := make([]int, qLen)
		for i := 0; i < qLen; i++ {
			node := queue[i]
			if !isReverse {
				list[i] = node.Val
			} else {
				list[qLen-i-1] = node.Val
			}
			// 左右子树入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将前面遍历过的数据删除
		queue = queue[qLen:]
		// 改变方向
		isReverse = !isReverse
		res = append(res, list)
	}
	return res
}

//func solution(root *TreeNode) {
//	queue := []*TreeNode{root}
//	direction := 0
//	for len(queue) != 0 {
//		// 先打印，确认方向从左到右
//		qLen, begin := len(queue), len(queue)
//		if direction == 0 {
//			for i := 0; i < qLen; i++ {
//				// 先将左右子树入队列
//				if queue[i].Left != nil {
//					begin++
//					queue[begin] = queue[i].Left
//				}
//				if queue[i].Right != nil {
//					begin++
//					queue[begin] = queue[i].Right
//				}
//				// 打印
//				fmt.Printf("val: %d\n", queue[i].Val)
//			}
//			// 将队列前面的删除
//			queue = queue[qLen:]
//			direction = 1
//		} else {
//			// 从右到左
//			for i := qLen - 1; i >= 0; i-- {
//				// 先将左右子树入队列
//				if queue[i].Left != nil {
//					begin++
//					queue[begin] = queue[i].Left
//				}
//				if queue[i].Right != nil {
//					begin++
//					queue[begin] = queue[i].Right
//				}
//				// 打印
//				fmt.Printf("val: %d\n", queue[i].Val)
//			}
//			// 将队列前面的删除
//			queue = queue[qLen:]
//			direction = 0
//		}
//	}
//}
