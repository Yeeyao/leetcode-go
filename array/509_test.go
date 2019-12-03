package array

import "testing"

func TestPro(t *testing.T) {

	t.Run("leetcode 238  Product of Array Except Self", func(t *testing.T) {
		input := 4
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input int) int {
	if input == 0 {
		return 0
	}
	if input == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 1; i < input; i++ {
		a, b = b, a+b
	}
	return b
}
