package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("154. Find Minimum in Rotated Sorted ArrayII", func(t *testing.T) {
		nums := []int{3, 3, 3, 3, 4, 5, 1, 2}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("154. Find Minimum in Rotated Sorted ArrayII2", func(t *testing.T) {
		nums := []int{4, 5, 1, 1, 1, 1, 2, 2}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("154. Find Minimum in Rotated Sorted ArrayII3", func(t *testing.T) {
		nums := []int{4, 5, 5, 5, 6, 7, 0, 1, 2}
		want := 0
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("154. Find Minimum in Rotated Sorted ArrayII4", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 2, 3, -3, -3, -3, -2, -2, -2, -1, 0}
		want := -3
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("154. Find Minimum in Rotated Sorted ArrayII5", func(t *testing.T) {
		nums := []int{11, 11, 11, 11, 13, 15, 17}
		want := 11
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
相比 153 这里的元素是允许重复的，因此在进行二分查找之前，需要将重复的元素排除掉，这里类似 33 和 81 的关系

给定一个元素都是唯一的升序数组反转多次(长度 n，次数 1-n)，找到数组中最小的元素数值
这里反转的意思不是以一个元素为中心两边反转，而是反转一次就将尾部的元素放到头部
需要在 O(nlogn) 时间完成

应该是类似二分查找的思路，但是这里需要怎么判断？
可以类似 33 如果是左边部分就直接向右移动，如果是右边部分就直接向左移动。这里一开始的思路是对的。
然后，如果区间长度是 1 就是循环条件不成立了，二分查找就结束了，直接返回 left 或者 right


[ref](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/48493/Compact-and-clean-C%2B%2B-solution)
[ref](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/solution/er-fen-cha-zhao-wei-shi-yao-zuo-you-bu-dui-cheng-z/)
需要学习这里的分析方式
*/
func solution(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		// 如果最左边小于最右边，就根本没有反转过
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := left + (right-left)/2
		// 在左边，向右移动，否则向左移动
		if nums[mid] > nums[right] {
			left = mid + 1
			// 在右边，向左移动
		} else if nums[mid] < nums[right] {
			right = mid
			// 过滤相同的元素，直接将 right 递减，因为这里是使用 mid 和 right 比较
		} else {
			right--
		}
	}
	return nums[left]
}

func solution2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		// 如果最左边小于最右边，就根本没有反转过
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := left + (right-left)/2
		// 过滤相同的元素
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
			// 在左边，向右移动，否则向左移动
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
