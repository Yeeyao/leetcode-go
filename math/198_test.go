package math

import "testing"

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that adjacent houses have security system connected
and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

抢劫房子，不能连续抢劫两个相邻的房子，给定房子可以抢劫到的钱数量列表，求最大可以抢劫到的数量
其实只有两种，单数或者偶数和然后取较大值，因为这里的数值都大于 0
不对，可以中间的不抢，所以应该要 DFS
	每个元素都作为开始进行递归遍历，然后后续跳过下一个元素直接后面的元素都遍历
	然后不相邻的都要处理

	helper 中需要先加上 begin 的数值，然后判断 begin 是最后两个就直接判断最大值并返回
*/

func TestPro(t *testing.T) {
	t.Run(" 198. House Robber ", func(t *testing.T) {
		input := []int{1, 2, 3, 1}
		want := 4
		got := rob(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 198. House Robber2", func(t *testing.T) {
		input := []int{2, 7, 9, 3, 1}
		want := 12
		got := rob(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 198. House Robber3", func(t *testing.T) {
		input := []int{2}
		want := 2
		got := rob(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 198. House Robber4", func(t *testing.T) {
		input := []int{2, 7}
		want := 7
		got := rob(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	TLE
*/
func rob(nums []int) int {
	sum := 0
	numsLen := len(nums)
	// 0, 4
	for i := 0; i < numsLen; i++{
		helper(nums, i, numsLen, 0, &sum)
	}
	return sum
}

func helper(nums []int, begin, end, temp int, sum *int){
	temp += nums[begin]
	if begin >= end - 2 {
		if temp > *sum{
			*sum = temp
		}
		return
	}
	for i := begin + 2; i < end; i++{
		helper(nums, i, end, temp, sum)
	}
}

/*
	用迭代
	如果只有一间房屋，则偷窃该房屋，可以偷到最高总金额
	如果有两间，则选择金额较高的偷窃
	如果数量大于 2 对于第 k 间房屋，有两个选项
		偷窃第 k 间房屋，就不能偷 k - 1 的，偷窃总额为前  k - 2 间的最高金额和第 k 间的金额
		不偷窃第 k 间房屋，则偷窃总金额为前 k - 1 间房屋的最高总金额
		上述两个选项中选择较大的

	dp[i] 表示前 i 间房屋可以偷盗的最高总金额，则有转移方程
	dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
	最终结果是 dp[n - 1]
*/
func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1{
		return nums[0]
	}
	dp := make([]int, numsLen)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < numsLen; i++{
		dp[i] = max((dp[i-2] + nums[i]), dp[i-1])
	}
	return dp[numsLen - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}