package array

import (
	"testing"
)

/*
18. 4Sum
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums
such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
*/
func TestPro(t *testing.T) {
	t.Run("334. Increasing Triplet Subsequence", func(t *testing.T) {
		nums := []int{35, 15, 38, 1, 10, 26}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
If no such indices exists, return false.

给定整型数组 nums，返回是否存在索引的三元组 (i, j, k) 存在 i < j < k 同时 nums[i] < nums[j] < nums[k]
如果不存在则返回 false

第一时间想到单调栈，单调递增栈，有三个元素就可以直接返回 true

暴力就是选择一个元素作为中间的元素，然后向左边找小的元素，向右边找大的元素，但是这样时间复杂度 O(n^2)

主要是能否记录下大小的信息呢？比如每个元素保存其左边的当前最小值和右边的当前最大值，这样下一个元素。向两边移动就只需要判断旁边的元素的大小来更新

这里从左到右遍历一次，记录下每个元素左边的最小值，从右到左遍历一次，记录每个元素右边的最大值。计算完其中一个后，另一个可以在遍历的时候提前判断
*/

func solution(nums []int) bool {
	numsLen := len(nums)
	minMaxSlice := make([][2]int, numsLen)
	for i := 0; i < numsLen; i++ {
		if i == 0 || nums[i] < minMaxSlice[i-1][0] {
			minMaxSlice[i][0] = nums[i]
		} else {
			minMaxSlice[i][0] = minMaxSlice[i-1][0]
		}
	}
	for i := numsLen - 1; i >= 0; i-- {
		if i == numsLen-1 || nums[i] > minMaxSlice[i+1][1] {
			minMaxSlice[i][1] = nums[i]
		} else {
			minMaxSlice[i][1] = minMaxSlice[i+1][1]
		}
		// 这里可以提前判断
		if nums[i] > minMaxSlice[i][0] && nums[i] < minMaxSlice[i][1] {
			return true
		}
	}
	return false
}
