package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 856. Score of Parentheses ", func(t *testing.T) {
		input := "(()(()))"
		want := 6
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run(" 856. Score of Parentheses2", func(t *testing.T) {
		input := "(()(()))"
		want := 6
		got := solution2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run(" 856. Score of Parentheses3 ", func(t *testing.T) {
		input := "(()(()))"
		want := 6
		got := solution3(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	计算括号对的数量 () 为 1 AB 为 A + B (A) 为 2 * A
	使用一个计数器，遇到左括号，计数器 + 1，遇到右括号，计数器 - 1。当计数器是 0
	的时候，就找到一对了
	遇到左括号则直接入栈，且设置 cur = 0 相邻的一对括号直接相加，但是需要判断嵌套
	遇到右括号则需要将 cur 成倍 或者至少为 1

*/
func solution(S string) int {
	SLen := len(S)
	// 栈保存的是 cur，当前并列的数值
	st := make([]int, SLen)
	stTop, cur := 0, 0
	for i := 0; i < SLen; i++ {
		// 空或者左括号，入栈，
		// 这里将上一次的数值保存
		if S[i] == '(' {
			st[stTop] = cur
			stTop++
			// 只有遇到右括号才需要保存下当前的数值，所以这里每次都将 cur 置 0
			cur = 0
			// 更新 cur 并出栈一个左括号
		} else {
			// 遇到多个右括号，需要将个数都加倍
			if cur > 0 {
				cur = cur * 2
				// 只遇到一个右括号，只需要将个数置 1
			} else {
				cur = 1
			}
			// 这里是并列关系，所以需要加
			cur = st[stTop-1] + cur
			stTop--
		}
	}
	return cur
}

/*
	第二种理解 只是使用数组而已
	向前累加
*/
func solution2(S string) int {
	// 输入最长只有 50
	res := make([]int, 31)
	i := 0
	for _, c := range S {
		if c == '(' {
			res[i] = 0
			i++
		} else {
			// 这里是向前赋值
			if res[i] == 0 {
				res[i-1] += 1
			} else {
				res[i-1] += res[i] * 2
			}
			i--
		}
	}
	return res[0]
}

/*
	只统计括号数量 遇到左括号，计数 + 1，遇到右括号，计数 - 1
	刚好遇到一对则需要累加，其中，有多少层左括号就需要加倍
	同时，嵌套只计算最里面的那一层，外面的不符合计数条件
	更像是智力题。。。
*/
func solution3(S string) int {
	res, count := 0, 0
	SLen := len(S)
	for i := 0; i < SLen; i++ {
		if S[i] == '(' {
			count++
		} else {
			count--
		}
		if S[i] == ')' && S[i-1] == '(' {
			// 累加
			res += 1 << count
		}
	}
	return res
}
