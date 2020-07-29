package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("47 求1+2+…+n", func(t *testing.T) {
		n := 5
		get := solution(n)
		want := 15
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
	需要将 n 分解成 2^1 + 2^2 + ... 然后将 sum 对应左移并累加
	但这种思路只是将乘除换成了移位操作，循环等好像还是需要处理
	递归，然后将循环缓存逻辑运算符

	func sum(n int) int {
		if n == 0 {
			return 0
		}
		return n + sum(n-1)
	}

*/
// 递归
var res int

func solution(n int) int {
	res = 0
	sum(n)
	return res
}

func sum(n int) bool {
	res += n
	return n > 0 && sum(n-1)
}
