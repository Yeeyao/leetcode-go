package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("31. Next Permutation", func(t *testing.T) {
		reservedSeats := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}
		n := 3
		want := 4
		got := solution2(n, reservedSeats)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找到下一个排列
	从后面向前找第一个位置，当前元素小于下一个元素
	将上面的位置到结尾位置的所有元素都交换
	如果上面的位置大于 0 表示元素不是降序排列的
	则需要从 上面的位置向后面找到第一个大于该位置的数然后交换两者
	需要注意等于的情况
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
