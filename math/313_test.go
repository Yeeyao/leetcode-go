package math

import "testing"

func TestPro(t *testing.T) {
	t.Run("313. Super Ugly Number", func(t *testing.T) {
		num := 9
		want := true
		got := solution(num)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
313 Super Ugly Number
编写函数找到第 n 个 super ugly number
定义：所有的素数因子存在于给定的素数列表中的正数
就是需要先找到所有素数因子存在于给定列表中的数字，然后找到第 n 个
这题有点类似 2，3，5 那题(264)，只是这里换成了素数数组。对 264 的一般情况推广？
why 1
这里使用 dp
一般性的丑数计算是，为每个质因数建立一个指针，然后将这几个质因数运算结果中找到最小的，对比数值然后将指针 + 1
这里我们需要将计算出来的丑数再次和 primes 里面的质因数结合，计算出新的丑数，算出来的丑数放在一个 dp 数组里面
这里建立一个 index 数组，存放每个质因数下一个将要结合的 dp 下标，这个下标从 0 开始，每结合一次 + 1
[ref](https://leetcode-cn.com/problems/super-ugly-number/solution/javazui-rong-yi-li-jie-de-dong-tai-gui-hua-fen-xi-/)
*/
func solution(n int, primes []int) int {
	// dp 数组
	const MAXINT = int(^uint(0) >> 1)
	dp := make([]int, n)
	primesLen := len(primes)
	// 每个因数的指针数组
	index := make([]int, primesLen)
	dp[0] = 1
	// 从所有的因数中计算乘积，找到最小乘积，保存到 dp 数组，然后将因数的指针递增
	for i := 1; i < n; i++ {
		min := MAXINT
		for j := 0; j < primesLen; j++ {
			if min > dp[index[j]]*primes[j] {
				min = dp[index[j]] * primes[j]
			}
		}
		dp[i] = min
		for j := 0; j < primesLen; j++ {
			if min == primes[j]*dp[index[j]] {
				index[j]++
			}
		}
	}
	return dp[n-1]
}

// 100%
func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n, n)
	ugly[0] = 1
	c := 0
	candidates := make([]int, len(primes), len(primes))
	copy(candidates, primes)
	idxs := make([]int, len(primes), len(primes))
	for c < n-1 {
		m := 1<<63 - 1
		for i := range candidates {
			for candidates[i] <= ugly[c] {
				idxs[i]++
				candidates[i] = primes[i] * ugly[idxs[i]]
			}
			m = min(m, candidates[i])
		}
		c++
		ugly[c] = m
	}
	return ugly[c]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
