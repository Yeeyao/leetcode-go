package tree

/*
102. Binary Tree Level Order Traversal
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
返回树的层序遍历
直接使用队列，将根先放入，然后取队首，将其左右子树入队列后队首元素保存到临时结果

根入队列
循环判断条件是队列非空
遍历队列元素，保存到答案，然后每个元素的左右子树保存到临时结果
临时结果赋值给队列
最后返回结果
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	var temp []*TreeNode
	var addRes []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queueLen := len(queue)
		temp = []*TreeNode{}
		addRes = []int{}
		for i := 0; i < queueLen; i++ {
			addRes = append(addRes, queue[i].Val)
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		res = append(res, addRes)
		queue = temp
	}
	return res
}
