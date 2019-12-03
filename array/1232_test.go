package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	//t.Run("1232. Check If It Is a Straight Line", func(t *testing.T) {
	//	input := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	//	want := true
	//	got := solution(input)
	//	if got != want {
	//		t.Errorf("got: %v, want: %v", got, want)
	//	}
	//})
	//
	//t.Run("1232. Check If It Is a Straight Line2", func(t *testing.T) {
	//	input := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 4}}
	//	want := false
	//	got := solution(input)
	//	if got != want {
	//		t.Errorf("got: %v, want: %v", got, want)
	//	}
	//})

	t.Run("1232. Check If It Is a Straight Line3", func(t *testing.T) {
		input := [][]int{{-4, -3}, {1, 0}, {3, -1}, {0, -1}, {-5, 2}}
		want := false
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1232. Check If It Is a Straight Line4", func(t *testing.T) {
		input := [][]int{{1, -3}, {1, 0}, {1, -1}, {1, -2}, {1, 2}}
		want := true
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(input [][]int) bool {
	first := input[0]
	second := input[1]
	for i := 2; i < len(input); i++ {
		cur := input[i]
		if (cur[0]-first[0])*(first[1]-second[1]) != (cur[1]-first[1])*(first[1]-second[1]) {
			return false
		}
	}
	return true
}

func solution(input [][]int) bool {
	dx0 := input[1][0] - input[0][0]
	dy0 := input[1][1] - input[0][1]
	for i := 1; i < len(input)-1; i++ {
		dx := input[i+1][0] - input[i][0]
		dy := input[i+1][1] - input[i][1]
		if dx0*dy != dy0*dx {
			return false
		}
	}
	return true
}
