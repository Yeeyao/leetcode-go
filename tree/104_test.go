package tree

// TLE
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if maxDepth(root.Left) > maxDepth(root.Right) {
		return maxDepth(root.Left) + 1
	} else {
		return maxDepth(root.Right) + 1
	}
}

/*
	使用队列，然后每次都是保存一行的元素
	初始化一个队列以及临时队列和一个高度
	循环判断队列非空
		将队列中的所有节点的左右子树都保存到临时数组
		高度 + 1
		使用临时数组更新队列
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	count := 0
	for len(queue) > 0 {
		temp := []*TreeNode{}
		for _, v := range queue {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		count++
		queue = temp
	}
	return count
}
