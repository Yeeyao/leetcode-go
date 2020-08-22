package math

import "testing"

func TestPro(t *testing.T) {
	t.Run("136. Single Number", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3}
		want := 3
		got := solution(nums)
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

给定非空整型数组，除了一个元素，其他元素都出现了两次，找到给元素

直接将所有的元素都进行异或，两个相同的数异或得到 0
0 和任何数异或得到数本身
*/

func solution(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
