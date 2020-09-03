package math

/*
[ref](https://leetcode-cn.com/problems/target-sum/solution/mu-biao-he-by-leetcode/)
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S.
Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.
Find out how many ways to assign symbols to make sum of integers equal to target S.
给定非负整型列表，只能使用 + 和 - ，给定 target 判断有多少种 + - 组合应用于所有的整型可以得到 s
这里给的是非负整型
*/

/*
	转换为背包问题
	把所有符号为正的数总和设为一个背包的容量 x，把所有符号为负的数总和设为一个背包容量 y
	在给定数组中，有多少种选择方法让背包装满，令 sum 为数组总和，则 x + y = sum
	两个背包的差 S，则 x - y = S，从而求得 x = (S + sum) / 2

	问题转换为背包问题，给定一个数组和一个容量为 x 的背包，求有多少种方式让背包装满

	先求所有元素的总和，然后判断背包数的整数以及目标和是否满足要求
		这里的意思是转换成了求和为 x 的背包问题，即给定数组有多少中方法将 x 装满
		这里 i 表示容量为 i 时有多少种装法
		dp[i] = dp[i-a] + dp[i-b] + dp[i-c] +... a, b, c 属于装进背包的东西

使用其中一个 x 即 (sum + S) / 2作为背包总空间
	第一个循环递归遍历 nums 每个元素
	第二个循环 i 从 x 的最大值开始，条件是 i >= n 即当前 nums 的元素值，每次递减
		dp[i] += dp[i-n]
	最后返回 dp[x]

*/
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	// 背包量需要是整数
	if (sum+S)%2 == 1 {
		return 0
	}
	// 目标和不能大于总数
	if S > sum {
		return 0
	}
	// 真正背包问题 这里算 x 的数量
	dpLen := (sum + S) / 2
	dp := make([]int, dpLen+1)
	dp[0] = 1
	for _, n := range nums {
		for i := dpLen; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[dpLen]
}

/*
dp
dp[i][j] 表示用数组中前 i 个元素，组成和为 j 的方案数
dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]]
也可以写成递推式
dp[i][j+nums[i] += d[i-1][j]
dp[i][j-nums[i] += d[i-1][j]

因为数组中所有数的和不超过 1000，则 j 最小值可以达到 -1000
我们需要给 dp[i][j] 的第二维先增加 1000
dp[i][j + nums[i] + 1000] += dp[i-1][j + 1000]
dp[i][j + nums[i] + 1000] += dp[i-1][j + 1000]
*/
func findTargetSumWays2(nums []int, S int) int {
	dp := make([][2001]int, len(nums))
	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000] = 1
	for i := 1; i < len(nums); i++ {
		for sum := -1000; sum <= 1000; sum++ {
			if dp[i-1][sum+1000] > 0 {
				dp[i][sum+nums[i]+1000] += dp[i-1][sum+1000]
				dp[i][sum-nums[i]+1000] += dp[i-1][sum+1000]
			}
		}
	}
	if S > 1000 {
		return 0
	}
	return dp[len(nums)-1][S+1000]
}

// 直接暴力枚举
func findTargetSumWays3(nums []int, S int) int {
	res := 0
	helper(nums, 0, 0, S, &res)
	return res
}

func helper2(nums []int, start, temp, S int, res *int) {
	// 处理完了
	if start == len(nums) {
		if temp == S {
			*res++
		}
	} else {
		helper(nums, start+1, temp+nums[start], S, res)
		helper(nums, start+1, temp-nums[start], S, res)
	}

}
