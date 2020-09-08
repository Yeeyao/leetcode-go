package array

/*
55. Jump Game
[ref](https://leetcode-cn.com/problems/jump-game/solution/tiao-yue-you-xi-by-leetcode-solution/)
Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Determine if you are able to reach the last index.
给定一个非负整型数组，从第一个元素开始，每个元素表示可以跳的最大步数，求是否可以跳到最后位置
3 2 1 0 4 第一个是 3 表示最大可以从到达 0 元素的位置

贪心
对任意位置 y 如果存在一个位置 x，x + num[x] >= y，则 y 可以到达
可以依次遍历数组每个位置，实时维护最远可以到达的位置，对当前遍历到的位置 x，
如果它在最远可到位置的范围内，则可以到达这个位置，可以用 x + nums[x] 来更新最远可达位置
遍历的过程中，如果最远可达位置大于等于数组最后位置，则说明最后一个位置可达

初始化 maxLen = 0, numsLen 为数组长度
循环遍历每个元素
	如果当前元素小于等于 maxLen
		用 i + nums[i] 更新 maxLen
		如果 maxLen >= numsLen - 1 返回 true
最后返回 false
*/
func canJump(nums []int) bool {
	numsLen := len(nums)
	maxLen := 0
	for i := 0; i < numsLen; i++ {
		if i <= maxLen {
			maxLen = max(i+nums[i], maxLen)
			// 提前结束
			if maxLen >= numsLen-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
