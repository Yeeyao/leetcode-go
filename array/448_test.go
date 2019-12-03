package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("448. Find All Numbers Disappeared in an Array", func(t *testing.T) {
		input := []int{4, 3, 2, 7, 8, 2, 3, 1}
		want := []int{5, 6}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("448. Find All Numbers Disappeared in an Array2", func(t *testing.T) {
		input := []int{3, 3, 2, 2}
		want := []int{1, 4}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("448. Find All Numbers Disappeared in an Array4", func(t *testing.T) {
		input := []int{3, 3, 3, 6, 1, 1}
		want := []int{2, 4, 5}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("448. Find All Numbers Disappeared in an Array5", func(t *testing.T) {
		input := []int{1, 1, 1, 1, 1}
		want := []int{2, 3, 4, 5}
		got := solution(input)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	相等就直接下一个，不相等，那就是等于上一个或者大于上一个
		等于上一个，就直接跳过
		大于上一个，判断大于1
			大于1，需要将中间的元素保存
			等于1，跳过处理

	需要排序，有跨过的元素就补全
	使用一个当前最大保存数值 Count 初始化为1
	遍历已有元素
		如果当前元素大于最大保存数值，表示中间有缺少的数值，直接将中间的补全
			这里如果相差的是1，则直接更新最大保存数值
			这里补全的数量是 Count到当前元素的数值 补全完了需要更新 Count
		如果当前元素等于最大保存数值，跳过处理

	上面的方法太复杂了，直接使用 int 数组来记录元素是否存在
	使用多一个辅助数组
	直接使用原数组来保存信息？
*/
//func solution(input []int) []int {
//	var retArr []int
//	sort.Ints(input)
//	inputLen := len(input)
//	currentNum := 0
//	for i := 0; i < inputLen; i++ {
//		if input[i] > currentNum {
//			for j := currentNum + 1; j < input[i]; j++ {
//				retArr = append(retArr, j)
//			}
//			currentNum = input[i]
//		}
//	}
//	fmt.Println(currentNum)
//	for j := currentNum + 1; j < inputLen+1; j++ {
//		retArr = append(retArr, j)
//	}
//	return retArr
//}

/*
	使用负数来记录是否已经存在 这里仅仅用负数来判断
*/

func solution(input []int) []int {
	var retArr []int
	inputLen := len(input)
	for i := 0; i < inputLen; i++ {
		var pos int
		if input[i] < 0 {
			pos = -input[i] - 1
		} else {
			pos = input[i] - 1
		}
		if input[pos] > 0 {
			input[pos] = -input[pos]
		}
	}
	for i := 0; i < inputLen; i++ {
		if input[i] > 0 {
			retArr = append(retArr, i+1)
		}
	}
	return retArr
}

/*
	使用负数来记录是否已经存在，直接全部改完，但是有多余的判断
*/
func solution(input []int) []int {
	var retArr []int
	inputLen := len(input)
	for i := 0; i < inputLen; i++ {
		pos := input[i] - 1
		// 这里是一改就改完，这样原来的数组信息就不会丢失
		for pos >= 0 {
			temp := input[pos] - 1
			input[pos] = -1
			pos = temp
		}
	}
	for i := 0; i < inputLen; i++ {
		if input[i] >= 0 {
			retArr = append(retArr, i+1)
		}
	}
	return retArr
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
