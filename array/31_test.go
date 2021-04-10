package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("31. Next Permutation", func(t *testing.T) {
		nums := []int{1, 2, 3}
		want := []int{1, 3, 2}
		solution(nums)
		if !reflect.DeepEqual(nums, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
}

/*
    i
12359123456789
12361234567899
[ref](https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/)
需要注意，这里是将后面的尽量小的大数和前面的小数交换，然后对大数之后的元素升序排列

从后面向前找到 第一个 nums[i] > nums[i-1] 同时 i 到数组结尾都是升序的
然后从 i 向后面找到最小的数字同时需要大于 nums[i-1]
这个数字跟 i 的数字交换后，i 后面的排序

1. 从后向前查找第一个相邻元素对(i,j) 满足 A[i] < A[j] 此时 j 到最后的元素都是降序的
2. 在 j 到最后的元素从后向前找到第一个满足 A[i] < A[k] 的 k 这里 A[k] 就是第一个大于 A[i] 的最小元素
3. A[i] 直接和 A[k] 交换
4. 之后 j 到最后的元素是降序的，直接反转它们
5. 如果 1 中找不到这样的相邻元素，直接 4

*/
func solution(nums []int) {
	numsLen := len(nums)
	// 这里比较最后两个元素
	i, k := numsLen-1, numsLen-1
	// 等于号的处理 从后面向前找第一个位置，当前元素大于上一个元素
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	// 将第一个元素小于后面的元素位置到结尾的所有元素位置反转
	// 如果是降序的话，i 已经是 0 同时这里将所有元素反转了
	for j := i; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
	// 不是降序排列 需要注意这里关键元素的位置
	if i > 0 {
		// 从 i 位置开始 即从大于上一个元素的位置开始
		k = i
		// 注意这里将 i 递减了才是 i 位置的元素原来是小于下一个元素的
		// 这里向后找第一个元素大于 i 位置的元素，并交换
		i--
		// 向后
		for nums[k] <= nums[i] {
			k++
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func nextPermutation(nums []int) {
	numsLen := len(nums)
	i, j, k := numsLen-2, numsLen-1, numsLen-1
	// 找到 nums[i] > nums[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	// 如果找到了，找最小的 nums[k] > nums[i] 然后交换
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 最后将 j 后面的降序
	for i, j := j, numsLen-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
