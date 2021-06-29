package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("374. Guess Number Higher or Lower II", func(t *testing.T) {
		n := 2
		got := solution(n)
		want := 2
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	猜测的游戏
	选择 1 到 n 的一个数字
	如果猜对了，赢得该游戏
	如果猜错了，将告诉你猜得太高还是太低了，然后继续猜测
	每次猜错了的数字是 x，你将会花费 x dollars，如果用完了钱将输掉游戏
	给定 n，然后无论选择什么数字需要保证胜利的最少金钱数量，这种其实应该需要想到是 DP
	但是感觉题目有歧义

	minMax 问题
	这里的意思是对每个数字有猜的数字是 k 以及不是 k 的情况，然后从 a 到 b 每个数字都要判断

	DP[a,b] 表示 [a, b] 区间内，至少需要的钱，k 表示本次需要猜测的数字，这里需要三个循环，第一个是枚举左端点 l 第二个是枚举右端点 r 第三个是枚举 k
	DP[a,b] = min(max(DP[a, k-1], DP[k+1, b] + k))

	DP[i] 定义为当前保证胜利所需要的最少金钱
	关系：DP[i] = max(DP[i], DP[i-1])

	[ref](https://blog.csdn.net/Site1997/article/details/100168676)
	[ref](https://artofproblemsolving.com/community/c296841h1273742)

	[ref](https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/solution/cai-shu-zi-da-xiao-ii-by-leetcode/)

	我们需要意识到我们在范围 (1,n) 中猜数字的时候，需要考虑最坏情况下的代价。也就是说要算每次都猜错的情况下的总体最大开销。
	先是暴力解法 从 (1,n) 任意选择一个数字，假设选择错误(最坏的情况)，需要最小代价猜需要的数字，则一次尝试后，答案在我们猜测的左边或者右边，需要比较两边的较大值
		cost(1, n) = i + max(cost(1, i-1), cost(i+1, n))
		这里就需要从 1 到 n 遍历下去，对每个数字都需要计算 cost 时间复杂度 O(n!)
	暴力优化
		暴力中，对 (i, j) 中每个数字，都需要从当前数值然后考虑左右两个区间的代价，但是一个重要发现是，如果从 i, (i+j)/2 中选择数字
		作为第一次尝试，右边区间比左边的大，所以只需要从右边区间获取最大开销就行，因为肯定比左边的开销大，因此第一次尝试从 (i+j)/2, j 中选取
		所以就不需要从 i 到 j 遍历每个数字，只需要从 (i+j)/2, j 遍历 时间复杂度还是 O(n!)
	DP TODO: 需要学习这部分的分析方法才能理解解法
		这里以 i 为第一次尝试找到最小开销的过程可以被分解为找左右区间内最小开销的子问题，对每个区间，重复问题拆分过程，得到更多子问题，因此想到 DP
		使用一个 dp 矩阵，dp(i, j) 表示 (i, j) 在最坏情况下最小开销代价
		考虑如何求出 dp 数组，如果区间剩下一个数字 k，则猜中代价永远是 0，对于长度是 2 的区间，需要所有长度是 1 的区间的结果。因此，为了求出长度为
		len 的区间的解，需要求出所有长度为 len - 1 的解，因此按照区间的长度由短到长来求出 dp 数组

		对每个 dp(i, j) 长度 len = j - i + 1 按照暴力饿方法，因此选择每个数字作为第一次尝试的答案，可以求出最小开销
			cost(i, j) = pivot + max(cost(i, pivot - 1), cost(pivot + 1, j))
		因为我们已经知道了小于 len 长度的 dp 数组的所有答案，因此 dp 方程可以变成
			dp(i, j) = minPivots(i, j)(pivot + max(dp(i, pivot - 1), dp(pivot + 1, j) minPivots(i, j) 表示将 (i, j) 中每个数字作为第一个尝试的数
		这里需要从小区间向大区间计算
			这里观察后面两个，第一个是 i 不动，然后 j 根据 pivot 变化，第二个是 j 不动，i 根据 pivot 变化，就前面说的下界需要枚举，
			上界也需要枚举，因此这里会有两个循环，最外层又有一个 pivot 的变化循环。因此是三个循环
*/

/*
	根据暴力方法的优化，
			dp(i + (len - 1) / 2, j) = minPivots(i, j)(pivot + max(dp(i, pivot - 1), dp(pivot + 1, j) minPivots(i, j) 表示将 (i, j) 中每个数字作为第一个尝试的数
*/
func solution(n int) int {
	const intMax = int(^uint(0) >> 1)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// length 为 2 时， start 从 1 到 n-1
	for length := 2; length <= n; length++ {
		for start := 1; start <= n+1-length; start++ {
			minRes := intMax
			for piv := start + (length-1)/2; piv < start+length-1; piv++ {
				res := piv + max(dp[start][piv-1], dp[piv+1][start+length-1])
				minRes = min(res, minRes)
			}
			dp[start][start+length-1] = minRes
		}
	}
	return dp[1][n]
}

/*
	这里主要是需要怎么编写循环。。。
*/
func solution2(n int) int {
	const intMax = int(^uint(0) >> 1)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// length 为 2 时， start 从 1 到 n-1
	for length := 2; length < n; length++ {
		for start := 1; start <= n+1-length; start++ {
			minRes := intMax
			for piv := start; piv < start+length-1; piv++ {
				res := piv + max(dp[start][piv-1], dp[piv+1][start+length-1])
				minRes = min(res, minRes)
			}
			dp[start][start+length-1] = minRes
		}
	}
	return dp[1][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
