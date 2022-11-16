package tree

/*
662. Maximum Width of Binary Tree
这里说了是完全二叉树
求树的最大宽度，即所有层级的宽度的最大值。每个层级的宽度是其最左到最右的非空空节点的长度，它们之间的 nil 节点也算入长度中
每一层都需要计算，至少需要两个非空的节点才能计算
因为需要计算中间的 nil 节点，看起来上层的信息需要保留？

BFS 使用 stack 保存每层的节点，然后判断每一层计算 width 并获得最大的
*/

func widthOfBinaryTree(root *TreeNode) int {
	var maxWidth int
	if root == nil {
		return maxWidth
	}
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		oldLen := len(nodeList)
		beginIndex := 0
		// 本层处理
		for i := 0; i < oldLen; i++ {
			if nodeList[i] != nil {
				nodeList = append(nodeList, nodeList[i].Left, nodeList[i].Right)
				width := i - beginIndex + 1
				if width > maxWidth {
					maxWidth = width
				}
			} else {
				nodeList = append(nodeList, nil, nil)
			}
		}
		nodeList = nodeList[oldLen:]
	}
	return maxWidth
}
