package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("219. Contains Duplicate II", func(t *testing.T) {
		input := []int{1, 2, 3, 1}
		k := 3
		want := true
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("219. Contains Duplicate II 2", func(t *testing.T) {
		input := []int{1, 0, 1, 1}
		k := 1
		want := true
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("219. Contains Duplicate II 3", func(t *testing.T) {
		input := []int{1, 2, 3, 1, 2, 3}
		k := 2
		want := false
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int, k int) bool {
	posArr := make(map[int]int)
	for i, v := range input {
		if w, ok := posArr[v]; ok {
			if i-w <= k {
				return true
			}
		}
		posArr[v] = i
	}
	return false
}
