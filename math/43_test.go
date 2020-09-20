package math

import "strconv"

/*
Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.

给定两个非负整型字符串，返回它们的乘积字符串
正常乘起来，使用额外数组保存当前的进位，同时额外数组同样需要处理相加
n1Len + n2Len 是结果的最大长度
感觉遍历比较长的比较好
*/

/*
[ref](https://leetcode-cn.com/problems/multiply-strings/solution/zi-fu-chuan-xiang-cheng-by-leetcode-solution/)
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	// 这里每个数字的相乘保存到临时结果 i + j + 1 保存结果
	for i := m - 1; i >= 0; i-- {
		n1 := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			n2 := int(num2[j]) - '0'
			res[i+j+1] += n1 * n2
		}
	}
	// 进位计算
	for i := m + n - 1; i > 0; i-- {
		res[i-1] += res[i] / 10
		res[i] %= 10
	}
	// 字符串计算
	resStr := ""
	index := 0
	// 如果最高位是 0 就不需要加入，不然就有前导 0
	if res[0] == 0 {
		index = 1
	}
	for ; index < m+n; index++ {
		resStr += strconv.Itoa(res[index])
	}
	return resStr
}

/*
	太长了，弃之
*/
//func multiply(num1 string, num2 string) string {
//	var short, long string
//	n1Len, n2Len := len(num1), len(num2)
//	if n1Len < n2Len {
//		short, long = num1, num2
//	} else {
//		short, long = num2, num1
//	}
//	temp := make([]int, n1Len+n2Len)
//	mul := 1
//	intLong, _ := strconv.Atoi(long)
//	for _, s := range short {
//		ts := mul * (int)(s-'0') * intLong
//		Add(ts, temp)
//		mul *= 10
//	}
//	mul = 1
//	for i := 0; i < len(temp) && temp[i] != 0; i++ {
//
//	}
//	return strconv.Itoa(temp)
//}
//
//func Add(a int, intSlice []int) {
//	carry, count := 0, 0
//	sum, add := 0, 0
//	for a > 0 {
//		sum = a + intSlice[count] + carry
//		add = sum % 10
//		intSlice[count] = add
//		carry = sum / 10
//		count++
//		a /= 10
//	}
//	for carry > 0 {
//		sum = intSlice[count] + carry
//		add = sum % 10
//		intSlice[count] = add
//		carry = sum / 10
//	}
//}
