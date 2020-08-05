package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("44 翻转单词顺序", func(t *testing.T) {
		s := "the sky is blue"
		get := solution(s)
		want := "blue is sky the"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序2", func(t *testing.T) {
		s := "  hello world!  "
		get := solution(s)
		want := "world! hello"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序3", func(t *testing.T) {
		s := "hello world!"
		get := solution(s)
		want := "world! hello"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序4", func(t *testing.T) {
		s := "the sky is blue"
		get := solution3(s)
		want := "blue is sky the"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序5", func(t *testing.T) {
		s := "  hello world!  "
		get := solution3(s)
		want := "world! hello"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序6", func(t *testing.T) {
		s := "hello world!"
		get := solution3(s)
		want := "world! hello"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序7", func(t *testing.T) {
		s := "the "
		get := solution3(s)
		want := "the"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序8", func(t *testing.T) {
		s := " the "
		get := solution3(s)
		want := "the"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("44 翻转单词顺序9", func(t *testing.T) {
		s := " the "
		get := solution3(s)
		want := "the"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序10", func(t *testing.T) {
		s := "e "
		get := solution3(s)
		want := "e"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序11", func(t *testing.T) {
		s := " e"
		get := solution3(s)
		want := "e"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("44 翻转单词顺序12", func(t *testing.T) {
		s := " e "
		get := solution3(s)
		want := "e"
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。
	例如输入字符串"I am a student. "，则输出"student. a am I"。

	无空格字符构成一个单词。
	输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个

	同 leetcode 115

	brute force 直接反向遍历字符串，需要额外的空间保存新的字符串
	利用双指针

*/
func solution(s string) string {
	var res []byte
	bytes := []byte(s)
	bsLen := len(bytes)
	k := bsLen - 1
	for k >= 0 {
		// 过滤掉结尾多余空格
		for k >= 0 && bytes[k] == ' ' {
			k--
		}
		// 开头的空格
		if k < 0 {
			break
		}
		// 记录单词的开始和结尾
		i, j := k, k
		for i >= 0 && bytes[i] != ' ' {
			i--
		}
		// 保存单词并加上空格
		res = append(res, bytes[i+1:j+1]...)
		res = append(res, ' ')
		// 当前位置继续向前
		k -= j - i
	}
	// 去掉结尾空格
	res = res[0 : len(res)-1]
	return string(res)
}

func solution2(s string) string {
	var res []byte
	bytes := []byte(s)
	bsLen := len(bytes)
	i, j := bsLen-1, bsLen-1
	// 过滤开头k空格
	for i >= 0 && bytes[i] == ' ' {
		i--
		j--
	}
	for i >= 0 {
		// 找到首个非空格
		for i >= 0 && bytes[i] != ' ' {
			i--
		}
		res = append(res, bytes[i+1:j+1]...)
		res = append(res, ' ')
		// 过滤掉单词间空格
		for i >= 0 && bytes[i] == ' ' {
			i--
		}
		j = i
	}
	// 去掉结尾空格
	res = res[0 : len(res)-1]
	return string(res)

}

func solution3(s string) string {
	var res []byte
	sByte := []byte(s)
	right := len(sByte) - 1
	for right >= 0 {
		// 过滤结尾开头的空格
		for right >= 0 && sByte[right] == ' ' {
			right--
		}
		if right < 0 {
			break
		}
		// 遍历单词
		end := 0
		for right >= 0 && sByte[right] != ' ' {
			end++
			right--
		}
		// 保存单词以及添加结尾
		res = append(res, sByte[right+1:right+end+1]...)
		res = append(res, ' ')
	}
	// 去除尾部空格
	res = res[0 : len(res)-1]
	return string(res)
}
