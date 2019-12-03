package string

import "testing"

func TestPro(t *testing.T) {
	t.Run("leetcode 709 To Lower Case", func(t *testing.T) {
		input := "RLRRLLRLRL"
		want := "rlrrllrlrl"
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(s string) string {
	var rets string
	const diff = 'a' - 'A'

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			rets = rets + string(v + diff)
		} else {
			rets = rets + string(v)
		}
	}
	return rets
}
