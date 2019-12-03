package math

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("728. Self Dividing Numbers", func(t *testing.T) {
		left := 1
		right := 22
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}
		got := solution(left, right)
		fmt.Println(got)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(left, right int) []int {
	var dividedSlice []int
	for i := left; i < right+1; i++ {
		if allDivided(i) {
			dividedSlice = append(dividedSlice, i)
		}
	}
	return dividedSlice
}

func allDivided(i int) bool {
	j := i
	for j > 0 {
		digit := j % 10
		if digit == 0 {
			return false
		} else if i%digit != 0 {
			return false
		}
		j /= 10
	}
	return true
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
