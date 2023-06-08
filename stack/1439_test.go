package stack

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("leetcode 1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows", func(t *testing.T) {
		mat := [][]int{{1, 3, 11}, {2, 4, 6}}
		k := 9
		got := kthSmallest(mat, k)
		want := 17
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

/*
给你一个 m n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。
你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小 数组和。

按照这里多路归并的思路，应该是每一行维护一个指向当前最小值的指针
一开始全部都指向第一个元素，得到的是第一个最小的数组和，然后将所有的指针的下一个进行比较，其中最小的就是需要移动的行的指针，作为第二个最小的数组和
但是，如果这里有多个指针的下一个是相同的数值，则会出现多种可能的结果，每个不同的指针表示一种可能的结果(因为会影响后续的指针移动)，不同的结果得到的数组和不一定相同。
这里提到每次分裂之后，极值发生变化，因此是一个动态求极值的题目，可以使用堆

堆配合元组的多路归并方法

TODO: out of range 和错误
*/

func kthSmallest(mat [][]int, k int) int {
	// m row n col
	m := len(mat)
	n := len(mat[0])

	// 先将当前最小的保存
	sumPos := &sumPosStruct{
		spList: make([]*sumAndPositionList, 0),
	}
	minSum := 0
	for _, v := range mat {
		minSum += v[0]
	}
	minPosList := make([]int, m)
	heap.Push(sumPos, &sumAndPositionList{
		sum:          minSum,
		positionList: minPosList,
	})
	// 这里是记录所有行的列指针的数组，如果已经出现过表示是同一种情况，不需要重复计算
	seenMap := make(map[string]struct{})
	// 循环 k 次
	minSp := &sumAndPositionList{}
	for i := 0; i < k; i++ {
		// 每次弹出当前最小的
		minSp = heap.Pop(sumPos).(*sumAndPositionList)
		minSpSum := minSp.sum
		misSpPositionList := minSp.positionList
		// 这里每一行的列索引都需要移动然后将本行的列索引移动后的结果入堆
		for row, pos := range misSpPositionList {
			// 只要没有到最后就每行的指针都移动一个位置
			if pos < n-1 {
				misPosListTemp := make([]int, m)
				copy(misPosListTemp, misSpPositionList)
				misPosListTemp[row]++
				addSum := minSpSum - mat[row][misPosListTemp[row]-1] + mat[row][misPosListTemp[row]]
				if _, ok := seenMap[getStrFromPositionList(misPosListTemp)]; !ok {
					seenMap[getStrFromPositionList(misPosListTemp)] = struct{}{}
					addSapl := &sumAndPositionList{
						sum:          addSum,
						positionList: misPosListTemp,
					}
					heap.Push(sumPos, addSapl)
				}
			}
		}
	}
	return minSp.sum
}

func getStrFromPositionList(intSlice []int) string {
	var str string
	for _, v := range intSlice {
		str += fmt.Sprintf("%v,", v)
	}
	return str
}

type sumPosStruct struct {
	spList []*sumAndPositionList
}

type sumAndPositionList struct {
	sum          int   // 当前的和
	positionList []int // 当前所有行对应的列的下标
}

func (sp *sumPosStruct) Less(i, j int) bool {
	return sp.spList[i].sum < sp.spList[j].sum
}

func (sp *sumPosStruct) Swap(i, j int) {
	temp := sp.spList[i]
	sp.spList[i] = sp.spList[j]
	sp.spList[j] = temp
}

func (sp *sumPosStruct) Len() int {
	return len(sp.spList)
}

func (sp *sumPosStruct) Push(v interface{}) {
	sp.spList = append(sp.spList, v.(*sumAndPositionList))
}

func (sp *sumPosStruct) Pop() interface{} {
	top := sp.spList[len(sp.spList)-1]
	sp.spList = sp.spList[:len(sp.spList)-1]
	return top
}
