package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 131. Palindrome Partitioning ", func(t *testing.T) {
		input := "aab"
		want := [][]string{{"aa", "b"}, {"a", "a", "b"}}
		got := solution(input)
		if StrSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	回文处理
	给定一个字符串，将该字符串进行划分，使得所有字符串子集都是回文
*/
func solution(s string) [][]string {
	retArr := make([][]string, 0)
	sLen := len(s)
	solutionHelper(s, sLen, []string{}, &retArr, 0)
	return retArr
}

/*
	从所有的位置都开始一次，然后跟剩下的位置组成的字符串都判断一次
*/
func solutionHelper(s string, sLen int, solStr []string, retArr *[][]string, start int) {
	// 已经遍历到最后的字符
	if start == sLen {
		*retArr = append(*retArr, append([]string{}, solStr...))
	} else {
		saLen := len(solStr)
		// 没有超过总和且不超过数组长度访问
		// 这里，每个元素都最为一个开始，然后对每个元素位置 i，又和所有剩下的元素组成的字符串判断回文
		for i := start; i < sLen; i++ {
			// 这里判断从 start 到 i 的字符串是不是回文
			if isPalindrome(s, start, i) {
				solStr = append(solStr, s[start:i+1])
				solutionHelper(s, sLen, solStr, retArr, i+1)
				// 把最近加入的字符串丢掉构造剩余的不同的组合
				solStr = solStr[:saLen]
			}
		}
	}
}

func isPalindrome(s string, begin, end int) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

func StrSliceEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
