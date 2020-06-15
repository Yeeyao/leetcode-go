package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("13 调整数组顺序使奇数位于偶数前面", func(t *testing.T) {
		n := 3
		want := 4
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，
并保证奇数和奇数，偶数和偶数之间的相对位置不变。
如果不要求相对位置不变，那直接从两边向中间遍历，然后左边的偶数和右边的奇数交换就行
如果对相对位置有要求 遍历两次，分别将奇数和偶数的元素保存到数组中
*/
func solution(nums []int) []int {
	numsLen := len(nums)
	res := make([]int, numsLen)
	for _, n := range nums {
		if n%2 != 0 {
			res = append(res, n)
		}
	}
	for _, n := range nums {
		if n%2 == 0 {
			res = append(res, n)
		}
	}
	return res
}

/*
	如果不使用额外的空间 遍历数组，然后统计偶数元素的数量
 */
func solution(nums []int) {

}
