package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1184. Distance Between Bus Stops", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		start := 0
		destination := 1
		want := 1
		got := solution(input, start, destination)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1184. Distance Between Bus Stops2", func(t *testing.T) {
		input := []int{5, 3, 5, 7}
		start := 0
		destination := 2
		want := 8
		got := solution(input, start, destination)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1184. Distance Between Bus Stops3", func(t *testing.T) {
		input := []int{7, 10, 1, 12, 11, 14, 5, 0}
		start := 7
		destination := 2
		want := 17
		got := solution(input, start, destination)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	distance == n
*/
func solution(input []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	totalSum := calcSum(input)
	cwArr := input[start:destination]
	cwSum := calcSum(cwArr)
	ccwSum := totalSum - cwSum
	if cwSum < ccwSum {
		return cwSum
	} else {
		return ccwSum
	}
}

func calcSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
