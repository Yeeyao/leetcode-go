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

分析左中右三个位置数值，然后分为四种情况，其中情况中有最小值出现在左右半边，然后对应收缩左右边界

详细分析，这里入手的地方是左中右值的位置比较情况
用二分法查找，需要始终将目标值（这里是最小值）套住，并不断收缩左边界或右边界。
下面四种情况其实是 2 * 2，二分查找的一般性情况
	1. 左值 < 中值 中值 < 右值：没有旋转，最小值在最左边，可以收缩右边界
	2. 左值 > 中值 中值 < 右值：发生了旋转，最小值在左半边(或者中间位置)，可以收缩右边界
	3. 左值 < 中值 中值 > 右值：发生了旋转，最小值在右半边(或者中间位置)，可以收缩左边界
	4. 左值 > 中值 中值 > 右值：不会发生

可以知道 1，2 是同一类；3 是另一类。这里通过向那边收缩了来分类。然后中值的判断是和收缩的一致
	如果中值 < 右值，最小值在左半边，可以收缩右边界
	如果中值 > 右值，最小值在右半边，可以收缩左边界
1，3 都是 左值 < 中值，但是最小值范围不同，如果使用左值和中值比较则不能确定最小值的范围，但是提交的确实是使用左值和中值比较...

循环不变式 left < right，中间 mid 的计算让 mid 更靠近 left，因此有 left <= mid mid < right
	while 循环内，mid 始终小于 right
	也正是因为这样，提交的可以使用左值和中值比较，但是实际的分析是需要按照上面的来
while 循环退出条件
	如果输入数组只有一个数，左右边界位置重合，直接输出
	如果多于一个数，循环到最后，只会剩下两个数(为什么，因为更多会继续移动)nums[left] == nums[mid] 和 nums[right] 这里 left == mid == right - 1
		如果 nums[left] == nums[mid] > nums[right] 则左边大，右边小，将执行 left = mid + 1，则 left == right 结束时 left, right 保存最小值
		如果 nums[left] == nums[mid] < nums[right] 则右边大，左边小，将执行 right = mid，则 left == right，三个都保存了最小值

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
		} else {
			right = mid
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
		// 在左边，向右移动，否则向左移动
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
