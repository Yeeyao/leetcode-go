package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("31  1～n整数中1出现的次数", func(t *testing.T) {
		n := 12
		get := solution(n)
		want := 5
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	需要使用公式处理
	数字 n 是 x 位数 ni 为第 i 位 n 可以记为 nxnx-1...n3n2n1
	ni 当前位记为 cur，其中 ni-1ni-2...n2n1 称为低位 nxnx-1...ni 称为高位
	10^i 称为位因子 digit
	分类
		cur = 0 1 出现的次数只由高位 high 决定 为 high * digit
		cur = 1 1 出现的次数计算公式 high * digit + low + 1
		cur = 2-9 (high + 1) * digit
	需要注意初始值
		high = n // 10 cur = n % 10 low = 0 digit = 1
	递归下去处理
		结束条件是 high == 0 以及 cur == 0
	循环变化
		high = n // 10
		cur = n % 10
		low += cur * digit
		digit *= 10
*/
func solution(n int) int {
	high, low, cur, digit := n/10, n%10, 0, 1
	count := 0
	for cur != 0 || high != 0 {
		if cur == 0 {
			count += high * digit
		} else if cur == 1 {
			count += high*digit + low + 1
		} else {
			count += (high + 1) * digit
		}
		cur = high % 10
		high = high / 10
		low += cur * digit
		digit *= 10
	}
	return count
}
