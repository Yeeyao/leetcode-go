package math

/*
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's
in their binary representation and return them as an array.

给定非负数 num，输出从 0 到 num 的每个数字的二进制表示中 1 的个数
这里数字数连续的
1 2 4 8 16

[ref](https://leetcode-cn.com/problems/counting-bits/solution/bi-te-wei-ji-shu-by-leetcode/)

brute force 是直接将每个数字调用辅助函数，辅助函数计算每个数字的 1 的位数

动态规划 + 最高有效位/最低有效位/最后设置位 分为三种方法
	观察 0 0，1 1，2 10，3 11
	2,3 通过在 0,1 前面加上 1 来得到
	同样对 [0,3] 进行处理可以得到 [4,7]
	状态转移方程 P(x) P(x + b) = P(x) + 1, b = 2^m > x
	利用这个从 0 开始生成所有结果

*/

/*
	使得 i 的二进制中 1 的数量减少 1
	比如 i = 14 i & (i - 1) 得到 12，然后因为是一路计算下来 res[12] 是已知的
	然后 14 对比 12 就是多个 1
	i = 17 i & (i - 1) 得到 16 同样也是得到 res[16] + 1

	res[i] = res[比 i 少一个 二进制 1] + 1
*/
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 1; i < num+1; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

/*
	所有数字只有两种，奇数和偶数
	其中奇数比前一个偶数的 1 的个数 + 1
	偶数的 1 的个数和偶数 / 2 的 1 的个数相同
*/
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	// 注意这里从 1 开始同时最后一个数字
	for i := 1; i < num+1; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}
