package jianzhi_offer

type TreeNode struct {
	val         int
	left, right TreeNode
}

func solution(preOrder, inOrder []int) TreeNode {
	poLen := len(preOrder)
	if poLen == 0 {
		return nil
	}
	// 先将中序的数值和索引保存到 map 中，方便 preOrder 中的根找到其左右子树的节点数量
	indexMap := make(map[int]int)
	for i := 0; i < poLen; i++ {
		indexMap[inOrder[i]] = i
	}
	// 两个遍历的开头和结尾以及两个遍历和 map
	root := solutionHelper(0, poLen - 1, 0, poLen - 1, preOrder, inOrder, indexMap)
	return root
}

/*
 */
func solutionHelper(poStart, poEnd, ioStart, ioEnd int ,preOrder, inOrder []int, indexMap map[int]int) TreeNode {
	// 索引出问题
	if poStart > poEnd {
		return nil
	}
	// 当前遍历的根的数值以及根数值在 inOrder 中的索引
	rootVal := preOrder[poStart]
	rootIndex := indexMap[rootVal]
	root := TreeNode{rootVal, nil, nil}
	if poStart == poEnd {
		return root
	}
	// 计算当前根左右子树的节点数量
	lstNum := rootIndex - ioStart
	//rstNum := ioEnd - rootIndex
	// 注意这里的下标更新 preOrder 用掉一个就需要向前，inOrder 用掉 rootIndex 就分别向两边移动，根据 rootIndex 移动
	// 同时两个遍历的递归范围是一样的
	// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
	leftSubTree := solutionHelper(poStart + 1, poStart + lstNum, ioStart, rootIndex - 1, preOrder, inOrder, indexMap)
	rightSubTree := solutionHelper(poStart + lstNum + 1, poEnd, rootIndex  + 1, ioEnd, preOrder, inOrder, indexMap)
	root.left = leftSubTree
	root.right = rightSubTree
	return root
}
