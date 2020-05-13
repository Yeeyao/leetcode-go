package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("122. Best Time to Buy and Sell Stock II", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		want := 7
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("122. Best Time to Buy and Sell Stock II 2", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		want := 7
		got := solution2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input []int) int {
	if len(input) == 0 {
		return 0
	}
	maxBenefit := 0
	buy := input[0]
	for _, v := range input {
		if v > buy {
			maxBenefit += v - buy
		}
		buy = v
	}
	return maxBenefit
}

/*
	k = +INF

	T[i][k][1] = T[i-1][k-1][1]
	T[i][k][0] = T[i-1][k-1][0]

	T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	T[i][k][1] = max(T[i-1][k][1], T[i-1][k][0] - prices[i])
*/
func solution2(prices []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	tik0, tik1 := 0, intMin
	for _, p := range prices {
		// // 这里需要 T[i-1][k][0]
		// 根据观察，这里不需要这样保存旧的数值
		// tip1k0 := tik0
		if tik0 < tik1+p {
			tik0 = tik1 + p
		}
		if tik1 < tik0-p {
			tik1 = tik0 - p
		}
	}
	return tik0
}
