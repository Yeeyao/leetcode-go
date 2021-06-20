package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("374. Guess Number Higher or Lower", func(t *testing.T) {
		n := 10
		got := solution(n)
		want := 6
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
// 又是变相二分，只是需要调用它们的函数
func solution(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
