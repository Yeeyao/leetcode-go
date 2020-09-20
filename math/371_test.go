package math

/*
371. Sum of Two Integers 类似 48 解法一样
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
给定两数字 a, b 返回它们的和，不许使用 + - 运算符
直接位运算
*/
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}
