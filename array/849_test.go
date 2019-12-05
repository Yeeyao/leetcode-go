package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("849. Maximize Distance to Closest Person", func(t *testing.T) {
		input := []int{1, 0, 0, 0, 1, 0, 1}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("849. Maximize Distance to Closest Person2", func(t *testing.T) {
		input := []int{1, 0, 0, 0}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("849. Maximize Distance to Closest Person3", func(t *testing.T) {
		input := []int{0, 0, 0, 1}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("849. Maximize Distance to Closest Person4", func(t *testing.T) {
		input := []int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	中间元素可通过计算两个 1 之间的元素数量来判断
	只需要记录开始和结束元素
	找到一个空位
*/
func solution(input []int) int {
	inputLen := len(input)
	var tempMax int
	maxLen := 0
	onePos := -1
	for i := 0; i < inputLen; i++ {
		if input[i] == 1 {
			if onePos == -1 {
				tempMax = i
			} else {
				tempMax = (i - onePos) / 2
			}
			if tempMax > maxLen {
				maxLen = tempMax
			}
			onePos = i
		}
	}
	// 这里处理只有一个 1 右边的情况
	left := inputLen - onePos - 1
	if left > maxLen {
		maxLen = left
	}
	return maxLen
}
