package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("66. Plus One", func(t *testing.T) {
		input := []int{1, 3, 5, 4, 7}
		want := []int{1, 3, 5, 4, 8}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("66. Plus One2", func(t *testing.T) {
		input := []int{9, 9, 9, 9}
		want := []int{1, 0, 0, 0, 0}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("66. Plus One3", func(t *testing.T) {
		input := []int{0}
		want := []int{1}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) []int {
	inputLen := len(input)
	retArr := make([]int, inputLen)
	carry := 0
	last := inputLen - 1
	input[last] += 1
	for i := last; i >= 0; i-- {
		retArr[i] = input[i] + carry
		carry = retArr[i] / 10
		retArr[i] = retArr[i] % 10
	}
	if carry == 1 {
		retArr = append([]int{1}, retArr...)
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
