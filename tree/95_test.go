package tree

import "fmt"

/*
	对比 95 这次要求返回所有的结果树，需要注意是 BST
	DFS 递归，只需要返回所有的结果树的 head
	开始，选择任意一个作为 head 结束，所有的节点都被使用了 答案，将结束的时候的 head 保存到返回的 head slice
*/

func generateTrees(n int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	fmt.Printf("out %p\n", ret)
	for i := 1; i <= n; i++ {
		head := &TreeNode{Val: i}
		usedI := make(map[int]struct{})
		usedI[i] = struct{}{}
		// 左边作为假的 head
		fakeHead := &TreeNode{Left: head}
		generateTreesHelper(n, usedI, fakeHead, head, &ret)
	}
	return ret
}

// head 确定了，然后下面需要选择下一个节点，作为左边或者右边，同时不能选择已经出现过的数字
// 需要 map 记录已经使用的数字？每个都需要作为当前 head 的左子树或者右子树的节点，head 上面的默认已经构造好了
func generateTreesHelper(n int, usedI map[int]struct{}, fakeHead, head *TreeNode, ret *[]*TreeNode) {
	isAllUsed := true
	for i := 1; i <= n; i++ {
		if _, ok := usedI[i]; !ok {
			isAllUsed = false
			usedI[i] = struct{}{}
			leftNode := &TreeNode{Val: i}
			head.Left = leftNode
			generateTreesHelper(n, usedI, fakeHead, head.Left, ret)
			head.Left = nil
			head.Right = leftNode
			generateTreesHelper(n, usedI, fakeHead, head.Right, ret)
		}
	}
	if isAllUsed {
		*ret = append(*ret, fakeHead.Left)
	}
}
