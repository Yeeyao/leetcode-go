package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("45 扑克牌的顺子", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		get := solution(nums)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("45 扑克牌的顺子2", func(t *testing.T) {
		nums := []int{0, 0, 1, 2, 5}
		get := solution(nums)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("45 扑克牌的顺子3", func(t *testing.T) {
		nums := []int{1, 0, 1, 2, 5}
		get := solution(nums)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
	2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。
	A 不能视为 14。

	要求除了大小王，其他牌不重复 同时 max - min < 5
	大小王直接跳过统计，其他需要统计出现次数以及最大最小值
*/
func solution(nums []int) bool {
	meet := make(map[int]bool)
	min, max := 15, 0
	for _, n := range nums {
		if n == 0 {
			continue
		}
		// 出现重复值
		if _, ok := meet[n]; ok {
			return false
		}
		meet[n] = true
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
		if max-min > 5 {
			return false
		}
	}
	return true
}
