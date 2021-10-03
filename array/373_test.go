package array

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("373. Find K Pairs with Smallest Sums", func(t *testing.T) {
		nums1 := []int{0, 0, 0, 0, 0}
		nums2 := []int{-3, 22, 35, 56, 76}
		k := 22
		want := [][]int{{1, 3}, {2, 3}, {1, 5}}
		got := solution(nums1, nums2, k)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	给定两个升序排序的整型数组 nums1, nums2 以及整型数字 k，定义 pair (u, v)，其中一个元素从第一个数组中获取，另一个元素从第二个数组中获取
	返回 pair (u, v) 和最小的 k 对。
	暴力方法是直接将枚举所有的序对，然后按照总和排序
*/

/*
	二分查找做法，直接从右上角开始判断，无法通过所有测试用例。。。
	这里使用快排 OOM 了
*/

type ele2 [][]int

func (e ele2) Len() int           { return len(e) }
func (e ele2) Less(i, j int) bool { return e[i][0] < e[j][0] }
func (e ele2) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (e *ele2) Push(x interface{}) {
	*e = append(*e, x.([]int))
}

func (e *ele2) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func solution(nums1, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	res := make([][]int, 0)
	if n == 0 || m == 0 || k == 0 {
		return res
	}
	// k 大于总数，直接全部作为结果了
	if k >= n*m {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				res = append(res, []int{nums1[i], nums2[j]})
			}
		}
		return res
	}
	low, high := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]
	for low < high {
		mid := low + (high-low)/2
		res = make([][]int, 0)
		if isEnough(m, n, k, mid, nums1, nums2, &res) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	// 如果两个相等，则最后一次获取需要在外部获取，因为相等了前面的循环会跳过最后一次
	if low == high {
		res = make([][]int, 0)
		isEnough(m, n, k, low, nums1, nums2, &res)
	}
	// 如果元素数量超过了 k 就需要将较大的剔除，有一个快速排序的变种，划分的时候判断一下左右两边的数量
	// 这里会 OOM，那就只能使用堆了
	if len(res) > k {
	}
	return res
}

// 这里从右上角开始
func isEnough(m, n, k, mid int, nums1, nums2 []int, res *[][]int) bool {
	i, j := 0, n-1
	count := 0
	/*
		如果结果数组数量足够了，就不需要继续添加了，这里是有问题的，因为虽然上面的元素可以保证是小于 mid 的
		但是下一行的元素也可能是小于 mid 的，同时下面的元素可能更小，这样达到目标数量就终止，就会返回更大的序对，而忽略了下面的序对了
		因此，这里不能这样计数，需要都加入，然后重新使用优先队列处理排除较大的元素，但是这样感觉不如直接优先队列了？
	*/
	// 这里以行为主来移动，每次加上当前行
	for i < m && j >= 0 {
		// 如果满足就可以直接将当前列的该行以及上面的元素数量统计，然后统计下一列
		if nums1[i]+nums2[j] <= mid {
			// 将元素添加到结果 slice
			for re := 0; re <= j; re++ {
				*res = append(*res, []int{nums1[i], nums2[re]})
			}
			count += j + 1
			i++
			// 如果大于，则需要行数递减来判断更小的元素
		} else {
			j--
		}
	}
	return count >= k
}

func qs(pairs *[][]int, left, right int) {
	if right > left {
		pivotIndex := left
		pivotIndexN := exchange(pairs, left, right, pivotIndex)
		qs(pairs, left, pivotIndexN-1)
		qs(pairs, pivotIndexN+1, right)
	}
}

func exchange(pairs *[][]int, left, right, pivotIndex int) int {
	pivotVal := (*pairs)[pivotIndex][0] + (*pairs)[pivotIndex][1]
	(*pairs)[pivotIndex][0], (*pairs)[right][0] = (*pairs)[right][0], (*pairs)[pivotIndex][0]
	(*pairs)[pivotIndex][1], (*pairs)[right][1] = (*pairs)[right][1], (*pairs)[pivotIndex][1]
	pCount := left
	for i := left; i < right; i++ {
		temp := (*pairs)[i][0] + (*pairs)[i][1]
		if temp < pivotVal {
			(*pairs)[i][0], (*pairs)[pCount][0] = (*pairs)[pCount][0], (*pairs)[i][0]
			(*pairs)[i][1], (*pairs)[pCount][1] = (*pairs)[pCount][1], (*pairs)[i][1]
			pCount++
		}
	}
	(*pairs)[right][0], (*pairs)[pCount][0] = (*pairs)[pCount][0], (*pairs)[right][0]
	(*pairs)[right][1], (*pairs)[pCount][1] = (*pairs)[pCount][1], (*pairs)[right][1]
	return pCount
}

// 优先队列 5% 主要问题是内存访问不友好导致频繁切页
type ele [][3]int

func (e ele) Len() int           { return len(e) }
func (e ele) Less(i, j int) bool { return e[i][0] < e[j][0] }
func (e ele) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (e *ele) Push(x interface{}) {
	*e = append(*e, x.([3]int))
}

func (e *ele) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

/*
	类似 378 的做法，总体上还是先一行（列）然后找下一个可能的最小元素
	[ref](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/discuss/84551/simple-Java-O(KlogK)-solution-with-explanation)
	这里的思路是，一开始将 nums1[i], i : [0, len(nums1)-1], nums2[j] j 最开始是 0 放入优先队列，然后将顶部元素出队列后，放入元素 nums1[i] nums2[j + 1]
	因为出队列的元素中，一开始 nums2[j] 都是相同的数值，因此下一个比较小的和应该是出队列的元素中的 nums1 的元素加上 nums2[j + 1] 去和队列中原来的元素比较
	TODO: 为什么比下面的快，这里是每次都判断下一个需要放入到优先队列的元素对，同时只有一对需要放入，而下面的做法是，将 i + j == k 的都放入了，显然多了不必要的内存访问
*/

func solution2(nums1, nums2 []int, k int) [][]int {
	res := make([][]int, 0)
	n, m := len(nums1), len(nums2)
	if n == 0 || m == 0 || k == 0 {
		return res
	}
	// 一开始将第一列放到优先队列中
	h := &ele{}
	for i := 0; i < n && i < k; i++ {
		heap.Push(h, [3]int{nums1[i] + nums2[0], i, 0})
	}
	for k > 0 && len(*h) > 0 {
		topE := heap.Pop(h).([3]int)
		k--
		res = append(res, []int{nums1[topE[1]], nums2[topE[2]]})
		if topE[2] == m-1 {
			continue
		}
		heap.Push(h, [3]int{nums1[topE[1]] + nums2[topE[2]+1], topE[1], topE[2] + 1})
	}
	return res
}

/*
	优先队列做法
		count = 0，count < k，递增，
		每次将 i + j == count 的元素放入到优先队列中，然后取出头部元素保存到结果，这里队列元素 {sum, i, j} i,j 记录 sum 的索引
*/
func solution3(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	res := make([][]int, 0)
	if k >= n*m {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				res = append(res, []int{nums1[i], nums2[j]})
			}
		}
		return res
	}
	// 优化内存访问
	h := &ele{}
	count := 0
	for i := 0; i < n && i < k; i++ {
		for j := 0; j < m && i+j <= k; j++ {
			heap.Push(h, [3]int{nums1[i] + nums2[j], i, j})
		}
	}
	for count < k {
		topE := heap.Pop(h).([3]int)
		res = append(res, []int{nums1[topE[1]], nums2[topE[2]]})
		count++
	}
	return res
}
