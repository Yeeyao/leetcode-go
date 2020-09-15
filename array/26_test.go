package array

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("26. Remove Duplicates from Sorted Array", func(t *testing.T) {
		input := []int{1, 1, 2}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("26. Remove Duplicates from Sorted Array2", func(t *testing.T) {
		input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		want := 5
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	更快一些，比下面的快
*/
func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	dupCount := 0
	for i := 0; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			dupCount++
		} else {
			nums[i-dupCount] = nums[i]
		}
	}
	return numsLen - dupCount
}

/*
	参考 80
*/
func removeDuplicates(nums []int) int {
	i := 0
	for _, n := range nums {
		if i < 1 || n > nums[i-1] {
			nums[i] = n
			i++
		}
	}
	return i
}

func solution(input []int) int {
	inputLen := len(input)
	if inputLen == 0 {
		return 0
	}
	count := 0
	for i := 0; i < inputLen-1; i++ {
		input[i-count] = input[i]
		for i < inputLen-1 && input[i] == input[i+1] {
			i++
			count++
		}
	}
	// 最后一个元素处理
	input[inputLen-count-1] = input[inputLen-1]
	fmt.Println(input)
	return inputLen - count
}

/*
	LeetCode 运行时间和上面的一样
*/
func solution(input []int) int {
	if len(input) < 2 {
		return len(input)
	}

	i := 0
	// 这里利用一个元素来作为当前处理的唯一的元素
	// 如果是相等则跳过处理，否则，需要存放并递增
	for _, v := range input[1:] {
		if v == input[i] {
			continue
		}
		input[i+1] = v
		i++
	}
	return i + 1
}
