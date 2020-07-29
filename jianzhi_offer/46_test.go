package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("46 圆圈中最后剩下的数字", func(t *testing.T) {
		n := 5
		m := 3
		get := solution(n, m)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。
	例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
	约瑟夫问题
	智力题
	长度为 n 的序列将删除 m % n 个元素，记为 f(n, m)，下一步将求 f(n-1, m)
	f(n,m) 开始的时候，需要移除的数字下标是 (m-1) % n，继续向后面数的下标
	((m-1)%n + x + 1) % n， x = f(n-1, m)
	f(n,m) = ((m-1)%n + x + 1) % n = (m + x) % n
	TODO:
*/
func solution(n, m int) int {
	ret := 0
	for i := 2; i < n+1; i++ {
		ret = (m + ret) % i
	}
	return ret
}
