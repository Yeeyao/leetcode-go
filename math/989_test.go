package math

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("989. Add to Array-Form of Integer", func(t *testing.T) {
		input := []int{1, 2, 0, 0}
		k := 34
		want := []int{1, 2, 3, 4}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer2", func(t *testing.T) {
		input := []int{2, 7, 4}
		k := 181
		want := []int{4, 5, 5}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer3", func(t *testing.T) {
		input := []int{2, 1, 5}
		k := 806
		want := []int{1, 0, 2, 1}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer4", func(t *testing.T) {
		input := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
		k := 1
		want := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer5", func(t *testing.T) {
		input := []int{2, 1, 5}
		k := 8061
		want := []int{8, 2, 7, 6}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer6", func(t *testing.T) {
		input := []int{0}
		k := 10000
		want := []int{1, 0, 0, 0, 0}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("989. Add to Array-Form of Integer7", func(t *testing.T) {
		input := []int{}
		k := 0
		want := []int{0}
		got := solution(input, k)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里是先构建结果数组，然后按顺序处理
	这里的 brute force 快一点
*/
func solution(input []int, k int) []int {
	inputLen := len(input)
	kLen := calcLen(k)
	maxLen := 0
	if inputLen > kLen {
		maxLen = inputLen + 1
	} else {
		maxLen = kLen + 1
	}
	retArr := make([]int, maxLen)
	carry := 0
	for i := 0; i < maxLen; i++ {
		kAdd := k % 10
		k /= 10
		ri := maxLen - i - 1
		j := inputLen - i - 1
		// maxLen 可能会大于 inputLen
		if j >= 0 {
			retArr[ri] = input[j] + kAdd + carry
		} else {
			retArr[ri] = kAdd + carry
		}
		carry = retArr[ri] / 10
		retArr[ri] = retArr[ri] % 10
	}

	if retArr[0] == 0 {
		retArr = retArr[1:]
	}
	return retArr
}

func calcLen(k int) int {
	kLen := 1
	for k/10 > 0 {
		kLen++
		k /= 10
	}
	return kLen
}

/*
	优化 不需要计算长度 只需要 append
	更加简洁，比我的慢一些。。。
*/
func solution(input []int, K int) []int {
	res := make([]int, 0)
	// 进位
	c := 0
	// 两部分一起判断
	for i := len(input) - 1; i >= 0 || K > 0; i-- {
		s := 0
		// 数组部分还有
		if i >= 0 {
			s = s + input[i]
		}
		s = s + K%10 + c
		c = s / 10
		K = K / 10
		res = append(res, s%10)
	}
	// 最后一个进位
	if c != 0 {
		res = append(res, c)
	}
	// 反转
	resLen := len(res)
	for i := 0; i < resLen/2; i++ {
		res[i], res[resLen-i-1] = res[resLen-i-1], res[i]
	}
	return res
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
