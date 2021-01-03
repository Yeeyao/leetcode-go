package math

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("263. Ugly Number", func(t *testing.T) {
		num := 9
		want := true
		got := solution(num)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	263. Ugly Number
	分解的因子仅仅是 2，3，5 的数字就是丑数
	循环判断直到数字是 1，每次循环需要 2，3，5 三个任意一个可以整除，如果都不能整除就直接返回 false
*/
func solution(num int) bool {
	if num == 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return true
}
