package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("7 旋斐波那契数列", func(t *testing.T) {
		n := 0
		want := 0
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("7 旋斐波那契数列2", func(t *testing.T) {
		n := 1
		want := 1
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("7 旋斐波那契数列3", func(t *testing.T) {
		n := 5
		want := 5
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。 n<=39
*/
func solution(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a + b
	}
	return a
}
