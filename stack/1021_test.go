package stack

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("leetcode 1021 Remove Outermost Parentheses", func(t *testing.T) {
		input := "(()())(())"
		want := "()()()"
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(str string) string {
	var strslice []rune
	var increment, counter = 0, 0
	var rets string
	for _, v := range str {
		// 需要将中间的元素出栈并存入返回的字符串，所以这里需要进行额外的计数
		if counter == 0 && increment != 0 {
			rets = rets + string(strslice[1:increment-1])
			strslice = strslice[0:0]
			increment, counter = 0, 0
		}
		if v == '(' {
			increment++
			counter++
			strslice = append(strslice, v)
		} else {
			increment++
			counter--
			strslice = append(strslice, ')')
		}
	}
	if counter == 0 && increment != 0 {
		rets = rets + string(strslice[1:increment-1])
	}
	return rets
}

// 因为都是 pairs 所以第一个 '(' counter 一定是 0
// 最后一个 ')' 一定是 1，所以只需要把这两个排除，中间的都保存就行
func removeOuterParentheses(S string) string {
	counter := 0
	formatted_s := ""

	for _, l := range S {
		// ( 且非 0
		if l == '(' {
			if counter != 0 {
				formatted_s += "("
			}
			counter++
		} else {
			// ) 且非 1
			if counter != 1 {
				formatted_s += ")"
			}

			counter--
		}
	}

	return formatted_s
}

// 使用两个 []byte stack 处理 len(stack) 就是 counter
// sub_str 是当前处理的字符串临时数据
// stack 就是用来计数的，当 stack 中 ( 和 ) 数量相等时，表示需要将 sub_str 中的元素保存到字符串中
func removeOuterParentheses(S string) string {
	s_len := len(S)
	if s_len == 0 || s_len == 2 {
		return ""
	}
	var i int
	str := ""
	var stack, sub_str []byte
	for i = 0; i < s_len; i++ {
		// 当前处理的栈
		sub_str = append(sub_str, S[i])
		if S[i] == '(' {
			stack = append(stack, S[i])
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 将中间的保存到字符串中
				str += string(sub_str[1 : len(sub_str)-1])
				// 处理完，清空
				sub_str = sub_str[:0]
			}
		}
	}

	return str
}
