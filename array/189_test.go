package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("914. X of a Kind in a Deck of Cards", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		want := []int{5, 6, 7, 1, 2, 3, 4}
		solution(input, k)
		if !IntSliceEqual(input, want) {
			t.Errorf("got: %v, want: %v", input, want)
		}
	})

	t.Run("914. X of a Kind in a Deck of Cards2", func(t *testing.T) {
		input := []int{-1, -100, 3, 99}
		k := 2
		want := []int{3, 99, -1, -100}
		solution(input, k)
		if !IntSliceEqual(input, want) {
			t.Errorf("got: %v, want: %v", input, want)
		}
	})
}

/*
	翻转的方法
*/
func solution(input []int, k int) {
	inputLen := len(input)
	k = k % inputLen
	reverse(input[:inputLen-k])
	reverse(input[inputLen-k:])
	reverse(input)
}

func reverse(input []int) {
	inputLen := len(input)
	for i := 0; i < inputLen/2; i++ {
		input[i], input[inputLen-i-1] = input[inputLen-i-1], input[i]
	}
}

/*
	brute force
*/
func solution(input []int, k int) {
	inputLen := len(input)
	k = k % inputLen
	if k == 0 {
		return
	}
	arr := make([]int, inputLen)
	copy(arr, input)
	for i := 0; i < inputLen; i++ {
		input[(i+k)%inputLen] = arr[i]
	}
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
