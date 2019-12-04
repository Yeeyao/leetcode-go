package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("119. Pascal's Triangle II", func(t *testing.T) {
		input := 3
		want := []int{1, 3, 3, 1}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input int) []int {
	tempArr := make([]int, input+1)
	retArr := make([]int, input+1)
	for i, _ := range tempArr {
		tempArr[i] = 1
		retArr[i] = 1
	}
	isTemp := true
	for i := 0; i < input+1; i++ {
		for j := 1; j < i; j++ {
			if isTemp {
				retArr[j] = tempArr[j-1] + tempArr[j]
			} else {
				tempArr[j] = retArr[j-1] + retArr[j]
			}
		}
		if isTemp {
			isTemp = false
		} else {
			isTemp = true
		}
	}
	if isTemp {
		return tempArr
	} else {
		return retArr
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
