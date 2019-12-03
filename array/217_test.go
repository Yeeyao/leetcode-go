package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("217. Contains Duplicate", func(t *testing.T) {
		input := []int{4, 3, 2, 7, 8, 2, 3, 1}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) bool {
	inputLen := len(input)
	countArr := make(map[int]bool, inputLen)
	for _, v := range input {
		if _, ok := countArr[v]; ok {
			return true
		} else {
			countArr[v] = true
		}
	}
	return false
}
