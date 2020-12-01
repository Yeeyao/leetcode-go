package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("674. Longest Continuous Increasing Subsequence", func(t *testing.T) {
		input := []int{1, 3, 5, 4, 7}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("674. Longest Continuous Increasing Subsequence2", func(t *testing.T) {
		input := []int{2, 2, 2, 2, 2, 2, 2}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("674. Longest Continuous Increasing Subsequence3", func(t *testing.T) {
		input := []int{1, 3, 5, 7}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	类似 830 的思路
*/
func solution(input []int) int {
	inputLen := len(input)
	if inputLen == 0 {
		return 0
	}
	maxLen := 1
	for i := 1; i < inputLen; i++ {
		tempLen := 1
		for i < inputLen && input[i] > input[i-1] {
			tempLen++
			i++
		}
		if tempLen > maxLen {
			maxLen = tempLen
		}
	}
	return maxLen
}
