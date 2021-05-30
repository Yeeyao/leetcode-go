package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("153. Find Minimum in Rotated Sorted Array", func(t *testing.T) {
		nums := []int{3, 4, 5, 1, 2}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("153. Find Minimum in Rotated Sorted Array2", func(t *testing.T) {
		nums := []int{4, 5, 1, 2}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("153. Find Minimum in Rotated Sorted Array3", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		want := 0
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("153. Find Minimum in Rotated Sorted Array4", func(t *testing.T) {
		nums := []int{1, 2, 3, -3, -2, -1, 0}
		want := -3
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("153. Find Minimum in Rotated Sorted Array5", func(t *testing.T) {
		nums := []int{11, 13, 15, 17}
		want := 11
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
给定一个元素都是唯一的升序数组反转多次(长度 n，次数 1-n)，找到数组中最小的元素数值
这里反转的意思不是以一个元素为中心两边反转，而是反转一次就将尾部的元素放到头部
需要在 O(nlogn) 时间完成

应该是类似二分查找的思路，但是这里需要怎么判断？
可以类似 33 如果是左边部分就直接向右移动，如果是右边部分就直接向左移动。这里一开始的思路是对的。
然后，如果区间长度是 1 就是循环条件不成立了，二分查找就结束了，直接返回 left 或者 right


[ref](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/48493/Compact-and-clean-C%2B%2B-solution)
[ref](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/solution/er-fen-cha-zhao-wei-shi-yao-zuo-you-bu-dui-cheng-z/)
需要学习这里的分析方式

用二分法查找，需要始终将目标值（这里是最小值）套住，并不断收缩左边界或右边界。
分析左中右三个位置数值，然后分为四种情况，其中情况中有最小值出现在左右半边，然后对应收缩左右边界
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
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
