package array

import (
	"sort"
	"testing"
)

/*
	对 arr2 构造一个 map，其中 key 是 arr2 的元素，value 是 arr2 的排名
	直接对 arr1 的元素进行遍历，不需要对 arr2 进行 map 的构造
	有一个需要注意的是 arr1 的元素不需要排序的，只需要相对顺序满足就行
	从两边遍历，三个索引？最后一个是不存在于 arr2 的，两个是
	使用两个 slices，一个用来保存不存在于 arr2 而存在于 arr1 的元素，需要排序
*/

type Arrays []int

var sortArr2 []int

func TestPro(t *testing.T) {
	t.Run("1122. Relative Sort Array", func(t *testing.T) {
		arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
		arr2 := []int{2, 1, 4, 3, 9, 6}
		want := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}
		got := solution(arr1, arr2)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1122. Relative Sort Array2", func(t *testing.T) {
		arr1 := []int{2, 3, 1, 3, 2, 4, 6, 9, 2}
		arr2 := []int{2, 1, 4, 3, 9, 6}
		want := []int{2, 2, 2, 1, 4, 3, 3, 9, 6}
		got := solution(arr1, arr2)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

//func solution(arr1, arr2 []int) []int {
//	retArr := Arrays(arr1)
//	// 注意这里因为有 1000 的，所以元素数量要是 1001
//	sortArr2 = make([]int, 1001)
//	for i := 0; i < len(arr2); i++ {
//		sortArr2[arr2[i]] = i + 1
//	}
//	sort.Sort(retArr)
//	return retArr
//}
//
//func (a Arrays) Len() int {
//	return len(a)
//}
//
///*
//	两个都存在于 arr2 则需要直接比较arr2的大小
//	其中一个存在其中一个不存在则取存在的一个
//	两个都不存在则直接比较原来的大小
//*/
//func (a Arrays) Less(i, j int) bool {
//	arr2i := sortArr2[a[i]]
//	arr2j := sortArr2[a[j]]
//	if arr2i != 0 {
//		if arr2j != 0 {
//			return arr2i < arr2j
//		} else {
//			return true
//		}
//	} else {
//		if arr2j != 0 {
//			return false
//		} else {
//			return a[i] < a[j]
//		}
//	}
//}
//
//func (a Arrays) Swap(i, j int) {
//	a[i], a[j] = a[j], a[i]
//}

func solution(arr1, arr2 []int) []int {
	var retArr []int
	var notInArr []int
	// 无差别统计 arr1 中元素出现的个数
	eleCounter := make([]int, 1001)
	for _, v := range arr1 {
		eleCounter[v]++
	}
	// 先将出现在 arr2 中的元素存放进去
	for _, v := range arr2 {
		for eleCounter[v] > 0 {
			retArr = append(retArr, v)
			eleCounter[v]--
		}
	}
	// 最后获得其他元素
	for _, v := range arr1 {
		for eleCounter[v] > 0 {
			notInArr = append(notInArr, v)
			eleCounter[v]--
		}
	}
	sort.Ints(notInArr)
	return append(retArr, notInArr...)
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

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// 这里 int slice 就够了吧？
	elems := make(map[int]int)
	// 统计 arr1 中的元素出现次数
	for _, v := range arr1 {
		elems[v]++
	}
	var ret []int
	var amount int
	// 将 arr2 的所有元素先存放进去，按照 arr2 中出现的顺序
	for _, v := range arr2 {
		amount = elems[v]
		for i := 0; i < amount; i++ {
			ret = append(ret, v)
			elems[v]--
		}
	}
	var a1 []int
	// 存放剩下的元素，排序后放到末尾
	for k, v := range elems {
		for i := 0; i < v; i++ {
			a1 = append(a1, k)
		}
	}
	sort.Ints(a1)
	return append(ret, a1...)
}
