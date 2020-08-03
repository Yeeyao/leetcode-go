package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("13 调整数组顺序使奇位于偶数前面", func(t *testing.T) {
		num := []int{1, 2, 3, 4}
		want := []int{1, 3, 2, 4}
		got := solution(num)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("13 调整数组顺序使奇位于偶数前面2", func(t *testing.T) {
		num := []int{1, 2, 3, 4}
		want := []int{1, 3, 2, 4}
		got := solution2(num)
		if !reflect.DeepEqual(got, want) {
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
	var res []int
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
func solution2(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if odd(nums[i]) {
			if odd(nums[j]) {
				j--
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		} else {
			if odd(nums[j]) {
				i++
				j--
			} else {
				i++
			}
		}
	}
	return nums
}

func odd(i int) bool {
	return i%2 == 0
}
