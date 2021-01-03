package string

import (
	"fmt"
	"testing"
)

// 思路就是，一次遍历，记录下当前两个字母分别出现的次数，当出现次数相同时，将总次数+1并重新计数
func TestPro(t *testing.T) {
	t.Run("leetcode 1221 Split a String in Balanced Strings", func(t *testing.T) {
		input := "RLRRLLRLRL"
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(s string) int {
	var rnum, jnum, sum = 0, 0, 0
	for _, v := range s {
		if rnum == jnum && rnum != 0 {
			sum++
			rnum, jnum = 0, 0
		}

		if v == 'R' {
			rnum++
		} else {
			jnum++
		}
	}

	// 循环结束最后
	if rnum == jnum && rnum != 0 {
		sum++
	}
	return sum
}
