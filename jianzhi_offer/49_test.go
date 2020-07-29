package jianzhi_offer

import (
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	b := []byte{'1'}
	fmt.Println(b[0] - '0')
	t.Run("49 把字符串转换成整数", func(t *testing.T) {
		str := "42"
		get := solution(str)
		want := 42
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("49 把字符串转换成整数2", func(t *testing.T) {
		str := "    -42"
		get := solution(str)
		want := -42
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("49 把字符串转换成整数3", func(t *testing.T) {
		str := "4193 with other words"
		get := solution(str)
		want := 4193
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("49 把字符串转换成整数4", func(t *testing.T) {
		str := "words and 42"
		get := solution(str)
		want := 0
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("49 把字符串转换成整数5", func(t *testing.T) {
		str := "-91283472332"
		get := solution(str)
		want := -2147483648
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("49 把字符串转换成整数6", func(t *testing.T) {
		str := "91283472332222"
		get := solution(str)
		want := 2147483647
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，
请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
先将前面的空格过滤，然后判断第一个符号还是数字
数字遍历内部判断最大最小值然后累加
*/
func solution(str string) int {
	nums := []byte(str)
	bsLen := len(nums)
	minus := 1
	res := 0
	if bsLen == 0 {
		return 0
	}
	const intMax = int(^uint32(0) >> 1)
	intMaxMt := intMax / 10
	// 过滤掉开头的空格
	i := 0
	for ; i < bsLen && nums[i] == ' '; i++ {

	}
	// 判断第一个字符
	if nums[i] == '-' {
		minus = -1
		i++
	} else if nums[i] == '+' {
		i++
	}
	for j := i; j < bsLen; j++ {
		// 非数字，直接退出循环
		if nums[j] < '0' || nums[j] > '9' {
			break
		}
		// 最大最小值判断，需要注意这里少算了一位，所以上面需要 / 10
		// 这里相等的时候，判断档期那遍历位是否大于 7，因为最大或者最小最后一位是 8
		if res > intMaxMt || res == intMaxMt && nums[j] > '7' {
			if minus == 1 {
				return intMax
			} else {
				return -intMax - 1
			}
		}
		res = res*10 + int(nums[j]) - '0'
	}
	return res * minus
}
