package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("989. Add to Array-Form of Integer", func(t *testing.T) {
		input := []int{1, 2, 0, 0}
		k := 34
		want := []int{1, 2, 3, 4}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer2", func(t *testing.T) {
		input := []int{2, 7, 4}
		k := 181
		want := []int{4, 5, 5}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer3", func(t *testing.T) {
		input := []int{2, 1, 5}
		k := 806
		want := []int{1, 0, 2, 1}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer4", func(t *testing.T) {
		input := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
		k := 1
		want := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer5", func(t *testing.T) {
		input := []int{2, 1, 5}
		k := 8061
		want := []int{8, 2, 7, 6}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int, k int) []int {
	inputLen := len(input)
	kLen := calcLen(k)
	maxLen := 0
	if inputLen > kLen {
		maxLen = inputLen + 1
	} else {
		maxLen = kLen + 1
	}

	retArr := make([]int, maxLen)
	carry := 0
	for i := 0; i < maxLen; i++ {
		kAdd := k % 10
		k /= 10
		ri := maxLen - i - 1
		j := inputLen - i - 1
		if j >= 0 {
			if carry == 1 {
				retArr[ri] = input[j] + kAdd + 1
			} else {
				retArr[ri] = input[j] + kAdd
			}
		} else {
			retArr[ri] = kAdd + carry
		}

		if retArr[ri] >= 10 {
			carry = 1
			retArr[ri] -= 10
		} else {
			carry = 0
		}
	}

	if retArr[0] == 0 {
		retArr = retArr[1:]
	}
	return retArr
}

func calcLen(k int) int {
	kLen := 1
	if k/10 > 0 {
		kLen++
		k /= 10
	}
	return kLen
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
