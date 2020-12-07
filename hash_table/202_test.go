package hash_table

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("202. Happy Number", func(t *testing.T) {
		n := 19
		want := true
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	202. Happy Number
	求一个数字的所有位的平方和，然后作为新的数字，一直循环知道遇到重复的数字或者 1 停止
	停止的时候是 1 就返回 true，否则返回 false
*/
func solution(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		seen[n] = true
		sum := calcSum(n)
		n = sum
		if _, ok := seen[n]; ok {
			return false
		}
	}
	if n == 1 {
		return true
	}
	return false
}

func calcSum(n int) int {
	sum := 0
	for n > 0 {
		i := n % 10
		sum += i * i
		n /= 10
	}
	return sum
}
