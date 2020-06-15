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
 */
func solution(num int) int {

}
