package jianzhi_offer

import (
	"strconv"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("61 序列化二叉树", func(t *testing.T) {
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
	请实现两个函数，分别用来序列化和反序列化二叉树。
	序列化就是层次遍历
	反序列化是从数组重建二叉树
*/
func solution(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	queue := []*TreeNode{root}
	var res []string
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			res = append(res, "nil")
		}
	}
	return res
}

/*
	反序列化 这里输入的是有 nil
	根节点之后递归调用左右子树
	设 m 为区间 [0,n] 中空节点(nil) 个数，
	node.Val 非 nil, node.left 索引为 2(n-m) + 1, node.right 索引为 2(n-m) + 2

	如果 data 为空，则直接返回 nil
	指针 i = 0, root 为根，队列 queue
	queue 为空则跳出
		节点出队，记为 node
		构建 node 的左节点，node.left 值为 vals[i]，将 node.left 入队列
		i += 1
		构建 node 的右节点，node.right 值为 vals[i]，将 node.right 入队列
		i += 1
*/
func solution2(data []string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := TreeNode{0, nil, nil}
	// 当前访问的数组位置索引
	i := 1
	queue := []*TreeNode{&root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// 左子树非 nil
		if data[i] != "nil" {
			iVal, err := strconv.Atoi(data[i])
			if err != nil {
				iVal = 0
			}
			node.Left = &TreeNode{iVal, nil, nil}
			queue = append(queue, node.Left)
		}
		i++
		// 右子树非 nil
		if data[i] != "nil" {
			iVal, err := strconv.Atoi(data[i])
			if err != nil {
				iVal = 0
			}
			node.Right = &TreeNode{iVal, nil, nil}
			queue = append(queue, node.Right)
		}
		i++
	}
	return &root
}

/*

 */
func solution2Helper() {

}
