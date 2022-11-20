package tree

import (
	"fmt"
)

/*
	给定二叉树的根以及目标先序遍历的节点列表(每个节点的数值都是唯一的)，判断能否通过将树的节点反转（交换其左右子树）来使得树实现给定的先序遍历顺序
	返回需要反转的节点的数值列表，如果不能通过反转实现要求的遍历顺序则返回 [-1] ，无需反转则返回 []
*/

/*
	从给定地 root 开始前序遍历，遇到和所期望的节点不一样的时候进行判断
	这里如果从 root 开始判断，如果 root 不一样，那直接返回 [-1] 否则，判断左右子树数值是否可以交换，如果不行则直接返回 [-1]
	如果可以，则直接交换来遍历。如果一样，则递归处理 root 的左右子树，然后将它们返回的结果汇总
	同时，这里交换后需要使用交换后的节点遍历
*/

/*
	[官方](https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/description/)
*/

var resNode []int
var index int

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	resNode = []int{}
	index = 0
	flipMatchVoyageHelper(root, voyage)
	// 可能会出现前面的遍历中保存了 -1 但是递归的后面有节点可以反转，因此这里需要再次判断。-1 如果是最后就会被清空，只返回 []int{-1}
	if len(resNode) > 0 && resNode[0] == -1 {
		return []int{-1}
	}
	return resNode
}

func flipMatchVoyageHelper(root *TreeNode, voyage []int) {
	// 只有该节点非 nil 才需要处理，因为 voyage 的节点都是非 nil
	if root != nil {
		if root.Val != voyage[index] {
			resNode = []int{-1}
			return
		}
		index++
		// 如果左子树不同，表示该节点需要反转，同时先判断右子树
		if index < len(voyage) && root.Left != nil && root.Left.Val != voyage[index] {
			resNode = append(resNode, root.Val)
			flipMatchVoyageHelper(root.Right, voyage)
			flipMatchVoyageHelper(root.Left, voyage)
		} else {
			flipMatchVoyageHelper(root.Left, voyage)
			flipMatchVoyageHelper(root.Right, voyage)
		}
	}
}

var resVisit []int

// 前序遍历
func preOrderVisit(root *TreeNode) {
	resVisit = []int{}
	preOrderVisitHelper(root)
}

func preOrderVisitHelper(root *TreeNode) int {
	resVisit = append(resVisit, root.Val)
	leftNum, rightNum := 0, 0
	if root.Left != nil {
		leftNum = preOrderVisitHelper(root.Left)
	}
	if root.Right != nil {
		rightNum = preOrderVisitHelper(root.Right)
	}
	fmt.Println(root.Val, leftNum+rightNum)
	return leftNum + rightNum + 1
}

/*
错误的做法，这里其实思路一样是根据 root 的左右子树和当前的 voyage 的元素比较判断下一步需要使用哪个子树进行判断
但是判断完成其中一个子树之后，需要判断另外一个子树（如果非空）的时候，需要计算跳过多少个 voyage 元素。不如答案的直接使用 index 来计算
当前比较的 voyage 的元素的索引来的简单直观
*/

/*
var resNode []int
var invalid bool

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	resNode = []int{}
	invalid = false
	flipMatchVoyageHelper(root, voyage)
	if invalid {
		return []int{-1}
	}
	return resNode
}

func flipMatchVoyageHelper(root *TreeNode, voyage []int) int {
	leftNum, rightNum := 0, 0
	if root.Val != voyage[0] {
		invalid = true
		return 0
	}
	if root.Left != nil && root.Right != nil {
		if len(voyage) > 1 {
			if root.Left.Val == voyage[1] {
				leftNum = flipMatchVoyageHelper(root.Left, voyage)
				// 这里中间需要跳过多少个 voyage 呢？
				rightNum = flipMatchVoyageHelper(root.Right, voyage[leftNum+1:])
			} else if root.Right.Val == voyage[1] {
				// 如果是右边等于，则需要先遍历右边再遍历左边了
				rightNum = flipMatchVoyageHelper(root.Right, voyage)
				// 这里中间需要跳过多少个 voyage 呢？
				leftNum = flipMatchVoyageHelper(root.Left, voyage[rightNum+1:])
			} else {
				invalid = true
				return 0
			}
		}
	}
	if root.Left != nil {
		if root.Left.Val != voyage[0] {
			invalid = true
			return 0
		} else {
			leftNum = flipMatchVoyageHelper(root.Left, voyage)
		}
	}
	if root.Right != nil {
		if root.Right.Val != voyage[0] {
			invalid = true
			return 0
		} else {
			rightNum = flipMatchVoyageHelper(root.Right, voyage)
		}
	}
	return leftNum + rightNum + 1
}
*/
