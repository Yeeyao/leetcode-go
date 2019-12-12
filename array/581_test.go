package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("581. Shortest Unsorted Continuous Subarray", func(t *testing.T) {
		input := []int{2, 6, 4, 8, 10, 9, 15}
		want := 5
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("581. Shortest Unsorted Continuous Subarray2", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		want := 0
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("581. Shortest Unsorted Continuous Subarray3", func(t *testing.T) {
		input := []int{2, 1}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找到最短连续的未排序的子数组，该数组有序则整体数组有序
	从左到右遍历，找到比当前 max 小的最近一个元素
	从右到做遍历，找到比当前 min 大的最近一个元素
*/
func solution(input []int) int {
	inputLen := len(input)
	// subarray index
	left, right := -1, -2
	// 递增
	min, max := input[inputLen-1], input[0]
	for i := 1; i < inputLen; i++ {
		rightIndex := inputLen - i - 1
		// 先更新最大最小值
		if input[i] > max {
			max = input[i]
		}
		if input[rightIndex] < min {
			min = input[rightIndex]
		}
		// 更新子数组的开始结束位置索引
		/*
			这里的意思是，需要排序的数组的结束位置，是在当前最大值的左边的
			开始位置，是在当前最小值的右边的。所以需要不断更新
		*/
		if input[i] < max {
			right = i
		}
		if input[rightIndex] > min {
			left = rightIndex
		}
	}
	return right - left + 1
}

func solution(input []int) int {
	inputLen := len(input)
	if inputLen == 0 || inputLen == 1 {
		return 0
	}
	// 找到比最大值小的最右边元素
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	max, end := intMin, -2
	for i, v := range input {
		if v > max {
			max = v
		}
		if v < max {
			end = i
		}
	}
	// 找到比最小值大的最左边元素
	min, begin := intMax, -1
	for i := inputLen - 1; i >= 0; i-- {
		if input[i] < min {
			min = input[i]
		}
		if input[i] > min {
			begin = i
		}
	}
	return end - begin + 1
}
