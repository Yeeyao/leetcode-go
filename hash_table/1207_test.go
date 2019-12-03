package hash_table

import "testing"

func TestPro(t *testing.T) {
	t.Run("1207. Unique Number of Occurrences", func(t *testing.T) {
		input := []int{1, 2, 2, 1, 1, 3}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1207. Unique Number of Occurrences2", func(t *testing.T) {
		input := []int{1, 2}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// 需要全部都统计出来，之后再判断出现次数的独立
func solution(A []int) bool {
	numCounter := make(map[int]int)
	for _, val := range A {
		numCounter[val]++
	}
	numbers := make(map[int]bool)
	for _, val := range numCounter {
		if _, ok := numbers[val]; ok {
			return false
		} else {
			numbers[val] = true
		}
	}
	return true
}
