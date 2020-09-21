package array

import "strconv"

/*
60. Permutation Sequence
The set [1,2,3,...,n] contains a total of n! unique permutations.
n 从 1 到 9 k 从 1 到 n!
1 到 n 个数的集合其排列有 n! 个，给定 n 和 k，求 1 到 n 的集合中第 k 个排列
这里排列需要字典序 这里还是类似 46 只是需要计数器来返回，这里怎么返回
同时根据 k 的大小还能直接过滤很多没必要的组合
每个数字开始的数字串，其含有的排列数量是 (n-1)! 不知道怎么深入想下去

[ref](https://leetcode-cn.com/problems/permutation-sequence/solution/di-kge-pai-lie-by-leetcode-solution/)
每个数字开始的数字串，其含有的排列数量是 (n-1)!
k <= (n-1)! 第一个元素是 1
(n-1)! < k <= 2(n-1)! 第一个元素是 2
...
(n-1)(n-1)! < k <= n! 第一个元素为 n
所以，第 k 个排列的首个元素是 (k-1) / (n-1)! + 1 前面需要向下取整
确定第一个元素 a1 之后类似可以确定下一个元素 a2，这个 a2 开始的排列有 (n-2)! 个
除去 a1 这些的编号从 (a1-1)*(n-1)! 到 a1*(n-1)!结束总共 (n-1)! 个
第 k 个排列 k'= (k-1)mod(n-1)! + 1

建立 factorial 数组然后每个元素初始化为索引的阶乘数值 f[1] = 1! f[2] = 2!
建立 valid 数组，长度是 n + 1 元素都初始化为 1
	这里是用来进行每个开始元素下后面的元素计数
k--
循环从 1 到 n 包含
	order = k / factorial[n-i] + 1
	第二个循环从 1 到 n 包含
		这里每次将 order 减少 1，当它是 0 就表示找到了该位置的元素，保存
	然后将 k %= factorial[n-i]
*/

func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	res := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}

	k--

	for i := 1; i <= n; i++ {
		// 当前的开头元素数值
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			// 找到了当前元素的索引
			if order == 0 {
				res += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return res
}
