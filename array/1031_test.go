package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 1031. Maximum Sum of Two Non-Overlapping Subarrays ", func(t *testing.T) {
		A := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
		L, M := 1, 2
		want := 20
		got := solution2(A, L, M)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	题目的意思是找一个数组内的两个指定长度的无重叠的子数组
	使得两个子数组的元素总和最大

	首先，会不会出现先找长度较长的最大和子数组，然后找较短的得到的结果和反过来的结果不一样的情况

	如果先找长的子数组，出现有相同的和的情况，则还要进行不同情况的区分 9467 找 1，2 长度的。
	如果先找短的子数组，12395 找 1，3 长度的

	先各自找最大的子数组，看是否有重复部分，有重复，那就需要其中一个移动来判断。没有重复则直接找到了

	类似 购买股票的题目编号：123。
	从左到右，找最大的 L 长度的和
	从右到左，找最大的 M 长度的和
	上述过程需要重复一次，因为 L 和 M 都要从两个方向执行一次

	从两边遍历，并且需要遍历完。先加上当前元素。
	如果当前的元素比前面的元素总和都大的话，当前总和需要更新为当前元素
	如果当前的遍历数量已经超过了给定的长度，则需要减去最早加入的元素值。
	注意，这里使用两个额外的数组，保存每个子数组到当前位置的和。
	最后，比较所有左右两边的总和得到最大总和
	左右两边都需要计算来取较大者

*/

func solution2(A []int, L, M int) int {
	first := solution2_2(A, L, M)
	second := solution2_2(A, M, L)
	if first > second {
		return first
	}
	return second
}

func solution2_2(A []int, L, M int) int {
	aLen := len(A)
	// 使用两个辅助数组保存当前元素下的累加和
	leftSumArr := make([]int, aLen+1)
	rightSumArr := make([]int, aLen+1)
	for i, j, leftSum, rightSum := 0, aLen-1, 0, 0; i < aLen; i, j = i+1, j-1 {
		leftSum += A[i]
		rightSum += A[j]
		// 如果当前的元素已经大于累加的和则保存当前的元素
		if leftSumArr[i] > leftSum {
			leftSumArr[i+1] = leftSumArr[i]
		} else {
			leftSumArr[i+1] = leftSum
		}
		if rightSumArr[j+1] > rightSum {
			rightSumArr[j] = rightSumArr[j+1]
		} else {
			rightSumArr[j] = rightSum
		}
		// 超过给定长度，需要减去最早加入的元素
		// 因为上面是先加上去计算，所以这里的边界需要判断等于
		if i+1 >= L {
			leftSum -= A[i+1-L]
		}
		if i+1 >= M {
			rightSum -= A[j-1+M]
		}
	}
	maxSum := 0
	for i := range A {
		tempSum := leftSumArr[i] + rightSumArr[i]
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}
