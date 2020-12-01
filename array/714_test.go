package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("714. Best Time to Buy and Sell Stock with Transaction Fee", func(t *testing.T) {
		prices := []int{1, 3, 2, 8, 4, 9}
		fee := 2
		want := 8
		got := solution(prices, fee)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("714. Best Time to Buy and Sell Stock with Transaction Fee2", func(t *testing.T) {
		prices := []int{1, 3, 2, 8, 4, 9}
		fee := 2
		want := 8
		got := solution2(prices, fee)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里多了一个交易的费用，因此，每次比较的时候，需要减去交易费用
	之前是直接用后面的减去前面的
	不对，因为需要交易手续费，所以，还需要考虑交易次数，越少越好
	一般来说使用 DP 则有 dp[i] = dp[i-1] + ...
	这里有另一种保存 DP 状态的方式
	2 states:
		hold 我们在时间 i - 1 上获得最大收益时购买了股票
		empty 我们在时间 i - 1 上获得最大收益时没有购买股票
	初始化：
		对于 hold 状态，我们的收益是 -prices[0]
		对于 empty 状态，我们的收益是 0

	对于时间 i 最大的收益是我们在 i - 1 时是 hold 的状态
	或者我们重新购买股票此时收益是 empty - prices[i]

	empty 状态下，最大的收益是我们仍然是 empty 或者我们销售我们
	hold 状态下的股票 获得 hold + prices[i] - fee

	最后返回的一定是 empty 因为 empty 才不会持有股票而亏损

*/
func solution(prices []int, fee int) int {
	hold, empty := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		if hold < empty-prices[i] {
			hold = empty - prices[i]
		}
		if empty < hold+prices[i]-fee {
			empty = hold + prices[i] - fee
		}
	}
	return empty
}

/*
	k = +INF fee 只需要在购买或者出售的时候减去 fee
	T[i][k][1] = T[i-1][k-1][1]
	T[i][k][0] = T[i-1][k-1][0]

	T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	T[i][k][1] = max(T[i-1][k][1], T[i-1][k][0] - prices[i] - fee)
	同样不需要保存上一个 T[i-1][k][0] 所以更新 tik1 的时候直接使用更新完的 tik0
*/
func solution2(prices []int, fee int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	tik0, tik1 := 0, intMin
	for _, p := range prices {
		if tik0 < tik1+p {
			tik0 = tik1 + p
		}
		if tik1 < tik0-p-fee {
			tik1 = tik0 - p - fee
		}
	}
	return tik0
}
