package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("918. Maximum Sum Circular Subarray", func(t *testing.T) {
		A := []int{3, -1, 2, -1}
		want := 4
		got := solution(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	[lee](https://leetcode.com/problems/maximum-sum-circular-subarray/discuss/178422/One-Pass)
	类比 53 这里分为两种情况，最大的子数组在中间以及在两边
	后面的情况可以转换为第一种
	所以最大的子数组和是 最大子数组和以及总和减去最小子数组的较大值

	max(the max subarray sum, the total sum - the min subarray sum)

	当所有元素是负数时， maxSum = max(A) minSum = sum(A) 此时需要返回 max(A)
*/
func solution(A []int) int {
	// 总和，最大和，当前最大和，最小和，当前最小和
	total, maxSum, curMax, minSum, curMin := 0, A[0], 0, A[0], 0
	// 注意这里从头开始遍历以及各个数值的更新顺序
	for _, a := range A {
		if a > curMax+a {
			curMax = a
		} else {
			curMax += a
		}
		if curMax > maxSum {
			maxSum = curMax
		}
		if a > a+curMin {
			curMin += a
		} else {
			curMin = a
		}
		if curMin < minSum {
			minSum = curMin
		}
		total += a
	}
	// 不是全都是负数
	if maxSum > 0 {
		if maxSum > total-minSum {
			return maxSum
		} else {
			return total - minSum
		}
	} else {
		return maxSum
	}
}
