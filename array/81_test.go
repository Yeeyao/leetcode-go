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
为什么这样可以过滤？，为什么不能直接使用 nums[0] 而需要改成 nums[left]

这里在一个位置反转(这里位置是所有的位置都可以)。将原来的单调递增的数组，变成两个单调递增的数组，其中中间的位置是数组元素的最大值
这里 mid 的计算，同样遵循 left <= mid mid < right

可能出现的情况，中值出现在左边的递增数组还是右边的递增数组判断
	1. 出现在左边，判断 target 和中值的关系
		1.1 target > nums[mid] 需要向右边收缩
		1.2 target < nums[mid] 需要向左边收缩或者向右边收缩
	2. 出现在右边，判断 target 和中值的关系
		2.1 target > nums[mid] 需要向右边收缩
		2.2 target < nums[mid] 需要向左边收缩
这里 1.2 的情况就不清晰，所以需要思考其他划分方式

可能的情况，直接判断左边和中间的数值关系
	1.左边数值小于等于中间数值（当前左边和中间数值都在左半递增数组，成立，都在右半递增数组，成立。左边在左半部分，中间在右半部分，不成立，所以可以确定是在一个递增数组的内部）
		1.1 如果 target > nums[mid] 则需要向右边收缩 left = mid + 1
		1.2 如果 target <= nums[mid] 则需要向左边收缩 right = mid
	2.左边数值大于中间数值（左边出现在左半的递增数组，右边出现在右半的递增数组），0, mid 中包含了旋转的位置
		2.1 target <= nums[mid] < nums[left]，需要向左边收缩
			因为 left 左边的部分比 mid 左边的部分都大，因此 target 在旋转点到 mid 之间(右边有序数组前面部分)
		2.2 nums[mid] < nums[left] <= target，需要向左边收缩
			只能向 left 的右边到 mid 之间（左边数组的 left 右边找）target 在 left 到 旋转点之间（左边有序数组后面部分）
		2.3 其他情况，向右边收缩
			这里主要的疑问就是剩下的 nums[mid] < target <  nums[left]，那一定是向右收缩呀，
			因为，nums[left] 的右边到旋转点的数值比 nums[left] 大，所以不可能在左边数组的右半部分(那左半部分呢？那不可能，因为是从左边收缩的)；
			同时，target > nums[mid]，所以也不可能在右边数组的左半部分，因此就只能在右边数组的右半部分寻找了，因此向右收缩

	最终的解法将上述的情况进行汇总，其中，用来判断究竟应该向哪边移动的，最初的方法都是需要 target 在一个递增的数组里面。

[ref](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-by-leetcode-solut/)
官方的做法是，因为某个位置旋转，会导致数组被分为左右两个单独递增的子数组。
这里去判断 l, mid 和 mid + 1, r 哪部分是有序的，根据有序的部分确定上下界的变化（因为可以根据有序的部分判断 target 是否在这个部分）
	如果 l, mid 有序，target 大小在 nums[l] nums[mid] 之间，则我们将收缩右边(但是其实右边是否存在 target 呢？)，target 大小不在这个之间，则收缩左边
	如果 mid, r 有序，target 大小在 nums[mid] nums[r] 之间，则我们将收缩左边(但是其实左边是否存在 target 呢？)，target 大小不在这个之间，则收缩右边
单独看，确实是，只有在有序的部分才能进一步判断 target 是否存在于有序的部分，如果大小属于有序的部分，就直接在该部分查找。

[ref](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/ji-jian-solution-by-lukelee/)
这里分析是
	1.2 2.1 2.2 都是向左边收缩，其他情况向右边收缩。
*/
func solution(nums []int, target int) bool {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		// 比 33 多了一步如果有相等的元素，先过滤 how it works，这里的意思是，三个相等就没法判断哪边是有序
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
