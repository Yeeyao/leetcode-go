package hash_table

import (
	"testing"
)

/*
排序找相等
鸽笼原理
下一种方法基于一种有趣的观察，我称之为reduce：把一个整个数组的性质（N个重复数字，不妨设为x；其他N个数字都是不重复的）
缩小到其中一部分元素的性质。在所有的长度为4的子序列中，其中必然至少有一个出现了两个x。
原因是这样的：考虑把这个子序列分成两个长度为2的子序列：由鸽巢原理，其中必然至少有一个数字是x。
但是单看长度为2的子序列是无法判断的（因为x只会出现N次，无法保证一定会出现包含两个x的子序列），所以只能看长度为4的子序列。
如果至少有一个出现了两个x的长度为2的子序列，则包含这个子序列的长度为4的子序列必然也至少包含两个x；
如果每个长度为2的子序列都只有一个x，那么仍然至少存在一个出现了两个x的长度为4的子序列

由上，直接找一开始出现次数是2的元素即所求
另一种方式是可以枚举出4个元素的排列，根据性质来计算
*/

func TestPro(t *testing.T) {
	t.Run("leetcode 961 N-Repeated Element in Size 2N Array", func(t *testing.T) {
		input := []int{8, 2, 3, 3}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// 因为对于 2N 长度来说，总有元素长度是重复 N 次的，所以只需要判断前面的
func solution(A []int) int {
	numCounter := make(map[int]int)
	for _, val := range A {
		numCounter[val]++
		if numCounter[val] >= 2 {
			return val
		}
	}
	return 0
}

// 这种更快，为什么呢？
func repeatedNTimes(A []int) int {
	numCounter := map[int]int{}
	for _, val := range A {
		numCounter[val]++
		if numCounter[val] >= 2 {
			return val
		}
	}
	return 0
}

func repeatedNTimes(A []int) int {
	for i := 2; i < len(A); i++ {
		if (A[i] == A[i-1]) || A[i] == A[i-2] {
			return A[i]
		}
		return A[0]
	}
}
