package array

/*
[ref](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-by-leetcode-solut/)
33. Search in Rotated Sorted Array 类似 81
You are given an integer array nums sorted in ascending order, and an integer target.
Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
If target is found in the array return its index, otherwise, return -1.
在旋转的已排序数组中查找元素，找不到返回 -1 使用二分查找

从中间 mid 分开时，一定有部分数据时有序的，左边部分 l, mid 右边 mid + 1, r
	如果 l, mid - 1 是有序，且 target 大小满足 [nums[l], nums[mid])，
	则我们应该将范围缩小到 l, mid - 1 否则在 mid + 1, r 中找
	如果 mid, r 是有序的，且 target 大小满足 (nums[mid], nums[r]],
	则我们将范围缩小到 mid + 1, r 否则在 l, mid - 1 中找

先判断长度为 0 和 1 的情况
初始化左右两个指针
循环条件是 left <= right
	先判断中间元素是否是所求
	判断左右两边那边是已排序的
		如果左边已排序 这里使用 nums[0] 和 nums[mid] 进行比较判断那边已排序
			如果 target 范围在左边则将 right = mid - 1 否则 left = mid + 1
		如果右边已排序
			如果 target 范围在右边则将 left = mid + 1 否则 right = mid - 1
*/
func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, numsLen-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}
		// 左边部分已排序
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
