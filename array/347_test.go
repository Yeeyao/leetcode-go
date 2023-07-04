package array

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

/*
	Given a non-empty array of integers, return the k most frequent elements.
	347. Top K Frequent Elements

	给定一个非空整形数组，返回出现频率最高的前 k 个元素
	算法至少是 O(nlogn) 时间复杂度，k 范围是 1 到 n

	brute force 直接统计每个元素的出现频率然后排序
	一点改进，同样是需要统计频率，只是排序变成了类似快排那种将元素按照频率进行划分
	另一种类似的是统计频率然后使用大顶堆，pop 调前 k 个
*/

func TestPro(t *testing.T) {
	t.Run("347. Top K Frequent Elements", func(t *testing.T) {
		input := []int{1, 1, 1, 2, 2, 3}
		k := 2
		want := []int{1, 2}
		got := topKFrequent(input, k)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	[ref](https://leetcode.com/problems/top-k-frequent-elements/discuss/430247/Go-golang-clean-solution)
	因为输入的元素是不确定的，所以不能直接用数组来初始化
	使用结构体 slice 来保存，然后对 slice 进行排序
	先用 map 来统计出现次数，然后将 map 转换为 slice 之后对 slice 进行排序降序后取前 k 个元素
*/
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	freqMap := make(map[int]int, 0)
	for _, n := range nums {
		freqMap[n]++
	}
	freqSlice := make([][]int, 0)
	for i, f := range freqMap {
		freqSlice = append(freqSlice, []int{i, f})
	}
	// 这里降序排列
	sort.Slice(freqSlice, func(a, b int) bool { return freqSlice[a][1] > freqSlice[b][1] })
	for i := 0; i < k; i++ {
		res = append(res, freqSlice[i][0])
	}
	return res
}

/*
	使用堆的方法
*/
func topKFrequentHeap(nums []int, k int) []int {
	freqMap := make(map[int]int, 0)
	for _, v := range nums {
		freqMap[v]++
	}

	h := eleWithCountSp{epList: []*eleWithCount{}}
	for ele, count := range freqMap {
		heap.Push(&h, &eleWithCount{
			value: ele,
			count: count,
		})
	}
	var res []int
	for ; k > 0; k-- {
		topEwc := heap.Pop(&h).(*eleWithCount)
		res = append(res, topEwc.value)
	}
	return res
}

type eleWithCountSp struct {
	epList []*eleWithCount
}

type eleWithCount struct {
	value int
	count int
}

// 这里将频率大的放在前面
func (sp *eleWithCountSp) Less(i, j int) bool {
	return sp.epList[i].count > sp.epList[j].count
}

func (sp *eleWithCountSp) Swap(i, j int) {
	temp := sp.epList[j]
	sp.epList[j] = sp.epList[i]
	sp.epList[i] = temp
}

func (sp *eleWithCountSp) Len() int {
	return len(sp.epList)
}

func (sp *eleWithCountSp) Push(v interface{}) {
	sp.epList = append(sp.epList, v.(*eleWithCount))
}

func (sp *eleWithCountSp) Pop() interface{} {
	top := sp.epList[len(sp.epList)-1]
	sp.epList = sp.epList[:len(sp.epList)-1]
	return top
}
