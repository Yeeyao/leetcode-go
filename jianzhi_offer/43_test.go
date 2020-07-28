package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("43 左旋转字符串", func(t *testing.T) {
		s := "abcdefg"
		k := 2
		get := solution(s, k)
		want := "cdefgab"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	我记得是三次旋转，n 左右两边旋转，然后整体旋转
	abcd 2 ba dc cdab
*/
func solution(s string, n int) string {
	str1 := reverse([]byte(s[:n]))
	str2 := reverse([]byte(s[n:]))
	str3 := reverse([]byte(str1 + str2))
	return str3
}

/*
	旋转
*/
func reverse(s []byte) string {
	begin, end := 0, len(s)-1
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
	return string(s)
}
