package string

import (
	"strconv"
	"strings"
	"testing"
)

/*
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits and
that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].
给定一个编码的字符串，返回解码字符串
编码规则是 k[s] 表示 s 出现 k 次，输入字符串是合法的
*/

func TestPro(t *testing.T) {
	t.Run("394. Decode String", func(t *testing.T) {
		input := "3[a]2[bc]"
		want := "aaabcbc"
		got := decodeString(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
[ref](https://leetcode-cn.com/problems/decode-string/solution/zi-fu-chuan-jie-ma-by-leetcode-solution/)
	字符串从后面向前遍历，如果遇到字母，直接保存，遇到 ] 将其入栈，
	使用栈 将所有字符先入栈，然后
	将字母，数字和括号看成是独立的字符，用栈来维护它们，具体是遍历栈
		如果当前字符是数字，解析出一个数字并进栈
		如果当前字符为字母或者左括号，直接进栈
		如果当前字符是右括号，开始出栈直到遇到左括号，出栈的序列反转拼成字符串，
		取出栈顶的数字，然后构造新的字符串并进栈
	重复操作直到栈元素按照从栈底到栈顶顺序拼接

	初始化栈和当前遍历的字符串索引
	循环条件是 索引小于输入字符串长度
		取当前遍历的索引，根据当前遍历的索引对应的字符进行不同处理
		如果当前是数字，需要将整个数字提取并入栈
		如果当前是字母或者 [ 直接入栈
		如果是 ]
			创建临时字符串 sub
			循环将栈中的元素保存到临时字符串直到遇到 [，然后将临时字符串反转
			将栈中的 [ 出栈
			取栈顶元素作为数字
			将字符串重复后入栈
*/
func decodeString(s string) string {
	stack := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		// 这里将数字都遍历并提取出来
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' || cur == '[' {
			// 字母或者 [ 直接入栈
			stack = append(stack, string(cur))
			ptr++
		} else {
			// 这里需要跳过遍历到的 ]
			ptr++
			sub := []string{}
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stack = stack[:len(stack)-1]
			num, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			add := strings.Repeat(getString(sub), num)
			stack = append(stack, add)
		}
	}
	return getString(stack)
}

// 循环提取出字符串的数字字符
func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

// 将字符串 slice 拼接成字符串
func getString(s []string) string {
	res := ""
	for _, s := range s {
		res += s
	}
	return res
}
