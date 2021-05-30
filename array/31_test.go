package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("31. Next Permutation", func(t *testing.T) {
		nums := []int{1, 2, 3}
		want := []int{1, 3, 2}
		nextPermutation(nums)
		if !reflect.DeepEqual(nums, want) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
}

func nextPermutation(nums []int) {
	numsLen := len(nums)
	i, j, k := numsLen-2, numsLen-1, numsLen-1
	// 从后面向前找到 nums[i] < nums[j] 第一个元素比后面元素小的位置，同时 j 后面的元素都是降序的
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	// 如果找到了，从后面找第一个 nums[k] > nums[i] 然后交换，因为 j 后面的元素是降序的，因此这里也是找到最小的大于 i 元素
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 最后将 j 后面的降序，满足排列
	for i, j := j, numsLen-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

/*
[ref](https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/)
需要注意，这里是将后面的尽量小的大数和前面的小数交换，然后对大数之后的元素升序排列

12534321 12541233
从后面向前找到 第一个 nums[i] > nums[i-1](等号需要继续向前找) 此时 i 到数组结尾都是降序的
然后从 i 向后面找到最小的数字同时需要大于 nums[i-1](因为第一个位置的不一定是后面最小的数字)
这个数字跟 i 的数字交换后，i 后面的排序

1. 从后向前 查找第一个相邻元素对(i,j) 满足 A[i] < A[j] 此时 j 到最后的元素都是降序的
2. 在 j 到最后的元素从后向前找到第一个满足 A[i] < A[k] 的 k 这里 A[k] 就是第一个大于 A[i] 的最小元素
3. A[i] 直接和 A[k] 交换
4. 之后 j 到最后的元素是降序的，直接反转它们
5. 如果 1 中找不到这样的相邻元素，直接 4

*/
func solution(nums []int) {
	numsLen := len(nums)
	// 这里比较最后两个元素
	i, k := numsLen-1, numsLen-1
	// 等于号的处理 从后面向前找 第一个位置，该位置比上一个位置的元素大
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	// 将前面找到的元素位置到结尾的所有元素位置反转(将原来的后面的降序数组变成升序的)
	// 如果是降序的话，i 已经是 0 同时这里将所有元素反转了，可以直接返回
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
		for nums[k] <= nums[i] {
			k++
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}
