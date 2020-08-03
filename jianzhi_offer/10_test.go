package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("10 矩形覆盖", func(t *testing.T) {
		n := 3
		want := 4
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	可以使用 2 * 1 的小矩形来覆盖目标的大矩形，不可重叠。其中大矩形的大小是 2 * n 求可以有多少种不同的覆盖方式
	2 * 1 覆盖 2 * n 不同的方法
	n = 1 有 1 中，n = 2 有  2 种，n = 3 有 3 种
	计 n 时有 f(n) 种， f(n) = f(n-1) + f(n-2) 类似斐波那契数列
*/
func solution(num int) int {
	a, b := 1, 1
	for i := 0; i < num; i++ {
		a, b = b, a+b
	}
	return a
}
