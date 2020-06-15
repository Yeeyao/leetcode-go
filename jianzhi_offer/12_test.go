package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("12 数值的整数次方", func(t *testing.T) {
		base := 2.0
		exponent := 5
		want := 32.0
		got := solution(base, exponent)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("12 数值的整数次方", func(t *testing.T) {
		base := 2.0
		exponent := 8
		want := 256.0
		got := solution(base, exponent)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	leetCode 50
	给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
	保证base和exponent不同时为0
	二分法 x^n = x^2 * x^(n // 2) 将 n 不断减少到 0
	n 为偶数, x^n = x^2 * x^(n // 2) n 为奇数，x^n = x * x^2 * x^(n//2)
	res = 1，初始 x^n = x^n * res 当 n 为奇数时，将多出来的一项 x 乘入 res，最终可以 x^n = x^0 * res = res
*/
func solution(base float64, exponent int) float64 {
	if base == 0 {
		return 0.0
	}
	res := 1.0
	// 负数指数 x^-n 转化为 (1/x)^n
	if exponent < 0 {
		base, exponent = 1/base, -exponent
	}
	// 位运算替代算术 如果指数对应的位数是 1 就需要将当前的累计乘积乘到 res
	// res = k * x^n + k * x^n-1 + ... + k * x^1
	for exponent > 0 {
		if exponent & 1 > 0 {
			res *= base
		}
		base *= base
		exponent >>= 1
	}
	return res
}
