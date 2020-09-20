package math

import "strconv"

/*
67. Add Binary
Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.

给定两个二进制字符串，返回它们的和
取两个字符串的较长作为遍历次数，然后 carry 保存当前的结果
每次判断遍历是否处在两个字符的长度区间，是就将数值加上去，然后保存结果并更新进位
遍历完需要再判断一次进位并更新数值
[ref](https://leetcode-cn.com/problems/add-binary/solution/er-jin-zhi-qiu-he-by-leetcode-solution/)
*/

func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}
	aLen, bLen := len(a), len(b)
	longer, carry := 0, 0
	if aLen > bLen {
		longer = aLen
	} else {
		longer = bLen
	}
	res := ""
	// 从最后的开始加
	for i := 0; i < longer; i++ {
		if i < aLen {
			carry += int(a[aLen-i-1] - '0')
		}
		if i < bLen {
			carry += int(b[bLen-i-1] - '0')
		}
		// 这里加到头部
		res = strconv.Itoa(carry%2) + res
		carry = carry / 2
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}
