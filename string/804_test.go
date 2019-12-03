package string

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("leetcode 1021 Remove Outermost Parentheses", func(t *testing.T) {
		input := []string{"gin", "zen", "gig", "msg"}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(words []string) int {
	exist := make(map[string]bool)
	oneword := ""
	morseTable := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
		"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, v := range words {
		for _, w := range v {
			oneword += morseTable[w-'a']
		}
		if _, ok := exist[oneword]; ok {
		} else {
			exist[oneword] = true
		}
		oneword = ""
	}
	return len(exist)
}
