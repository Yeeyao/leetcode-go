package stack

import (
	"container/heap"
	"math"
	"sort"
)

/*
leetcode++ 介绍是使用 q/w 表示卷王指数，这个数值越大，表示干越多活，拿越少工资。因为要保证 k 个人中每个人至少需要有最低工资，因此这个数值必须是 k 人里面最小的
使用元组 eff[(q / w, q, w)] 保存，然后根据 -q/w 进行排序
因为需要 k 个工人，就从第一个开始到 k 进行计算，每个群体的最后一个的 q/w 就是 k 个里面最小的，计算所有群体的总工资然后取最小值
有点暴力解法的味道

这里提到工人群体的大小固定，因此就有类似固定堆的解题思路

```python
class Solution:
    def mincostToHireWorkers(self, quality: List[int], wage: List[int], K: int) -> float:
        eff = [(q / w, q, w) for a, b in zip(quality, wage)]
        eff.sort(key=lambda a: -a[0])
        ans = float('inf')
		// 因为一定要 k 个人，所以从 K-1 开始，先按照 q/w 降序排列，每次取 k 个中最右边的就是当前的最小的 q/w 然后计算目前的最小工资
        for i in range(K-1, len(eff)):
            h = []
            k = K - 1
			// 按照排序结果，这个就是当前最小的 q/w
            rate, _, total = eff[i]
            # 找出工作效率比它高的 k 个人，这 k 个人的工资尽可能低。
            # 由于已经工作效率倒序排了，因此前面的都是比它高的，然后使用堆就可得到 k 个工资最低的。
            for j in range(i):
				// 这里每次都从头到 j 将工资入堆
                heapq.heappush(h, eff[j][1] / rate)
            while k > 0:
				// 利用堆的特性，每次都将堆中最小的工资累加，然后计算极值
                total += heapq.heappop(h)
                k -= 1
            ans = min(ans, total)
        return ans

固定堆思路，这里使用 effs[(q/w, q)] 两元组，同样 q/w 按照降序排列。total 保存当前的总工作量，然后 h 保存遍历过的每个工人的工作量，
这里负数是方便群体变化的时候将当前的最大的的工作量(利用堆)直接相加进行移除，因为是按照 q/w 降序排列，最右边的元素的 q/w 就是最小的，
直接和总工作量计算得到总工资，移动的时候减去当前的最大的工作量，加上下一个工人的工作量
可以在一个循环里面完成

class Solution:
    def mincostToHireWorkers(self, quality: List[int], wage: List[int], K: int) -> float:
        effs = [(q / w, q) for q, w in zip(quality, wage)]
        effs.sort(key=lambda a: -a[0])
        ans = float('inf')
        h = []
        total = 0
        for rate, q in effs:
			// 负数刚好可以将最大的工资放到堆顶
            heapq.heappush(h, -q)
            total += q
			// 工资数量超过 k 就将最小（工资最大的）移除堆
            if len(h) > K:
                total += heapq.heappop(h)
			// 数量刚好就计算
            if len(h) == K:
                ans = min(ans, total / rate)
        return ans

*/

/*
雇佣 K 名工人的最低成本

There are n workers. You are given two integer arrays quality and wage where quality[i] is the quality of the ith worker and wage[i] is the minimum wage expectation for the ith worker.

We want to hire exactly k workers to form a paid group. To hire a group of k workers, we must pay them according to the following rules:

    Every worker in the paid group should be paid in the ratio of their quality compared to other workers in the paid group.
    Every worker in the paid group must be paid at least their minimum wage expectation.

Given the integer k, return the least amount of money needed to form a paid group satisfying the above conditions. Answers within 10-5 of the actual answer will be accepted.


N 名工人，第 i 个的工作质量是 quality[i]，最低期望工资是 wage[i]

现在我们想雇佣K名工人组成一个工资组。在雇佣一组 K 名工人时，我们必须按照下述规则向他们支付工资：

对工资组中的每名工人，应当按其工作质量与同组其他工人的工作质量的比例来支付工资。
工资组中的每名工人至少应当得到他们的最低期望工资。
返回组成一个满足上述条件的工资组至少需要多少钱。

这里要求按照工作质量与其他人的工作质量比例来支付工资，同时需要每个工人满足最低工资期望。
可以使用 w[i]/q[i] 表示单位质量所需要的工资，因为后一个条件需要满足，因此需要找到 w[i]/q[i] 最高的工人，这里只找出最大值就好了，同时，这里只需要找到
k 个工人里面的 w/q 的最大值，同时需要总工资最小，因此需要 q 尽量小

```

times: 1

*/

// 参考上面第二种方法
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	// 构造然后根据 q/w 进行排序
	ws := wqSlice{Wqs: []*wageQuality{}}
	for i := 0; i < len(quality); i++ {
		ws.Wqs = append(ws.Wqs, &wageQuality{
			Qdw:     float64(quality[i]) / float64(wage[i]),
			Quality: quality[i],
		})
	}
	sort.Sort(&ws)

	var qualityHeap qualitySlice
	ret := math.MaxFloat64
	var total int
	// 入堆然后计算
	for _, v := range ws.Wqs {
		rate, quality := v.Qdw, v.Quality
		heap.Push(&qualityHeap, -quality)
		total += quality
		// 如果数量多于 k 就将最大的工资移除
		if qualityHeap.Len() > k {
			total += heap.Pop(&qualityHeap).(int)
		}
		if qualityHeap.Len() == k {
			if float64(total)/rate < ret {
				ret = float64(total) / rate
			}
		}
	}
	return ret
}

type qualitySlice struct {
	sort.IntSlice
}

type wqSlice struct {
	Wqs []*wageQuality
}

type wageQuality struct {
	Qdw     float64
	Quality int
}

func (ws *wqSlice) Less(i, j int) bool {
	return ws.Wqs[i].Qdw > ws.Wqs[j].Qdw
}

func (ws *wqSlice) Swap(i, j int) {
	temp := ws.Wqs[i]
	ws.Wqs[i] = ws.Wqs[j]
	ws.Wqs[j] = temp
}

func (ws *wqSlice) Len() int {
	return len(ws.Wqs)
}

func (qs *qualitySlice) Push(v interface{}) {
	qs.IntSlice = append(qs.IntSlice, v.(int))
}

func (qs *qualitySlice) Pop() interface{} {
	top := qs.IntSlice[len(qs.IntSlice)-1]
	qs.IntSlice = qs.IntSlice[:len(qs.IntSlice)-1]
	return top
}
