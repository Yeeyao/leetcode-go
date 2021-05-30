package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("33. Search in Rotated Sorted Array", func(t *testing.T) {
		nums := []int{1, 2, 3}
		target := 3
		got := solution(nums, target)
		want := 2
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
	t.Run("33. Search in Rotated Sorted Array2", func(t *testing.T) {
		nums := []int{1, 0, 2, 3, 4}
		target := 3
		got := solution(nums, target)
		want := 3
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
}

/*
	3 4 5 1 2
	4 5 6 1 2 3
	首先计算二分的中点，然后判断目标值和中间值关系同时确定中间值的位置。这里不是需要移动，而是需要重新设置左右两个边界的位置
		如果中间值大于目标值，则需要找更小的中间值，如果在左边部分，需要将右边设置为中间值 + 1，或者将左边设置为中间值；如果在右边部分，需要将右边设置为中间值 + 1
			如果在左边部分，然后将右边设置为中间值则会一直在左边寻找(如果数值在右边则有问题)，将左边设置为中间值，则一直在右边寻找（如果数值在左边则有问题）
		如果中间值小于目标值，则需要找更大的中间值，如果在左边部分，将左边设置为中间值，如果在右边部分，将左边设置为中间值或者需要将右边设置为中间值 + 1
			这里能否只向右移动
	按照上面的分析，还是没有办法最终确定应该怎么移动？再想
	所以上面的思路不对

	如果 mid 数值 大于等于 nums[0] ，则表示当前在递增数组的左边部分（较大的部分），如果小，则在右边部分（较小的部分）

	[参考这里的二分查找](https://www.zhihu.com/question/36132386)
		for left < right
		right = mid, left = mid + 1
		return left

	我们将数组从中间分开左右两部分的时候，一定有一部分数组是有序的，这里可以用反证法证明
	使用常规二分查找的时候，查看当前 mid 分割位置分割处理两个部分 [l, mid] [mid + 1, r] 哪个部分是有序的，根据有序的部分确定如何改变二分查找的上下界
	因为我们可以根据有序的部分判断 target 是否存在于这个部分
		如果 [l, mid - 1] 是有序数组，target 在这个范围内，则需要在这个范围内搜索。否则在 [mid + 1, r] 中搜索
		如果 [mid, r] 是有序数组，target 在这个范围内，则需要在这个范围内搜索。否则在 [l, mid - 1] 中搜索
		这里只是需要在某一边查找，所以直接将边界调整
	同时需要注意等号
*/
func solution(nums []int, target int) int {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 左边有序
		if nums[left] <= nums[mid] {
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
		return left
	} else {
		return -1
	}
}
