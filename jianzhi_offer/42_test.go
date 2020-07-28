package jianzhi_offer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("42 和为s的两个数字", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		get := solution(nums, target)
		want := []int{2, 7}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("42 和为s的两个数字2", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		get := solution2(nums, target)
		want := []int{2, 7}
		if !reflect.DeepEqual(get, want) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	类似 2-sum 但是有一个就返回
	使用一个 map 然后找
*/
func solution(nums []int, target int) []int {
	intMap := make(map[int]bool)
	// 每个元素把要求的差值保存进去，如果遍历到了差值的元素就直接返回
	for _, n := range nums {
		if _, ok := intMap[n]; ok {
			return []int{target - n, n}
		}
		intMap[target-n] = true
	}
	fmt.Println(intMap)
	return []int{}
}

/*
	注意这里是已排序数组，使用双指针
	思路是
*/
func solution2(nums []int, target int) []int {
	numsLen := len(nums)
	i, j := 0, numsLen-1
	for i < j {
		left := target - nums[i]
		if left == nums[j] {
			return []int{nums[i], nums[j]}
		} else if left > nums[j] {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
