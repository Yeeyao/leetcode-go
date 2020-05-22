package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("496. Next Greater Element I", func(t *testing.T) {
		nums1 := []int{1, 3, 5, 2, 4}
		nums2 := []int{6, 5, 4, 3, 2, 1, 7}
		want := []int{-1, 3, -1}
		got := solution(nums1, nums2)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定两个没有重复元素的数组 nums1, nums2，其中 nums1 的元素是 nums2 元素的子集，
	找到 nums1 中所有元素在 nums2 的下一个较大元素，如果不存在则返回 -1
	brute force 对 nums1 每个元素都遍历 nums2 查找，这样时间复杂度将是 O(nm)
	使用单调递减栈
	使用 map 保存每个元素的下一个大于它的元素的位置
	遍历 nums2 如果栈非空且栈顶元素小于当前的元素，则保存下 {栈顶元素，当前元素}

	这种模拟方式有一个缺点是，如果给定的目标数据很大，可能无法分配那么多空间
*/
func solution(nums1, nums2 []int) []int {
	nums2Len := len(nums2)
	locMap := make(map[int]int, nums2Len)
	retArr := make([]int, len(nums1))
	// 栈中保存的是元素的值 主要注意栈顶的变化，因为循环后面将栈顶位置 + 1
	// 所以循环前访问栈顶元素的位置应该 - 1
	st := make([]int, nums2Len)
	stTop := 0
	for _, v := range nums2 {
		// 非空且当前元素比栈顶元素大，当前元素就是栈顶元素的后面第一个比它大的元素
		// 因为栈顶元素是先遍历到的 这里需要注意使用的是 for 只要元素不断大于栈顶
		// 就不断更新
		for stTop > 0 && v > st[stTop-1] {
			locMap[st[stTop-1]] = v
			stTop--
		}
		st[stTop] = v
		stTop++
	}
	for i, v := range nums1 {
		if l, ok := locMap[v]; ok {
			retArr[i] = l
		} else {
			retArr[i] = -1
		}
	}
	return retArr
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
