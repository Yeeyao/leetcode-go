package array

/*
416. Partition Equal Subset Sum
LEETCODE 494

Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets
such that the sum of elements in both subsets is equal.
Note:
    Each of the array element will not exceed 100.
    The array size will not exceed 200.

给定一个仅含有正整数非空数组，判断是否可以被划分为两个子集，两个子集的总和相等
跟 + - 的很像 第一部分 x，第二部分 y s = 0
x + y = sum x - y = s x = (sum + s) / 2 x = sum / 2
先求和，然后判断是否可以满足数量条件 dp[i] 表示空间为 i 背包能否被存放
然后直接背包问题先创建 dp 长度为 x + 1，然后每个元素都遍历，每个元素进行 dp 计算
初始值 dp[0] = 1
*/
func canPartition(nums []int) bool {
	numsLen := len(nums)
	if numsLen == 0 {
		return false
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	// 无法组成整数
	if (sum & 1) == 1 {
		return false
	}
	xNum := sum / 2
	// 初始化全都是 false
	dp := make([]bool, xNum+1)
	dp[0] = true
	for _, n := range nums {
		for i := xNum; i >= n; i-- {
			dp[i] = dp[i] || dp[i-n]
		}
	}
	return dp[xNum]
}
