package tree

/*
653. Two Sum IV - Input is a BST
Given a Binary Search Tree and a target number,
return true if there exist two elements in the BST such that their sum is equal to the given target.

给定二叉搜索树以及一个 target，判断是否存在两个节点的和等于 target
直接先序遍历将所有元素保存到新建立的数组，然后利用二分查找
可以直接遍历树操作？
*/
var nodeS []int

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	nodeS = make([]int, 0)
	helper(root)
	SLen := len(nodeS)
	left, right := 0, SLen-1
	for left < right {
		if nodeS[left]+nodeS[right] == k {
			return true
		} else if nodeS[left]+nodeS[right] < k {
			left++
		} else {
			right--
		}
	}
	return false
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		helper(root.Left)
	}
	nodeS = append(nodeS, root.Val)
	if root.Right != nil {
		helper(root.Right)
	}
}

/*
	另一种，直接用一个 map 类似 普通 2Sum 只是遍历树的元素
	[ref](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/discuss/106067/C%2B%2BPython-Straight-Forward-Solution)
	这种没有利用到 BST 的特性，但是能提前终止，所以速度回快一些
*/
var m map[int]bool

func findTarget2(root *TreeNode, k int) bool {
	m = make(map[int]bool, 0)
	return helper2(root, k)
}

func helper2(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if m[root.Val] == true {
		return true
	}
	m[k-root.Val] = true
	return helper2(root.Left, k) || helper2(root.Right, k)
}
