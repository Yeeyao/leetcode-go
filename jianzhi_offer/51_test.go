package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("51 构建乘积数组", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		get := solution(nums)
		want := []int{120, 60, 40, 30, 24}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。
	不能使用除法。
	就是当前的 B[i] 是除了 A[i] 所有其他数乘积
	构建两个数组保存左右两边累计乘积 两个数组数值都是除了它当前位其他的左边或者右边乘积
	leetcode 238
*/
func solution(a []int) []int {
	aLen := len(a)
	// 1 1 2 6 24
	mul1 := make([]int, aLen)
	// 120 60 20 5 1
	mul2 := make([]int, aLen)
	res := make([]int, aLen)
	for i := 0; i < aLen; i++ {
		mul1[i], mul2[i] = 1, 1
	}
	for i := 1; i < aLen; i++ {
		mul1[i] = mul1[i-1] * a[i-1]
		mul2[aLen-i-1] = mul2[aLen-i] * a[aLen-i]
	}
	for i := 0; i < aLen; i++ {
		res[i] = mul1[i] * mul2[i]
	}
	return res
}
