package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("28  数组中出现次数超过一半的数字", func(t *testing.T) {
		intSlice := []int{1, 2, 3, 4, 5, 6, 3, 4, 3, 4, 3, 3, 3, 3}
		get := solution(intSlice)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("28  数组中出现次数超过一半的数字2", func(t *testing.T) {
		intSlice := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
		get := solution(intSlice)
		want := 2
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

func solution(nums []int) int {
	Ele := nums[0]
	Num := 1
	numsLen := len(nums)
	for i := 1; i < numsLen; i++ {
		if nums[i] == Ele {
			Num++
		} else {
			Num--
		}
		if Num == 0 {
			Ele = nums[i]
			Num = 1
		}
	}
	return Ele
}
