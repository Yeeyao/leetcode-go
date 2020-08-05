package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("leetcode 238  Product of Array Except Self", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := []int{24, 12, 8, 6}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 238  Product of Array Except Self2", func(t *testing.T) {
		input := []int{2, 3, 4, 5}
		want := []int{60, 40, 30, 24}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	思路1：
		使用辅助数组保存两边的乘积，其实就是对每个元素，两边保存好该元素两边的乘积（除了元素本身）
		第一次遍历计算两边乘积，第二次遍历直接将两边的乘积相乘就得到结果
		O(n)时间以及O(n)空间

	思路2：
		一次遍历中完成计算处理 需要注意相乘的顺序
		这里，每个结果将会有两次相乘的机会，第一次是遍历到的第一次乘以左边的积，第二次是遍历到右边去乘以右边的积
		启示是，一次遍历，两边的统计，然后一个乘积可以乘以两次
*/

func solution(input []int) []int {
	inputLen := len(input)
	retProduct := make([]int, inputLen)
	for i := 0; i < inputLen; i++ {
		retProduct[i] = 1
	}
	leftProduct := 1
	rightProduct := 1
	// 这里需要注意顺序
	for i := 0; i < inputLen; i++ {
		retProduct[i] *= leftProduct
		leftProduct *= input[i]
		retProduct[inputLen-i-1] *= rightProduct
		rightProduct *= input[inputLen-i-1]
	}
	return retProduct
}

// 需要考虑有 0 元素存在的情况 O(n) 并不使用除法，常数空间复杂度
func solution(input []int) []int {
	inputLen := len(input)
	leftProduct := make([]int, inputLen)
	rightProduct := make([]int, inputLen)
	retProduct := make([]int, inputLen)
	for i := 0; i < inputLen; i++ {
		leftProduct[i] = 1
		rightProduct[i] = 1
	}
	for i := 1; i < inputLen; i++ {
		leftProduct[i] = leftProduct[i-1] * input[i-1]
		rightProduct[inputLen-i-1] = rightProduct[inputLen-i] * input[inputLen-i]
	}
	for i := 0; i < inputLen; i++ {
		retProduct[i] = leftProduct[i] * rightProduct[i]
	}
	return retProduct
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	LP := 1
	for i, num := range nums {
		LP *= num
		res[i] = LP
	}
	res[length-1] = res[length-2]
	RP := nums[length-1]
	for j := length - 2; j >= 1; j-- {
		res[j] = res[j-1] * RP
		RP *= nums[j]
	}
	res[0] = RP

	return res
}
