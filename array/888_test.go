package array

import (
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("888. Fair Candy Swap", func(t *testing.T) {
		arrA := []int{1, 1}
		arrB := []int{2, 2}
		want := []int{1, 2}
		got := solution(arrA, arrB)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("888. Fair Candy Swap2", func(t *testing.T) {
		arrA := []int{2}
		arrB := []int{1, 3}
		want := []int{2, 3}
		got := solution(arrA, arrB)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("888. Fair Candy Swap3", func(t *testing.T) {
		arrA := []int{1, 2, 5}
		arrB := []int{2, 4}
		want := []int{5, 4}
		got := solution(arrA, arrB)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(arrA, arrB []int) []int {
	retArr := make([]int, 2)
	bMap := make(map[int]int, len(arrB))
	sumA, sumB := calcSum(arrA), calcSum(arrB)
	diff := (sumA - sumB) / 2
	// 使用 map的处理 在 map 中找
	for _, v := range arrB {
		bMap[v]++
	}
	for _, v := range arrA {
		want := v - diff
		if bMap[want] > 0 {
			retArr[0] = v
			retArr[1] = want
			return retArr
		}
	}
	return retArr
}

/*
	a 在 S1 中，b 在 S2 中 所以需要提示
	S1 + b - a = S2 + a - b
	a = (S1 - S2) / 2 + b
*/
// O(n^2)
//func solution(arrA, arrB []int) []int {
//	retArr := make([]int, 2)
//	sumA, sumB := calcSum(arrA), calcSum(arrB)
//	diff := (sumA - sumB) / 2
//	sort.Ints(arrA)
//	sort.Ints(arrB)
//	for _, v := range arrA {
//		want := v - diff
//		for _, w := range arrB {
//			if w == want {
//				retArr[0] = v
//				retArr[1] = w
//				return retArr
//			} else if w < want {
//				continue
//			} else {
//				break
//			}
//		}
//	}
//	return retArr
//}

/*
	先求和，如果两个和相等则直接返回空列表
	较小和的数组升序排序以及较大和的数组降序排序，分别记为A，B
	双指针的方法，交换元素并判断
	左右两个指针交换，即两个数组元素交换的时候，需要判断
		使得A较大，B较小 可以 1：继续看A的下一个元素 2：继续看B的下一个元素
			这里可以A的元素不变，继续遍历B的元素直到A的元素大于等于B的元素
*/
// 交换两个数组的元素使得两者的和相等
// 这种方法太麻烦了
//func solution(arrA, arrB []int) []int {
//	arraSum, arrbSum := calcSum(arrA), calcSum(arrB)
//	var smallerArr, biggerArr []int
//	var smallerSum, biggerSum int
//	retArr := make([]int, 2)
//	if arraSum < arrbSum {
//		smallerArr = arrA
//		smallerSum = arraSum
//		biggerArr = arrB
//		biggerSum = arrbSum
//	} else {
//		smallerArr = arrB
//		smallerSum = arrbSum
//		biggerArr = arrA
//		biggerSum = arraSum
//	}
//	sort.Ints(smallerArr)
//	sort.Ints(biggerArr)
//	lenBigger := len(biggerArr)
//	for i := 0; i < lenBigger/2; i++ {
//		t := lenBigger - 1 - i
//		biggerArr[i], biggerArr[t] = biggerArr[t], biggerArr[i]
//	}
//	for i, j := 0, 0; smallerArr[i] < biggerArr[j]; {
//		ss := smallerSum - smallerArr[i] + biggerArr[j]
//		bs := biggerSum - biggerArr[j] + smallerArr[i]
//		if ss == bs {
//			retArr[0] = smallerArr[i]
//			retArr[1] = biggerArr[j]
//			return retArr
//			// 可以左边不变，然后右边不断增加
//			// 也可以右边不变，左边不断增加
//		} else if ss > bs {
//			for k, l := i, j; smallerArr[k] < biggerArr[l] && k < len(smallerArr) && l < len(biggerArr); {
//				ss2 := smallerSum - smallerArr[k] + biggerArr[l]
//				bs2 := biggerSum - biggerArr[l] + smallerArr[k]
//				if ss2 == bs2 {
//					retArr[0] = smallerArr[k]
//					retArr[1] = biggerArr[l]
//					return retArr
//				} else if ss2 > bs2 {
//					l++
//				} else {
//					break
//				}
//			}
//			i++
//		} else {
//			return retArr
//		}
//	}
//	return retArr
//}

func calcSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
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
