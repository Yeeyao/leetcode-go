package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("53 表示数值的字符串", func(t *testing.T) {
		s := "+100"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串2", func(t *testing.T) {
		s := "5e2"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串3", func(t *testing.T) {
		s := "-123"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串4", func(t *testing.T) {
		s := "5e2"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串5", func(t *testing.T) {
		s := "3.146"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串6", func(t *testing.T) {
		s := "01234"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串7", func(t *testing.T) {
		s := "1.5e220"
		get := solution(s)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串8", func(t *testing.T) {
		s := "1.5e"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串9", func(t *testing.T) {
		s := "15e"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串10", func(t *testing.T) {
		s := "1a3.14"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串11", func(t *testing.T) {
		s := "+-5"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串12", func(t *testing.T) {
		s := "-1E-16"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串13", func(t *testing.T) {
		s := "12E+5.4"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("53 表示数值的字符串14", func(t *testing.T) {
		s := "+1+"
		get := solution(s)
		want := false
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"、"-1E-16"及"12e+5.4"都不是。
同 leetcode 65
	首先需要分析第一部分，开头是 + 或者数字
	中间可以是 . 或者 e
	最后是数字
*/
func solution(s string) bool {
	byteS := []byte(s)
	bsLen := len(byteS)
	// 前面部分的符号以及 e 后面的符号
	sign := 0
	eSign := 0
	i := 0
	// 过滤符号位并记录
	if byteS[i] == '-' {
		sign = 1
		i++
	} else if byteS[i] == '+' {
		sign = 2
		i++
	}
	// 前面的数字
	for i < bsLen && byteS[i] >= '0' && byteS[i] <= '9' {
		i++
	}
	// 只有前面部分
	if i == bsLen {
		return true
	}
	// 前面必须要有数字部分
	if sign != 0 && i < 2 || sign == 0 && i < 1 {
		return false
	}
	// 小数点
	if byteS[i] == '.' {
		// 小数点后面要有数字
		i++
		if i == bsLen {
			return false
		}
		// 小数点后数字部分
		for i < bsLen && byteS[i] >= '0' && byteS[i] <= '9' {
			i++
		}
		// 仅仅是小数
		if i == bsLen {
			return true
		}
	}
	// e 部分
	if byteS[i] == 'e' || byteS[i] == 'E' {
		// e 前面的部分不能是负数
		if sign == 1 {
			return false
		}
		eSign = 1
		i++
	}
	// e 后面必须要有数字
	if i == bsLen {
		return false
	}
	// e 后面的符号
	if byteS[i] == '+' || byteS[i] == '-' {
		// 有 e 才能判断符号
		if eSign == 1 {
			i++
		} else {
			return false
		}
	}
	// e 后的数字部分
	for i < bsLen && byteS[i] >= '0' && byteS[i] <= '9' {
		i++
	}
	// 必须全部是数字
	if i == bsLen {
		return true
	} else {
		return false
	}
}
