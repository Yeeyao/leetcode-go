package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("50 数组中重复的数字", func(t *testing.T) {
		nums := []int{2, 3, 1, 0, 2, 5, 3}
		get := solution(nums)
		want := 2
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
	直接用 map
*/
func solution(nums []int) int {
	intMap := make(map[int]bool)
	for _, n := range nums {
		if _, ok := intMap[n]; ok {
			return n
		}
		intMap[n] = true
	}
	return -1
}
