package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("121. Best Time to Buy and Sell Stock", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		want := 5
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("121. Best Time to Buy and Sell Stock2", func(t *testing.T) {
		input := []int{7, 6, 4, 3, 1}
		want := 0
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("121. Best Time to Buy and Sell Stock3", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		want := 5
		got := solution2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("121. Best Time to Buy and Sell Stock4", func(t *testing.T) {
		input := []int{7, 6, 4, 3, 1}
		want := 0
		got := solution2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	可以怎么知道当前是最低价位？
	初始当前最大收益是 0 初始化第一个元素是最小价格，遍历每个价格
	如果当前价格-最小价格大于最大收益就更新最大收益，然后如果当前元素小于最小价格就更新最小价格
*/
func solution(input []int) int {
	if len(input) == 0 {
		return 0
	}
	maxBenefit := 0
	minPrice := input[0]
	for _, v := range input {
		benefit := v - minPrice
		if benefit > maxBenefit {
			maxBenefit = benefit
		}
		if v < minPrice {
			minPrice = v
		}
	}
	return maxBenefit
}

/*
	每天只有两个变量 T[i][1][0] 以及 T[i][1][1]
	T[i][1][0] = max(T[i-1][1][0], T[i-1][1][1] + prices[i])
	T[i][1][1] = max(T[i-1][1][1], T[i-1][0][0] - prices[i])
	这里 T[i-1][0][0] 是 0
*/
func solution2(prices []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	ti10, ti11 := 0, intMin
	for _, p := range prices {
		if ti10 < ti11+p {
			ti10 = ti11 + p
		}
		if ti11 < -p {
			ti11 = -p
		}
	}
	return ti10
}
