package array

/*
34. Find First and Last Position of Element in Sorted Array
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
Your algorithm's runtime complexity must be in the order of O(log n).
If the target is not found in the array, return [-1, -1].

给定已排序数组，找到数组中目标元素出现的开始和结束位置
直接如果长度等于 0 返回 -1，-1
二分查找找到最左边的位置，如果该位置数值不等于 target 直接返回 -1,-1
否则，从左边的位置向右边查找，需要注意不要越界，找到右边的位置
*/
func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}
	left, right := 0, numsLen-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// left 是左边k开始 继续向右边找
	if nums[left] == target {
		right := left
		for right < numsLen && nums[right] == target {
			right++
		}
		return []int{left, right - 1}
	} else {
		// 找不到
		return []int{-1, -1}
	}
}
