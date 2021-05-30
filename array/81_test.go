package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("81. Search in Rotated Sorted Array II", func(t *testing.T) {
		nums := []int{2, 5, 6, 0, 0, 1, 2}
		target := 0
		want := true
		got := solution(nums, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("81. Search in Rotated Sorted Array II2", func(t *testing.T) {
		nums := []int{1, 0, 1, 1, 1}
		target := 0
		want := true
		got := solution(nums, target)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	对比 33，这里的元素不是唯一的，因此在判断前进行元素过滤，同时只需要返回是否存在
*/
func solution(nums []int, target int) bool {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		// 比 33 多了一步如果有相等的元素，先过滤
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
			// 左边有序 33 的 0 需要修改成 left TODO: 为什么？
		} else if nums[left] <= nums[mid] {
			// 如果 target 在左边范围，在左边查找，这里没有等号是因为前面已经判断了
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				// target 在右边范围，在右边查找
				left = mid + 1
			}
		} else {
			// 右边有序
			// 如果 target 在右边范围，在右边查找
			if nums[mid] < target && target <= nums[numsLen-1] {
				left = mid + 1
			} else {
				// target 在左边范围，在左边查找
				right = mid
			}
		}
	}
	if left == right && nums[left] == target {
		return true
	} else {
		return false
	}
}

/*
	[ref](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/solution/zai-javazhong-ji-bai-liao-100de-yong-hu-by-reedfan/)
	已排序数组在某个位置进行了旋转，给的一个元素判断数组中是否包含该元素
	二分查找？同时需要判断当前所在的部分
	主要是观察元素，什么条件下可以判断

	每次检查中间部分是否等于索要找的
	检查左边部分是否已经排序，即 nums[left] <= nums[mid]
		如果是则第三步，否则第四步
	检查所找的是否在 left 到 mid - 1 之间，即 nums[left] < target < nums[mid - 1]
		如果是就在左边部分查找，令 right = mid - 1
	检查是否在	mid + 1 到 right 之间，如果是就在右边部分查找，令 left = mid + 1
*/
func solution2(nums []int, target int) bool {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	// 注意这里要处理等于
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		// 过滤一下相等的元素
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// 判断哪部分已经排序 需要注意等号的边界 左边部分已排序
			// 这里两边需要判断相等闭区间，中间的开区间
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右边部分已排序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return false
}
