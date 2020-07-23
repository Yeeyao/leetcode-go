package jianzhi_offer

import (
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("27  字符串的排列", func(t *testing.T) {
		s := "abc"
		get := solution(s)
		want := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
		if !reflect.DeepEqual(get, s) {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	使用回溯法 重复情况需要剪枝
	第 0 个位置存放所有的元素一次，然后递归下去第 1 个位置存放剩下的元素。。。
	这里就是对于每个位置，需要将每个元素都固定到，然后递归直到最后的位置
	注意这里使用 string 和 []byte 的相互转换
*/
func solution(s string) []string {
	var strSlice []string
	sLen := len(s)
	temp := []byte(s)
	solutionHelper(&strSlice, sLen, 0, s, temp)
	return strSlice
}

func solutionHelper(strSlice *[]string, sLen, beginIndex int, s string, temp []byte) {
	// 已经全部位置都确定好了
	if beginIndex == sLen {
		*strSlice = append(*strSlice, string(temp))
	}
	// 当前的位置需要排除重复的元素
	// 需要注意，后续遍历从 beginIndex 开始，所以，需要通过交换将前面的元素向后存放才能再次遍历到
	appearMap := make(map[int]bool)
	for i := beginIndex; i < sLen; i++ {
		if _, ok := appearMap[i]; ok {
			continue
		}
		appearMap[i] = true
		temp[beginIndex], temp[i] = temp[i], temp[beginIndex]
		solutionHelper(strSlice, sLen, beginIndex+1, s, temp)
		temp[i], temp[beginIndex] = temp[beginIndex], temp[i]
	}
}

/*
	字符串的排列
	使用递归，每个元素都作为开头，然后如果到了结尾，字符数量还不够，就返回头部来追加
	这里的问题是超过最后位置的时候处理麻烦
	这里需要想办法处理，毕竟转换应该相当耗时
*/
//func solution(s string) []string {
//	var strSlice []string
//	sLen := len(s)
//	var temp string
//	solutionHelper(&strSlice, sLen, 0, s, temp)
//	return strSlice
//}

/*
		0		1		2
      1   2  0    2   0   1
      2   1  2    0   1   0
*/
//func solutionHelper(strSlice *[]string, sLen, beginIndex int, s, temp string) {
//	if len(temp) == sLen {
//		*strSlice = append(*strSlice, temp)
//		return
//	}
//	// 如果到了结尾就需要重新到开头开始
//	if beginIndex == sLen - 1 {
//		beginIndex = 0
//	}
//	tempBefore := temp
//	for i := beginIndex; i < sLen; i++{
//		temp = temp + string(s[i])
//		solutionHelper(strSlice, sLen, i + 1, s, temp)
//		temp = tempBefore
//	}
//}
