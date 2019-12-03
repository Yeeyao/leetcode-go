package string

import "testing"

func TestPro(t *testing.T) {
	t.Run("leetcode 1108 Defanging an IP Address", func(t *testing.T) {
		input := "1.1.1.1"
		want := "1[.]1[.]1[.]1"
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(str string) string {
	var rets string
	const add = "[.]"
	for _, v := range str {
		if v == '.' {
			rets = rets + add
		} else {
			rets = rets + string(v)
		}
	}
	return rets
}
