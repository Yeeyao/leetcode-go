package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("9 变态跳台阶", func(t *testing.T) {
		n := 3
		want := 4
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
转换为找出所有子数组和是 num 的子数组数量
*/
func solution(num int) int {
	res := 0
	solutionHelper(num, 0, &res)
	return res
}

func solutionHelper(num, sum int, res *int) {
	if sum == num {
		*res = *res + 1
		return
	} else if sum > num {
		return
	}
	for i := 1; i < num + 1; i++ {
		sum += i
		solutionHelper(num, sum, res)
		sum -= i
	}
}
