package hash_table

import "testing"

func TestPro(t *testing.T) {
	t.Run("leetcode 771 Jewels and Stones", func(t *testing.T) {
		inputone := "aA"
		inputtwo := "aAAbbbb"
		want := 3
		got := solution(inputone, inputtwo)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(j, s string) int {
	var sum = 0
	for _, v := range s {
		if charinj(j, v) {
			sum += 1
		}
	}
	return sum
}

func charinj(j string, c rune) bool {
	for _, v := range j {
		if v == c {
			return true
		}
	}
	return false
}
