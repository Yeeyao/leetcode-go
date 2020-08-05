package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("30   连续子数组的最大和", func(t *testing.T) {
		intSlice := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		get := solution2(intSlice)
		want := 6
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("30   连续子数组的最大和2", func(t *testing.T) {
		intSlice := []int{-2, -11, -3, -4, -1, -2, -1, -5, -14}
		get := solution2(intSlice)
		want := -1
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
	要求时间复杂度为O(n)
	开始的总和设置为最小的负数
	需要先加上当前的数值才判断，如果当前的总和小于 0 ， 下次的总和需要设置为 0
*/
func solution(nums []int) int {
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	maxSum := intMin
	tempSum := 0
	for _, v := range nums {
		tempSum += v
		if tempSum > maxSum {
			maxSum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
		}
	}
	return maxSum
}

func solution2(nums []int) int {
	maxSum := nums[0]
	tempSum := nums[0]
	numsLen := len(nums)
	for i := 1; i < numsLen; i++ {
		tempSum += nums[i]
		if tempSum > maxSum {
			maxSum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
		}
	}
	return maxSum
}
