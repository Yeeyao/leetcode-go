/*
Given an array arr.  You can choose a set of integers and remove all the occurrences of these integers in the array.

Return the minimum size of the set so that at least half of the integers of the array are removed.
*/

// 给定一个整型数组，返回一个唯一整型数字的数量使得移除这些数字后，整个
// 数组的大小至少为原来的一半。
// 同时要求该数字尽可能小
package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {

	t.Run("1338. Reduce Array Size to The Half", func(t *testing.T) {
		input := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1338. Reduce Array Size to The Half2", func(t *testing.T) {
		input := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1338. Reduce Array Size to The Half3", func(t *testing.T) {
		input := []int{7, 7, 7, 7, 7, 7}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1338. Reduce Array Size to The Half4", func(t *testing.T) {
		input := []int{1, 9}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1338. Reduce Array Size to The Half5", func(t *testing.T) {
		input := []int{1000, 1000, 3, 7}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1338. Reduce Array Size to The Half6", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 5
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// 这里统计完了，直接将所有元素保存到另外一个 slice 然后，只需要将元素的数量
// 进行统计，之后，返回元素的数量就行，不需要关心实际的元素
// 这里的思路是，利用 map 来统计，然后转成 slice 进行排序
func solution(arr []int) int {
	arrLen := len(arr)
	sum := 0
	m := make(map[int]int)
	var counter []int

	// 统计出现的次数
	for _, v := range arr {
		m[v]++
	}

	// 这里使用 int slice 来排序
	for _, v := range m {
		counter = append(counter, v)
	}

	// 出现次数的降序排列
	// 这里不反序，下面的就要从尾部向前遍历 不反序更快一点
	sort.Sort(sort.Reverse(sort.IntSlice(counter)))

	for index, v := range counter {
		sum += v
		if sum*2 >= arrLen {
			return index + 1
		}
	}
	return -1
}

/*
	deprecated 啰嗦
*/

// type Pair struct {
// 	Key   int
// 	Value int
// }

// type PairList []Pair

// func (p PairList) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

// func (p PairList) Len() int {
// 	return len(p)
// }

// // 这里升序排列
// func (p PairList) Less(i, j int) bool {
// 	return p[i].Value > p[j].Value
// }

// // 初步思路是将所有出现的元素进行出现次数的统计
// // 之后按照出现次数降序排列，然后一个计数器存放当前已经被剔除的数量，
// // 每次剔除下一个数字之前判断是否已经足够一半了，足够则直接返回
// // 这里想直接用 struct 但是不知道怎么查找什么的。。。
// // TODO: 怎么去找到 struct slice 里面的元素
// func solution(arr []int) int {
// 	num, sum := 0, 0
// 	arrLen := len(arr)
// 	counterMap := make(map[int]int)
// 	for _, v := range arr {
// 		counterMap[v]++
// 	}
// 	p := make(PairList, len(counterMap))
// 	i := 0
// 	for k, v := range counterMap {
// 		p[i] = Pair{k, v}
// 		i++
// 	}
// 	sort.Sort(p)
// 	for _, v := range p {
// 		if arrLen > sum*2 {
// 			AddValue := v.Value
// 			sum += AddValue
// 			num++
// 		} else {
// 			return num
// 		}
// 	}
// 	return num
// }
