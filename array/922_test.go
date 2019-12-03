package array

import "testing"

func TestPro(t *testing.T) {

	t.Run("leetcode 238  Product of Array Except Self", func(t *testing.T) {
		input := []int{4, 2, 5, 7}
		want := []int{24, 12, 8, 6}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) []int {
	inputLen := len(input)
	retProduct := make([]int, inputLen)
	evenCounter := 0
	oddCounter := 1
	for i := 0; i < inputLen; i++ {
		if input[i]%2 == 0 {
			retProduct[evenCounter] = input[i]
			evenCounter += 2
		} else {
			retProduct[oddCounter] = input[i]
			oddCounter += 2
		}
	}
	return retProduct
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
