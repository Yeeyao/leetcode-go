package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("123. Best Time to Buy and Sell Stock III", func(t *testing.T) {
		input := []int{1, 2, 3, 0, 2}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("123. Best Time to Buy and Sell Stock III2", func(t *testing.T) {
		input := []int{10, 20, 30, 40, 50, 60}
		want := 50
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	with cooldown 出售时不受影响，购买时有间隔 保存 T[i-2]

	T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	T[i][k][1] = max(T[i-1][k][1], T[i-2][k][0] - prices[i])
	注意这里保存的 T[i - 2] 这里第二次的 tik0old 还是 T[0][k][0]
	tik0old 0 0 1 2
	tik0p   0 0 0 1
	tik0p   0 0 1 2
*/
func solution(prices []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	tik0, tik1, tik0p := 0, intMin, 0
	for _, p := range prices {
		tik0old := tik0
		if tik0 < tik1+p {
			tik0 = tik1 + p
		}
		if tik1 < tik0p-p {
			tik1 = tik0p - p
		}
		tik0p = tik0old
	}
	return tik0
}
