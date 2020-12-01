package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("962. Maximum Width Ramp", func(t *testing.T) {
		A := []int{3, 1, 3, 6}
		want := false
		got := solution(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	使用 stack 单调递减栈
	使用二分查找找第一个更小的元素
	当数字小于上一个，就入栈
*/
func solution(A []int) int {
	// 保存数组索引
	st := make(stack, 0)
	ALen := len(A)
	res := 0
	// 直接构建单调递减栈
	for i := range A {
		if st.Empty() || A[st.Peek()] > A[i] {
			st.Push(i)
		}
	}
	// 从最后的元素向前遍历 单调递减栈，栈顶就是最小的数值索引
	for i := ALen - 1; i > res; i-- {
		// 非空以及当前数值大于等于栈顶元素的时候更新结果
		for !st.Empty() && A[st.Peek()] <= A[i] {
			tempRes := i - st.Pop()
			if tempRes > res {
				res = tempRes
			}
		}
	}
	return res
}

/*
	自定义 stack
*/
type stack []int

func (s *stack) Push(val int) { *s = append(*s, val) }
func (s stack) Peek() int     { return s[len(s)-1] }
func (s stack) Len() int {
	return len(s)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return x
}
func (s stack) Empty() bool {
	return len(s) == 0
}

/*
	找最大的两个元素的索引距离 i, j，要求索引对应数组数值
	满足 A[i] <= A[j] 这个速度太慢了
*/
func solution2(A []int) int {
	ALen := len(A)
	res := 0
	// 注意这里从开始位置和最后的位置向前找
	// 所以如果差值
	for i := 0; i < ALen; i++ {
		// 如果当前位置差值已经小于最大的值了，直接跳出循环
		if ALen-i < res {
			break
		}
		for j := ALen - 1; j > i; j-- {
			if j-i < res {
				break
			}
			if A[i] <= A[j] {
				if j-i > res {
					res = j - i
					break
				}
			}
		}
	}
	return res
}
