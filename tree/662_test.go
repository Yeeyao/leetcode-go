package tree

/*
662. Maximum Width of Binary Tree
求树的最大宽度，即所有层级的宽度的最大值。每个层级的宽度是其最左到最右的非空空节点的长度，它们之间的 nil 节点也算入长度中
每一层都需要计算，至少需要两个非空的节点才能计算
因为需要计算中间的 nil 节点，看起来上层的信息需要保留？

BFS 使用 stack 保存每层的节点，然后判断每一层计算 width 并获得最大的

从本层最左边开始的节点开始计算。所以说最大宽度的层类似完全二叉树的结构
*/

// 这里出现一层只有一个节点的情况需要处理，本层之后列表都要清空
// 这种方法 OOM 了
func widthOfBinaryTree(root *TreeNode) int {
	var maxWidth int
	if root == nil {
		return maxWidth
	}
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		oldLen := len(nodeList)
		// 本层处理
		isLastLayer := true
		beginIndex := -1
		for i := 0; i < oldLen; i++ {
			// 需要判断是否是最后一层
			if nodeList[i] != nil {
				if beginIndex == -1 {
					beginIndex = i
				}
				isLastLayer = false
				nodeList = append(nodeList, nodeList[i].Left, nodeList[i].Right)
				width := i - beginIndex + 1
				if width > maxWidth {
					maxWidth = width
				}
			} else {
				if beginIndex > -1 {
					nodeList = append(nodeList, nil, nil)
				}
			}
		}
		if isLastLayer {
			break
		}
		nodeList = nodeList[oldLen:]
	}
	return maxWidth
}
