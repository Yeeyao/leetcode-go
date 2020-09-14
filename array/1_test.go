package array

import "testing"

/*
1. Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

给定整型数组以及 target 整型，返回两个数的下标，要求它们的和是 target
*/
func TestPro(t *testing.T) {
	t.Run("1. Two Sum", func(t *testing.T) {
		input := []int{2, 7, 11, 15}
		target := 9
		want := []int{0, 1}
		got := solution(input, target)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	对每个元素，将差保存在另一个数组中，
	如果当前元素在数组中找到了，则直接返回起始元素位置以及当前元素位置
*/
func solution(input []int, target int) []int {
	inputLen := len(input)
	diffMap := make(map[int]int)
	for i := 0; i < inputLen; i++ {
		v := input[i]
		if pos, ok := diffMap[v]; ok {
			return []int{pos, i}
		}
		dest := target - v
		diffMap[dest] = i
	}
	return []int{}
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
