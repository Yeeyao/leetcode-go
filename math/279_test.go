package math

import (
	"math"
	"testing"
)

/*
[ref](https://leetcode-cn.com/problems/perfect-squares/solution/wan-quan-ping-fang-shu-by-leetcode/)
279. Perfect Squares
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...)
which sum to n.
给定正整数 n，找到合适的最少数量的平方数
使用回溯法 找到小于 n 的最大平方数
*/
func TestPro(t *testing.T) {
	t.Run("279. Perfect Squares", func(t *testing.T) {
		n := 12
		want := 3
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares2", func(t *testing.T) {
		n := 13
		want := 2
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares3", func(t *testing.T) {
		n := 1
		want := 1
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares4", func(t *testing.T) {
		n := 0
		want := 0
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares5", func(t *testing.T) {
		n := 9
		want := 1
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("279. Perfect Squares6", func(t *testing.T) {
		n := 2
		want := 2
		got := numSquares(n)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	贪心枚举
	初始化 factor map

	isDividedBy 函数
		count == 返回是否存在
		遍历 map 然后递归调用 isDividedBy

	主循环里面初始化 count = 1 递增到 n 然后对每个 count 进行 isDividedBy 判断
	如果返回 true 就表示这个 count 满足要求，因为是从 1 开始，所以一定是最小的先找到
	这里只是简单判断能否通过 count 来满足条件
*/

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	maxLess := 1
	var factor []int
	for ; maxLess*maxLess <= n; maxLess++ {
		factor = append(factor, maxLess*maxLess)
	}
	maxLess--
	if n == 1 || maxLess*maxLess == n {
		return 1
	}
	count := 1
	for ; count <= n; count++ {
		if isDividedBy(n, count, factor) {
			return count
		}
	}
	return count
}

/*
	判断可否划分
*/
func isDividedBy(n, count int, factor []int) bool {
	// 如果 count == 1 就判断是否只含有一个刚好等于 n
	if count == 1 {
		return isContain(n, factor)
	}
	// 对 factor 的每一个都要判断
	for _, i := range factor {
		if isDividedBy(n-i, count-1, factor) {
			return true
		}
	}
	return false
}

// 因为这里已经是有序的，所以直接二分查找
func isContain(n int, factor []int) bool {
	left, right := 0, len(factor)-1
	for left < right {
		mid := left + (right-left)/2
		if factor[mid] < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return factor[left] == n
}

/*
	数学方法 Bachet 猜想 它指出每个自然数都可以表示为四个整数平方和
*/
func numSquares2(n int) int {
	if isSquare(n) {
		return 1
	}
	// 当 n 可以被写成 4^k*(8*m + 7) 结果是 4
	for n&3 == 0 {
		n >>= 2
	}

	if n&7 == 7 {
		return 4
	}

	// 分解
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isSquare(n int) bool {
	sqrtN := int(math.Sqrt(float64(n)))
	return sqrtN*sqrtN == n
}

/*
	numSquares(n) = numSquares(n-k) + 1 k 属于 1, 4, 9, ...
	先构建小于等于 n 的平方数组 数组最大数值等于 n 或者 n = 1 就返回 1
	dp[0] 初始化为 0
	循环 i 从 1 到 n
		初始化 min = n
		循环 j 从 0 开始表示 平方数组的元素索引，从小于等于 i 所有满足 i - k
		的 dp 数中找到最小值，然后将它 + 1 作为当前的 dp 值
	最后返回 dp[n]
*/
func numSquares3(n int) int {
	if n == 0 {
		return 0
	}
	maxLess := 1
	var factor []int
	for ; maxLess*maxLess <= n; maxLess++ {
		factor = append(factor, maxLess*maxLess)
	}
	if n == 1 || maxLess*maxLess == n {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	factorLen := len(factor)
	// 对每个 dp 需要判断之前的元素 dp[i-1] dp[i-4] dp[i-9] ...
	for i := 1; i <= n; i++ {
		min := n
		for j := 0; j < factorLen && factor[j] <= i; j++ {
			if dp[i-factor[j]] < min {
				min = dp[i-factor[j]]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}

/*
	TLE
*/
func numSquares10(n int) int {
	if n == 0 {
		return 0
	}
	maxLess := 1
	var factor []int
	for ; maxLess*maxLess <= n; maxLess++ {
		factor = append(factor, maxLess*maxLess)
	}
	if n == 1 || maxLess*maxLess == n {
		return 1
	}
	res := n
	helper(factor, len(factor)-1, n, 0, &res)
	return res
}

func helper(factor []int, start, n, temp int, count *int) {
	if n == 0 {
		if temp < *count {
			*count = temp
		}
		return
	}
	for i := start; i >= 0; i-- {
		if n-factor[i] >= 0 {
			n -= factor[i]
			temp++
			helper(factor, i, n, temp, count)
			n += factor[i]
			temp--
		} else {
			continue
		}
	}
}
