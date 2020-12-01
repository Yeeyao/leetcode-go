/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.
Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/

/*
这里的思路是利用两个指针来处理排序，其中，可以利用已经排序的特性，从两边向中间遍历处理，
这里指针相交则可以停止，最后翻转一下。
也可以从中间选取两个指针，向两边移动来遍历处理。
*/

package array

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("977. Squares of a Sorted Array", func(t *testing.T) {
		input := []int{-4, -1, 0, 3, 10, 20}
		want := []int{0, 1, 9, 16, 100, 400}
		got := solution3(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array2", func(t *testing.T) {
		input := []int{0, 1, 2, 3}
		want := []int{0, 1, 4, 9}
		got := solution3(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array3", func(t *testing.T) {
		input := []int{2, 3}
		want := []int{4, 9}
		got := solution3(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array4", func(t *testing.T) {
		input := []int{-2, -1, 0}
		want := []int{0, 1, 4}
		got := solution3(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array5", func(t *testing.T) {
		input := []int{-3, -2, -1}
		want := []int{1, 4, 9}
		got := solution3(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("977. Squares of a Sorted Array6", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := []int{1, 4, 9}
		got := solution3(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	使用双指针，先遍历一次找到负数部分的最大值的位置 N，然后负数从 N 向 数组开头遍历，
	正数从 N + 1 向数组后面遍历
	循环终止条件是 i >= N 且 j < len(A)
		内部比较条件是直接用平方来比较
	最后，处理遍历了一边，另外一边没有遍历的情况
	这里需要处理好两个边界以及增长关系
*/
func solution(A []int) []int {
	aLen := len(A)
	var i, j int
	// i 指向第一个非负数
	for i = 0; i < aLen && A[i] < 0; i++ {

	}
	// j 是最后一个负数的开始索引
	j = i - 1
	var res []int
	for j >= 0 && i < aLen {
		left, right := A[i]*A[i], A[j]*A[j]
		if left < right {
			res = append(res, left)
			i++
		} else {
			res = append(res, right)
			j--
		}
	}
	for i < aLen {
		res = append(res, A[i]*A[i])
		i++
	}
	for j >= 0 {
		res = append(res, A[j]*A[j])
		j--
	}
	return res
}

func solution2(A []int) []int {
	// 找到第一个 >=0 的元素以及它的上一个元素，向两边遍历过去
	// 因为原数组是按照升序排列的，中间元素绝对值是最小的
	aLen := len(A)
	// 左右索引
	var leftIndex int
	var rightIndex int
	var retAttr = make([]int, aLen)
	for ; rightIndex < aLen && A[rightIndex] < 0; rightIndex++ {

	}
	leftIndex = rightIndex - 1
	// 遍历数组
	for i := 0; i < aLen; i++ {
		var leftSquare int
		var rightSquare int
		// 先判断边界条件 左右两边是否到边界
		// 左边判断
		if leftIndex >= 0 {
			leftSquare = A[leftIndex] * A[leftIndex]
		} else {
			// 左边已经遍历完了
			rightSquare = A[rightIndex] * A[rightIndex]
			retAttr[i] = rightSquare
			rightIndex++
			continue
		}
		// 右边判断
		if rightIndex < aLen {
			rightSquare = A[rightIndex] * A[rightIndex]
		} else {
			// 右边已经遍历完
			leftSquare := leftSquare
			retAttr[i] = leftSquare
			leftIndex--
			continue
		}
		// 两边都没有遍历完
		if leftSquare < rightSquare {
			retAttr[i] = leftSquare
			leftIndex--
		} else {
			retAttr[i] = rightSquare
			rightIndex++
		}
	}
	return retAttr
}

//func sortedSquares(A []int) []int {
//	var leftIndex int
//	var rightIndex int
//	var retArr []int
//	aLen := len(A)
//	if A[0] >= 0 {
//		// all positive
//		for _, v := range A {
//			retArr = append(retArr, v*v)
//		}
//		return retArr
//	} else if A[aLen-1] <= 0 {
//		// all negative
//		for _, v := range A {
//			retArr = append(retArr, v*v)
//		}
//		return retArr
//	} else {
//		// positive and negative
//		// 找到第一个大于等于 0 的元素以及下一个元素
//		// 所以这里找错了
//		for i, v := range A {
//			if v >= 0 {
//				leftIndex = i
//				if i+1 < aLen-1 {
//					rightIndex = i + 1
//				}
//				break
//			}
//		}
//	}
//	// 比较这里，一个循环就可以了，这里也有问题
//	for leftIndex >= 0 && rightIndex <= aLen-1 {
//		squareLeft := A[leftIndex] * A[leftIndex]
//		squareRight := A[rightIndex] * A[rightIndex]
//		if squareLeft < squareRight {
//			retArr = append(retArr, squareLeft)
//			leftIndex--
//		} else {
//			retArr = append(retArr, squareRight)
//			rightIndex++
//		}
//	}
//	for i := leftIndex; i >= 0; i-- {
//		retArr = append(retArr, A[i]*A[i])
//	}
//	for i := rightIndex; i < aLen; i++ {
//		retArr = append(retArr, A[i]*A[i])
//	}
//	return retArr
//}

// 两边向中间，最后需要翻转
//func solution(A []int) []int {
//	aLen := len(A)
//	leftIndex := 0
//	rightIndex := aLen - 1
//	var retArr []int
//	for leftIndex <= rightIndex {
//		squareLeft := A[leftIndex] * A[leftIndex]
//		squareRight := A[rightIndex] * A[rightIndex]
//		if squareLeft > squareRight {
//			retArr = append(retArr, squareLeft)
//			leftIndex++
//		} else {
//			retArr = append(retArr, squareRight)
//			rightIndex--
//		}
//	}
//	for i := 0; i < aLen/2; i++ {
//		retArr[i], retArr[aLen-i-1] = retArr[aLen-i-1], retArr[i]
//	}
//	return retArr
//}

//func IntSliceEqual(a, b []int) bool {
//	if len(a) != len(b) {
//		return false
//	}
//
//	if (a == nil) != (b == nil) {
//		return false
//	}
//
//	for i, v := range a {
//		if v != b[i] {
//			return false
//		}
//	}
//
//	return true
//}

func solution3(arr []int) []int {
	left, right := 0, len(arr)-1
	res := make([]int, 0)
	for left <= right {
		leftRes := arr[left] * arr[left]
		rightRes := arr[right] * arr[right]
		if leftRes > rightRes {
			res = append(res, leftRes)
			left++
		} else {
			res = append(res, rightRes)
			right--
		}
	}
	left, right = 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}
