package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("167. Two Sum II - Input array is sorted", func(t *testing.T) {
		input := []int{2, 7, 11, 15}
		target := 9
		want := []int{1, 2}
		got := solution(input, target)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	用 map 保存另一个数
*/
func solution(input []int, target int) []int {
	inputLen := len(input)
	retArr := make([]int, 2)
	for i, j := 0, inputLen-1; i < j; {
		if input[i]+input[j] == target {
			retArr[0], retArr[1] = i+1, j+1
			return retArr
		} else if input[i]+input[j] > target {
			j--
		} else {
			i++
		}
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
