package jianzhi_offer_copy

/*
309. Best Time to Buy and Sell Stock with Cooldown
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete as many transactions as you like
(ie, buy one and sell one share of the stock multiple times) with the following restrictions:
    You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
    After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

给定股价数组，计算最大收益，注意卖了一次之后，需要等待一天才能再次购买
使用动态规划
*/

/*
	将买入和卖出分开考虑，买入为负收益，卖出为正收益
	开始只能买入，卖出获得正收益，尽可能降低负收益而提高正收益
	使用动态规划维护股市每一天结束后可以获得的累计最大收益，并进行状态转移
	f[i] 表示第 i 天结束后的累计最大收益
		f[i][0] 表示目前持有一支股票对应的累计最大收益
		f[i][1] 表示不持有任何股票且在冷冻期，这里的意思是这天卖了然后不持有股票
		f[i][2] 表示不持有任何股票且不在冷冻期，即之前卖了
	第 i 天进行操作依赖于第 i - 1 天的情况
		f[i][0] 可能 i - 1 天持有了股票到 i 也可能 i - 1 天在非冷冻期然后现在购买了
		f[i][0] = max(f[i-1][0], f[i][2] - prices[i])
		f[i][1] = f[i][0] + prices[i]
		f[i][2] = max(f[i-1][2], f[i-1][1])
	最终结果就是 max(f[n-1][1], f[n-1][2]) 毕竟最后一天持有股票是亏
	初始值 f[0][0] = -prices[0] f[0][1] = 0 f[0][2] = 0

如果给定数组长度为 0 直接返回 0
初始化 dp 数组以及初始值，然后转移方程
最后返回
*/
func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen == 0 {
		return 0
	}
	f := make([][3]int, pricesLen)
	f[0][0], f[0][1], f[0][2] = -prices[0], 0, 0
	for i := 1; i < pricesLen; i++ {
		f[i][0] = max(f[i-1][2]-prices[i], f[i-1][0])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][2], f[i-1][1])
	}
	return max(f[pricesLen-1][1], f[pricesLen-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
