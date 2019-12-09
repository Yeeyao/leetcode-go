package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("532. K-diff Pairs in an Array", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5}
		k := 2
		want := 2
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("532. K-diff Pairs in an Array2", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		k := 1
		want := 4
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("532. K-diff Pairs in an Array3", func(t *testing.T) {
		input := []int{1, 3, 1, 5, 4}
		k := 0
		want := 1
		got := solution(input, k)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	找到特定和的对数 two-sum 加强版
*/
func solution(input []int, k int) int {
	if k < 0 {
		return 0
	}
	sum := 0
	counter := make(map[int]int)
	for _, v := range input {
		counter[v]++
	}
	for v, w := range counter {
		if k == 0 {
			if w > 1 {
				sum++
			}
		} else {
			// 只需要计算一个
			bigger := v + k
			if _, ok := counter[bigger]; ok {
				sum++
			}
		}
	}
	return sum
}

func buildFreqTable(numbers []int) map[int]int {
	var freqTable map[int]int = make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		freqTable[numbers[i]]++
	}

	return freqTable
}

func findPairs(numbers []int, k int) int {
	if k < 0 || len(numbers) < 1 {
		return 0
	}
	var freqTable map[int]int = buildFreqTable(numbers)

	var count = 0
	for key, value := range freqTable {
		if k == 0 && value >= 2 {
			count++
		} else if k != 0 && freqTable[key+k] > 0 {
			count++
		}
	}
	return count
}
