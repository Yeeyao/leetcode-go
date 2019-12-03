package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("1089. Duplicate Zeros", func(t *testing.T) {
		input := []int{1, 0, 2, 3, 0, 4, 5, 0}
		want := []int{1, 0, 0, 2, 3, 0, 0, 4}
		solution(input)
		fmt.Println(input)
		if !IntSliceEqual(input, want) {
			t.Errorf("got: %v, want: %v", input, want)
		}
	})
}

func solution(arr []int) {
	arrLen := len(arr)
	arr2 := make([]int, arrLen)
	copy(arr, arr2)
	for i, j := 0, 0; j < arrLen; i++ {
		fmt.Println(i)
		if arr2[i] == 0 {
			if j+1 < arrLen {
				arr[j] = 0
				arr[j+1] = 0
				j += 2
			} else {
				return
			}
		} else {
			arr[j] = arr2[i]
			j++
		}
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
