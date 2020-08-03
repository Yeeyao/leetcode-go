package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("11 二进制中1的个数", func(t *testing.T) {
		n := uint32(3)
		want := 2
		got := solution(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("11 二进制中1的个数2", func(t *testing.T) {
		n := uint32(3)
		want := 2
		got := solution2(n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	leetcode 191
	输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
	负数也需要计算 直接无符号 1 和数字做与运算，如果得到 1 就表示当前的位是 1
	之后需要不断逻辑右移
	brute force 移位然后逐个比较
*/
func solution(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num & uint32(1))
		num >>= 1
	}
	return res
}

/*
	记得好像 CSAPP 的实验有这道题
	n & n - 1 将原来的最后一个 1 变成 0
	1111 & 1110 ---> 1110 1110 & 1101 -- > 1100
	1010 & 1001 -- > 1000 1011 & 1010 -- > 1010
*/
func solution2(num uint32) int {
	res := 0
	for num > 0 {
		res++
		num &= num - 1
	}
	return res
}
