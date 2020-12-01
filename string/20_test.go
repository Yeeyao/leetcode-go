package string

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.
使用栈，然后遇到左边的括号就入栈，遇到右边的括号，
	栈为空或者栈顶元素不等于左边的括号就返回 false
	否则，将栈顶元素出栈
最后栈空则返回 true
使用辅助的 map 来匹配括号
*/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
