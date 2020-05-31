package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1130. Minimum Cost Tree From Leaf Values", func(t *testing.T) {
		nums := []int{6, 2, 4}
		want := 32
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定数组，这里二叉树的定义：
	每个节点有 0/2 个子节点
	arr 中的值是一个中序遍历的树的叶子节点的数值
	每个非叶子节点的值是其左右子树的分别的最大叶子节点的数值的乘积
	计算给定的 arr 组成的上述二叉树最小非叶子节点的和

	dp 找到 (i, j) 区间的 cost，为了建立 [i, j]，
	我们需要将其划分为左右子树
	这里的和等于左右两个子树的数值加上当前节点的左右两个叶子节点的最大的乘积
	dp[i, j] = dp[i, k] + dp[k + 1, j] + max(A[i, k]) * max(A[k + 1, j])
	但这里是 brute force
	[lee](https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/339959/One-Pass-O(N)-Time-and-Space)

	stack solution 在左边和右边找到数组中的下一个较大的元素
	使用单调递减栈 需要先保存一个最大值在栈中来计算
	遍历数组，然后当前元素大于等于栈顶的话，需要将栈顶移除同时计算
	当前元素入栈，最后，当栈元素数量大于 2，则需要继续计算乘积
	主要难点在于如何将问题转换
*/
func solution(arr []int) int {
	arrLen := len(arr)
	MAX_INT := int(^uint(0) >> 1)
	// 这里栈保存的是数值
	st := make([]int, arrLen)
	// 先保存一个最大的数值
	stTop := 0
	res := 0
	st[stTop] = MAX_INT
	stTop++
	for i := 0; i < arrLen; i++ {
		// 当前元素大于栈顶 构建单调递减栈
		for st[stTop-1] <= arr[i] {
			if st[stTop-2] < arr[i] {
				res += st[stTop-1] * st[stTop-2]
			} else {
				res += st[stTop-1] * arr[i]
			}
			// 栈顶元素出栈
			stTop--
		}
		// 当前元素入栈
		st[stTop] = arr[i]
		stTop++
	}
	stLen := len(st)
	// 求和
	for i := 2; i < stLen; i++ {
		res += st[i] * st[i-1]
	}
	return res
}
