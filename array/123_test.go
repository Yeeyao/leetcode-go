package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("123. Best Time to Buy and Sell Stock III", func(t *testing.T) {
		input := []int{3, 3, 5, 0, 0, 3, 1, 4}
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
k = 2，因为情况较少，所以可以直接将 4 个变量都列举出来
121, 122 都没有针对 k 需要特殊处理
T[i][1][0] = max(T[i-1][1][0], T[i-1][1][1] + prices[i])
T[i][1][1] = max(T[i-1][1][1], T[i-1][0][0] - prices[i])
T[i][2][0] = max(T[i-1][2][0], T[i-1][2][1] + prices[i])
T[i][2][1] = max(T[i-1][2][1], T[i-1][1][0] - prices[i])
*/
func solution(prices []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	ti10, ti11 := 0, intMin
	ti20, ti21 := 0, intMin
	for _, p := range prices {
		// 注意一下变量的赋值顺序
		if ti20 < ti21+p {
			ti20 = ti21 + p
		}
		if ti21 < ti10-p {
			ti21 = ti10 - p
		}
		if ti10 < ti11+p {
			ti10 = ti11 + p
		}
		if ti11 < -p {
			ti11 = -p
		}
	}
	return ti20
}
