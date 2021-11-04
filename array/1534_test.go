package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1534. Count Good Triplets", func(t *testing.T) {
		nums := []int{35, 15, 38, 1, 10, 26}
		want := 1
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1995. Count Special Quadruplets2", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 5}
		want := 4
		got := solution(nums)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定一个整型数组 arr 以及三个整型 a, b, c 需要找到 good triplets 的数量
	good triplet 定义 (arr[i], arr[j], arr[k]) 满足
	- 0 <= i < j < k < arr.length
	- | arr[i] - arr[j] | <= a
	- | arr[j] - arr[k] | <= b
	- | arr[i] - arr[k] | <= c

	绝对值判断

	这里暴力的优化，主要是利用计算得到的绝对值吧

	O(n^2) 枚举满足 |arr[j]-arr[k]| <= b 的二元组 (j,k)，统计该二元组下有多少 i 满足条件。i 的条件是 | arr[i] - arr[j] | <= a，
	| arr[i] - arr[k] | <= c；去除绝对值，得到 a[i] 数值需要在 [a[j] - a, a[j] + a]，[a[k] - c, a[k] + c]，的两个区间的交集，
	记作 [l,r] 因此，枚举 j，k 二元组的时候，就需要统计出 i < j 且 arr[i] 的数值在 [l,r] 范围内的数量

	维护一个 arr[i] 频次数组的前缀和 sum，对一个二元组 (j,k) 我们可以 O(1) 得到答案为 sum[r]-sum[l-1]。如何维护频次数组存的数下标符合
	i < j 的限制，只需要从小到大枚举 j 每次 j 移动指针 + 1 时将 arr[j] 的值更新到 sum 这样保证枚举到 j 的时候，sum 数组存的数值下标满足限制

*/
func solution(arr []int, a, b, c int) int {
	count := 0
	arrLen := len(arr)
	// 统计元素数量
	sumArr := make([]int, 1001)
	// 每一对 j,k 进行统计
	for j := 0; j < arrLen; j++ {
		for k := j + 1; k < arrLen; k++ {
			if abs(arr[j]-arr[k]) <= b {
				// 计算然后取交集
				lj, rj := arr[j]-a, arr[j]+a
				lk, rk := arr[k]-c, arr[k]+c
				l, r := max(0, max(lj, lk)), min(1000, min(rj, rk))
				if l <= r {
					if l == 0 {
						count += sumArr[r]
					} else {
						count += sumArr[r] - sumArr[l-1]
					}
				}
			}
		}
		for k := arr[j]; k <= 1000; k++ {
			sumArr[k]++
		}
	}
	return count
}

/*
	暴力方法就是遍历所有元素，然后计算。O(n^3)
	直接枚举所有的元素
*/
func solution2(arr []int, a, b, c int) int {
	count := 0
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			for k := j + 1; k < arrLen; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
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
