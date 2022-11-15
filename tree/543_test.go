package tree

/*
	这里 节点 的 diameter 其实就是左右子树的高度和
	同时因为 diameter 可能不经过 root，因此每个节点都需要计算一次，然后更新最大值
	一开始不确定可以先计算每个节点的高度，然后将其左右子树的高度相加就是当前节点的 diameter，
	然后发现计算 height 的时候其实已经得到的当前节点的 diameter 因此直接在计算高度的时候顺便判断
*/
var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	getHeight(root)
	return maxDiameter
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	diameter := leftHeight + rightHeight
	if diameter > maxDiameter {
		maxDiameter = diameter
	}
	if leftHeight > rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}

// 最开始的版本
func diameterOfBinaryTreeOld(root *TreeNode) int {
	maxDiameter = 0
	diameterOfBinaryTreeHelper(root)
	return maxDiameter
}

// 这里需要计算的同时判断是不是最大的
func diameterOfBinaryTreeHelper(root *TreeNode) {
	if root == nil {
		return
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	diameter := leftHeight + rightHeight
	if diameter > maxDiameter {
		maxDiameter = diameter
	}
}

func getHeightOld(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if leftHeight > rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}
