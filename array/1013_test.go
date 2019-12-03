package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("1013. Partition Array Into Three Parts With Equal Sum", func(t *testing.T) {
		input := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1013. Partition Array Into Three Parts With Equal Sum2", func(t *testing.T) {
		input := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1013. Partition Array Into Three Parts With Equal Sum3", func(t *testing.T) {
		input := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	从左边以及右边向中间遍历并累加
	终止条件是两个指针相交
	如果其中一个的和等于总和的三分之一，则只需要另外一个移动并判断
*/
//func solution(input []int) bool {
//	sum := calcSum(input)
//	if sum%3 != 0 {
//		return false
//	}
//	inputLen := len(input)
//	partSum := sum / 3
//	leftSum, rightSum := 0, 0
//	leftIndex, rightIndex := 0, inputLen-1
//	// 边界条件需要考虑
//	for leftIndex < rightIndex {
//		// 左边可以停止遍历了
//		if leftSum == partSum {
//			// leftIndex 会多一个
//			for rightIndex >= leftIndex {
//				if rightSum == partSum {
//					// 两边都满足，直接返回
//					return true
//				} else {
//					// 右边还需要遍历
//					rightSum += input[rightIndex]
//					rightIndex--
//				}
//			}
//			return false
//		}
//		if rightSum == partSum {
//			for leftIndex <= rightIndex {
//				if leftSum == partSum {
//					// 两边都满足，直接返回
//					return true
//				} else {
//					// 右边还需要遍历
//					leftSum += input[leftIndex]
//					leftIndex++
//				}
//			}
//			return false
//		}
//		leftSum += input[leftIndex]
//		leftIndex++
//		rightSum += input[rightIndex]
//		rightIndex--
//	}
//	return false
//}

func solution(input []int) bool {
	sum := calcSum(input)
	if sum%3 != 0 {
		return false
	}
	partSum := sum / 3
	count := 0
	tempSum := 0
	for _, v := range input {
		tempSum += v
		if tempSum == partSum {
			count++
			tempSum = 0
		}
	}
	return count == 3
}

func calcSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func canThreePartsEqualSum(nums []int) bool {
	// 子数组和方法
	total := sum(nums)
	if total%3 != 0 {
		return false
	}
	target := total / 3
	subtotal := 0
	count := 0
	for i := range nums {
		subtotal += nums[i]
		if subtotal == target {
			subtotal = 0
			count++
		}
	}
	return count == 3
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
