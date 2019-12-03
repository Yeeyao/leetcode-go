package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("169. Majority Element", func(t *testing.T) {
		input := []int{3, 2, 3}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("169. Majority Element2", func(t *testing.T) {
		input := []int{2, 2, 2, 1, 3, 4}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) int {
	majNum := input[0]
	majNumCount := 0
	for _, v := range input {
		if v == majNum {
			majNumCount++
		} else {
			if majNumCount == 0 {
				majNum = v
				majNumCount = 1
			} else {
				// 这里只是减少，需要到下一个循环才会替换
				majNumCount--
			}
		}
	}
	return majNum
}
