package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("8 跳台阶", func(t *testing.T) {
		n := 2
		want := 2
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	跳台阶
	一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）
	定义 f(n) 为到 n 台阶时有多少种跳法。跳最后一次时，最后一次和前一步属于同一种
	最后一步的时候，跳上一级或者二级台阶，一级时，有 f(n - 1) 种，二级时，有 f(n - 2) 种。
	直接 DP 了，对于第 n 级，设 f(n) 为到第 n 级的不同跳法，则
	f(n) = f(n - 1) + f(n - 2) f(1) = 1 f(2) = 2
*/
func solution(num int) int {
	a, b := 1, 1
	for i := 0; i < num; i++ {
		a, b = b, a+b
	}
	return a % 1000000007
}

/*
	变种，如果青蛙不能连续跳两次 2级台阶
*/
