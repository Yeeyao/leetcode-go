package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 926. Flip String to Monotone Increasing ", func(t *testing.T) {
		input := "00110"
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定只含有 0，1 的字符串，需要让字符串满足升序，求需要 0，1 转换的最小次数
	第二种方法的优化，遍历时直接取较小值
	Co 记录 1 的数量，Cf 记录 0 以及当前的最小反转次数
	当 1 的数量小于 Cf 的时候，表示将 1 反转为 0 的次数更少，所以将 Cf 变成 Co
	此时，假定数组将前面的 1 都反转为 0 了，那后面的 Cf 继续增加就还是目前的 0
	的数量，Co 一样只统计 1 的数量。同时，Co 不管怎么增加，因为反转 Cf 次已经将
	前面的 1 都变成了 0 了，所以后面的 1 增加也符合题目条件。
	直到 0 的数量再次超过 1，
*/
func solution(S string) int {
	countOne, countFlip := 0, 0
	for _, v := range S {
		if v == '1' {
			countOne++
		} else {
			countFlip++
		}
		if countOne < countFlip {
			countFlip = countOne
		}
	}
	return countFlip
}

/*
	dp 方法 可以利用两个数组，分别记录在所有位置上原数组前面将 1 转化为 0 的次数和原数组后面
	将 0 转换为 1 的次数。然后最终结果是找到这两个次数和的最小值。即某个位置上和最小
	这种应该是可以想出来的
	这种方法是每个位置的左右转换次数都记录然后找最小值
*/
func solution2(S string) int {
	sLen := len(S)
	arr1 := make([]int, sLen)
	arr2 := make([]int, sLen)
	for i, j := 1, sLen-1; i < sLen; i, j = i+1, j-1 {
		if S[i-1] == '1' {
			arr1[i] += arr2[i-1] + 1
		}
		if S[j] == '0' {
			arr2[j] += arr2[j+1] + 1
		}
	}
	minFlip := sLen
	for i := 0; i < sLen; i++ {
		tempSum := arr1[i] + arr2[i]
		if tempSum < minFlip {
			minFlip = tempSum
		}
	}
	return minFlip
}

/*
	prefix 方法
	遍历，跳过第一个 1 前面的 0
	用 onesCount 记录下 1 的数量 (prefix)
	我们遇到 1 之后的 0 可能需要被反转，用 flipCount 记录
	如果 flipCount 超过了 onesCount
	应该将 0 反转为 1
	但是跟第一个方法很类似，甚至第一个方法还不用那么麻烦
*/
