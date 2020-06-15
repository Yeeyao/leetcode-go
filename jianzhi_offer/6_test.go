package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("6 旋转数组的最小数字", func(t *testing.T) {
		nums := []int{3, 4, 5, 1, 2}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("6 旋转数组的最小数字2", func(t *testing.T) {
		nums := []int{2, 2, 2, 0, 1}
		want := 0
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("6 旋转数组的最小数字3", func(t *testing.T) {
		nums := []int{2, 3, 4, 5, 1}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("6 旋转数组的最小数字4", func(t *testing.T) {
		nums := []int{5, 1, 1, 1, 2, 2, 2, 3, 4}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
3 4 5 6 1 2
i j m
0 5 2
3 5 4
mid = (i + j) // 2 因此可以知道 i <= m < j 这里判断 m 二分点在哪个部分
	nums[m] > nums[j] 就表示 m 在左边的子数组(m < j)，所求的元素在 m+1,j 之间，因此 i = m + 1
	nums[m] < nums[j] 则 m 在右边的子数组(因为出现在左边之前循环就应该终止了)，所求元素在 i,m 之间，执行 j = m
	nums[m] == nums[j] 无法判断，执行 j = j - 1 来缩小范围 这里有证明
i = j 跳出循环，返回 nums[i]
 */
func solution(nums[]int)int{
	numsLen := len(nums)
	i, j := 0, numsLen-1
	for i < j {
		// 计算中间索引
		m := i + (j-i)/2
		if nums[m] > nums[j] {
			i = m + 1
		} else if nums[m] < nums[j] {
			j = m
			// 这里处理相等的情况
		} else {
			j = j - 1
		}
	}
	return nums[i]
}

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
这里是找出最小值 类似 81
虽然自己写出来了，但是还没有弄懂，尴尬 TODO:
显然这里是寻找右边子数组的第一个开始元素，使用二分法
同样 left <= mid < right
nums[left] < nums[mid] 就表示 left 一定在左边的子数组中，所求的元素出现在 mid, right 之间，所以需要将 left + 1
nums[left] > nums[mid] 就表示 left 一定在右边的子数组中，所求元素出现在
*/
//func solution2(nums []int) int {
//	numsLen := len(nums)
//	left, right := 0, numsLen-1
//	for left < right {
//		// 计算中间索引
//		mid := left + (right-left)/2
//		// 过滤相等的元素
//		for nums[left] == nums[mid] {
//			left++
//		}
//		for nums[right] == nums[mid] {
//			right--
//		}
//		// 还在左边部分 只要小于就不断从中间向前一索引
//		if nums[left] < nums[mid] {
//			left = mid + 1
//			// 在右边部分
//		} else {
//			right = mid
//		}
//	}
//	// 升序后的最小元素比较
//	return nums[right]
//}
