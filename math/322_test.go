package math

import "math"

/*
322. Coin Change
You are given coins of different denominations and a total amount of money amount.
Write a function to compute the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.
给定不同面额的硬币以及一个金额，编写一个函数返回用最少硬币组成面额的数量，无法组成返回 -1

这不就是背包问题吗
[ref](https://leetcode-cn.com/problems/coin-change/solution/322-ling-qian-dui-huan-by-leetcode-solution/)
*/
/*
dp 自下而上
F[i] 是组成金额 i 所需的最少硬币数量，计算 F[i] 前已经计算了 F[0]...F[i-1] 的结果
F[i] = min(F[i-cj] + 1, F[i])

初始化 dp 数组，长度是 amount + 1 初始值是 count + 1 dp[0] = 0
第一个循环从 1 到 amount 包含遍历
	第二个循环遍历 coins 每个元素，如果 coins[j] <= i 表示当前的金额可以放下当前的硬币
	判断 dp[i-coins[j]] + 1 和 dp[i] 取较小值 即这里找到每个金额所需最小硬币数量

*/
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i, _ := range dp {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 计算 F[i] 中 F[i-Cj] 的最小值
		for j := 0; j < len(coins); j++ {
			// 面额不超过 i 的才需要计算
			if coins[j] <= i {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

/*
DP

自上而下
F(S) 组成金额 S 所需的最少硬币数量 [c0...cn-1] 可选硬币面额
假如知道最后一枚硬币面值为 C，则 F(S) = F(S-C) + 1
然而我们并不知道最后一枚硬币的面值，所以需要枚举然后取最小值
F(S) = min(F(S-Ci) + 1 F(0) = 0
O(Sn) 时间，O(S) 空间

初始化 count 数组保存金额所需最少硬币数量
递归调用函数 参数是 硬币面额，dp 数组以及金额
	如果剩余金额小于 0 直接返回 -1 表示当前的面额太大了
	剩余金额等于 0 直接返回 0，因为金额为 0 不需要硬币了
	如果 dp 数组中的数值非 0 表示之前已经计算了，直接返回该数值
	对每个面值硬币调用递归函数得到结果来更新 min，结果大于等于 0 小于当前的 min 值就更新 min 值 最后更新 count 数组
*/
func coinChange2(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	count := make([]int, amount)
	return Change(coins, count, amount)
}

func Change(coins, count []int, rem int) int {
	// 剩下的金额
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	// 这里避免重复计算
	if count[rem-1] != 0 {
		return count[rem-1]
	}
	min := math.MaxInt64
	for _, c := range coins {
		res := Change(coins, count, rem-c)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == math.MinInt64 {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}
