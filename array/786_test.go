package array

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("786. K-th Smallest Prime Fraction", func(t *testing.T) {
		arr := []int{1, 2, 3, 5}
		k := 3
		got := solution(arr, k)
		want := 6
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定已排序数组 arr，包含 1 和 prime 数字，所有元素都是不同的。给定 k，对每个 i j 满足 0 <= i < j < len(arr) 我们得到分数
	arr[i] / arr[j] 返回第 k 个最小的分数，返回的结果放在数组中，其中第一个元素是 arr[i] 第二个是 arr[j]

	有点类似 373？可以使用优先队列，也可以使用二分查找。
	这里是同一个数组里面选择元素组成所求的数字，主要还是访问数组元素造成 cache 频繁切换的问题\

	这里是找第 k 个最小的元素
*/

type ele []*eleStr

type eleStr struct {
	i, j     int
	fraction float64
}

func (e ele) Len() int           { return len(e) }
func (e ele) Less(i, j int) bool { return e[i].fraction < e[j].fraction }
func (e ele) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (e *ele) Push(x interface{}) {
	*e = append(*e, x.(*eleStr))
}

func (e *ele) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func solution(arr []int, k int) []int {
	h := &ele{}
	arrLen := len(arr)
	// 需要将第一行全部放入
	for i := 0; i < arrLen-1; i++ {
		heap.Push(h, &eleStr{0, i + 1, float64(arr[0]) / float64(arr[i+1])})
	}
	for i := 0; i < k-1 && len(*h) > 0; i++ {
		top := heap.Pop(h).(*eleStr)
		// 当前的分母的分子还有后面的元素则将新的元素放入，因为当前的最小的分母不变，下一个较小的元素是当前分子的下一个元素和当前分母构成的
		if top.i < arrLen {
			heap.Push(h, &eleStr{top.i + 1, top.j, float64(arr[top.i]) / float64(arr[top.j])})
		}
	}
	top := heap.Pop(h).(*eleStr)
	return []int{arr[top.i], arr[top.j]}
}
