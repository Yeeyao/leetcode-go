package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1671. Minimum Number of Removals to Make Mountain Array", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*

 */
func solution(nums []int) int {

}
