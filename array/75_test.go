package array

import (
	"reflect"
	"testing"
)

/*
75. Sort Colors
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent,
with the colors in the order red, white and blue.
Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
Note: You are not suppose to use the library's sort function for this problem.

给定一个含有 n 个元素的数组，元素表示 红白蓝 三种，将他们排序使得相同颜色的在一起，顺序是红白蓝，需要在位排序
使用 0， 1， 2 分别表示红白蓝三原色
跑两次的双指针 第一次将所有 0 放在前面，第二次将 1 放在前面
001122
010022 l 0 r 5 l 1 r 5 l 1 r 4 l 1 r 3 l 2 r 2
*/

func TestPro(t *testing.T) {
	t.Run("75. Sort Colors", func(t *testing.T) {
		input := []int{2, 0, 2, 1, 1, 0}
		want := []int{0, 0, 1, 1, 2, 2}
		sortColors(input)
		if !reflect.DeepEqual(want, input) {
			t.Errorf("got: %v, want: %v", input, want)
		}
	})
	t.Run("75. Sort Colors2", func(t *testing.T) {
		input := []int{0, 0, 1, 1, 2, 2}
		want := []int{0, 0, 1, 1, 2, 2}
		sortColors(input)
		if !reflect.DeepEqual(want, input) {
			t.Errorf("got: %v, want: %v", input, want)
		}
	})
}

func sortColors(nums []int) {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	for left < right {
		if nums[left] == 2 || nums[left] == 1 {
			if nums[right] == 0 {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
			right--
		} else {
			left++
		}
	}
	right = numsLen - 1
	for left < right {
		if nums[left] == 2 {
			if nums[right] == 1 {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
			right--
		} else {
			left++
		}
	}
}
