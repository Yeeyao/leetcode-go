package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("34. Find First and Last Position of Element in Sorted Array", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		got := solution(nums, target)
		want := []int{3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
	t.Run("34. Find First and Last Position of Element in Sorted Array2", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 6
		got := solution(nums, target)
		want := []int{-1, -1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
	t.Run("34. Find First and Last Position of Element in Sorted Array3", func(t *testing.T) {
		nums := []int{}
		target := 0
		got := solution(nums, target)
		want := []int{-1, -1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
}

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

/*
	先是二分查找的代码，但是这里如果找到了，需要怎么处理
*/
func solution(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 找到的情况应该怎么处理？直接将 right = mid 让区间向左边收缩
		// 正确性证明 如果不存在，则 left == right 然后判断 nums[left] != target
		// 如果存在且只有一个，同样 left == right nums[left] == target
		// 如果存在且存在多个，则最终 left 和 right 在最左边的位置重合
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// [1] 1 问题，访问数组的时候需要注意越界问题
	// 已经找到了，则向后寻找右边界
	if nums[left] == target {
		for right < numsLen && nums[right] == target {
			right++
		}
		return []int{left, right - 1}
	}
	// 找不到，则直接返回
	return []int{-1, -1}
}
