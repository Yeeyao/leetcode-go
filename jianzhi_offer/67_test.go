package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("67 剪绳子", func(t *testing.T) {
		n := 2
		get := solution(n)
		want := 1
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("67 剪绳子2", func(t *testing.T) {
		n := 10
		get := solution(n)
		want := 36
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
根据数学推导，分成长度为 3 的每个小段d得到最大的乘积
	n <= 3 返回 n - 1
	n > 3
		n / 3 得到 a 余数 b
			b == 0 直接返回 3^a
			b == 1 返回 3^(a-1)*4
			b == 2 返回 3^a * 2
*/
func solution(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return pow(3, a)
	} else if b == 1 {
		return pow(3, a-1) * 4
	} else {
		return pow(3, a) * 2
	}
}

/*
	n^m 2 3
*/
func pow(n, m int) int {
	res := 1
	for m > 0 {
		if 1&m > 0 {
			res *= n
		}
		n *= n
		m >>= 1
	}
	return res
}
