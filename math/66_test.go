package math

/*
66. Plus One
Given a non-empty array of digits representing a non-negative integer, increment one to the integer.
The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contains a single digit.
You may assume the integer does not contain any leading zero, except the number 0 itself.

给定非空数组表示一个非负整型，将该整型数字 + 1 数组最高位在最前面，没有前导 0
返回 + 1 后的数组

先在最后的数字位置 + 1 然后不断判断最后的位置是否 > 9 如果是就不断将前面的进位，同时当前为置 0
*/

func plusOne(digits []int) []int {
	dLen := len(digits)
	digits[dLen-1]++
	for i := dLen - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i-1]++
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] > 9 {
		digits = append([]int{1}, digits...)
		digits[1] = 0
	}
	return digits
}
