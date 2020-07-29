package jianzhi_offer

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("48 不用加减乘除做加法", func(t *testing.T) {
		a, b := 3, 5
		fmt.Println(a & b)
		get := solution(a, b)
		want := 8
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
	直接位运算
		无进位和是需要将 a b 分别位异或，有进位则需要将 a b 分别位与
*/
func solution(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	// 进位是 0 就停止
	for b != 0 {
		// 计算进位 同时这里进行左移 1 位
		c := (a & b) << 1
		// 非进位和
		a ^= b
		// b 进位
		b = c
	}
	return a
}
