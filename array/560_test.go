package array

import (
	"testing"
)

/*
560. Subarray Sum Equals K 类似 494 454 1 1248
Given an array of integers and an integer k, you need to find the total number of continuous subarrays
whose sum equals to k.
Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Constraints:
    The length of the array is in range [1, 20,000].
    The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].
给定要给整型数组和 k 值，需要找到数组中连续元素和等于 k 的子数组的数量
数组长度是 1 到 20000 元素范围 -1000,1000 k -1e7 1e7
*/

func TestPro(t *testing.T) {
	t.Run("560. Subarray Sum Equals K", func(t *testing.T) {
		input := []int{1, 1, 1}
		k := 2
		want := 2
		got := subarraySum(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
[ref](https://leetcode-cn.com/problems/subarray-sum-equals-k/solution/he-wei-kde-zi-shu-zu-by-leetcode-solution/)
	前缀和 哈希表优化 瓶颈在于对于每个元素，需要枚举后面所有的元素求和来判断
	计算前缀和，然后用 map 保存每个前缀和出现的次数
	前缀和数组是 pre，pre[i] = nums[0] + ... + nums[i] pre[i] = pre[i-1] + nums[i]
	i 到 j 的和就是 pre[j] - pre[i - 1]，目标是求 pre[j] - pre[i-1] == k 组合数量
	即 pre[i-1] = pre[j] - k 题目转换为考虑以 j 结尾和为 k 的连续子数组个数就统计
	有多少个前缀和为 pre[j] - k 的 pre[i]

	初始化 count, pre map m 同时 m[0] = 1
	遍历元素，先将累加和 pre 加上当前元素，然后判断 m[pre-k] 是否存在(pre - (pre -k) = k)，
	是则将计数 + 1。然后当前的累加和计数 + 1
	就类似 two-sum 只不过是累加然后统计累加和的数量，然后找差值的出现数量

	需要理解的是，这里 pre 是从前面到后面累加的，所以，对于当前的 pre，只需要向前找
*/
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	// 和为 0 的有一个
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		// 先加上当前元素统计累加和
		pre += nums[i]
		// 是否找到对应的差值
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		// 将当前的累加和数量 + 1
		m[pre] += 1
	}
	return count
}

/*
	一每个元素作为开始 但是这样是 O(n^2)
*/
func subarraySum2(nums []int, k int) int {
	count := 0
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		temp := 0
		for j := i; j < numsLen; j++ {
			temp += nums[j]
			if k == temp {
				count++
			}
		}
	}
	return count
}
