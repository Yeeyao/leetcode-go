package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("628. Maximum Product of Three Numbers", func(t *testing.T) {
		input := []int{-4, -3, -2, -1, 60}
		want := 720
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("628. Maximum Product of Three Numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := 24
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里只需要求乘积
	或者求最大的三个正数以及最小的两个负数
	可以排序后处理
*/
func solution(input []int) int {
	max1, max2, max3 := -1001, -1001, -1001
	min1, min2 := 1001, 1001
	for _, v := range input {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	product1 := min1 * min2 * max1
	product2 := max1 * max2 * max3
	if product1 > product2 {
		return product1
	} else {
		return product2
	}
}
