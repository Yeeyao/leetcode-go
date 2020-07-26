package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("37  数组中数字出现的次数 ", func(t *testing.T) {
		nums := []int{4, 1, 4, 6}
		get := solution(nums)
		want := []int{1, 6}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	找出数组中只出现一次的元素，要求时间复杂度 O(n) 空间复杂度 O(1)
	除了两个数字外其他都出现了 2 次
	用 map 保存然后筛选是不行的
分组异或
	异或的结果，两个相同的数异或得 0 不同的数异或得 1
	需要将数组分为两组进行异或就知道哪两个数字不同
		重复的数字需要分到同一组，利用奇偶分组
		不同数字的分组

	a, b 是两个不同的数字
`	所有的数字进行异或的结果就是 a 和 b 异或的结果 x
	x 二进制形式 xkxk-1...x2x1x0

算法
	对所有数字进行异或得到 a b 的异或值
	在异或的结果中找到任意为 1 的位
	根据这位对所有数字进行分组
	每个组内进行异或，得到两个数字
*/
func solution(nums []int) []int {
	ret := 0
	// a b 异或结果
	for _, n := range nums {
		ret ^= n
	}
	// 找到第一个异或结果是 1 的位，这里其实只要找到一个是 1 的位就可以
	div := 1
	for div&ret == 0 {
		div <<= 1
	}
	// 分组异或 注意这里 div 和有相同两个数字的进行异或一定会分到同一组，
	// 因为相同的两个数字，和 div 1 对应的位数也是相同的
	// 同时，跟 a 或者 b 异或将它们分到不同的分组
	// 因为 a b 异或的结果产生的 1，原来 a 和 b 的对应位数一定是 0 和 1，所以一定分到不同的组
	a, b := 0, 0
	for _, n := range nums {
		if div&n == 1 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}
