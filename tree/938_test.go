package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	return rangeSumBST2(root, L, R, 0)
}

func rangeSumBST2(root *TreeNode, L, R, sum int) int {
	if  root == nil {
		return sum
	}

	// 这个只是加上去
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	// 后面两个条件，获得的是加上去的 sum
	// 满足条件的才继续往左遍历
	if root.Val > L {
		sum = rangeSumBST2(root.Left, L, R, sum)
	}

	// 满足条件的才继续往右遍历
	if root.Val < R {
		sum = rangeSumBST2(root.Right, L, R, sum)
	}

	return sum
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if  root == nil {
		return 0
	}

	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}

	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
}
