package binary

import "sort"
import "math/rand" // 默认导入的 rand 不是这个库，需要显式指明

/*
给定一个整型数组 nums 以及 lower, upper 两个整数，返回位于 [lower, upper] 闭区间内的返回和的数量。
范围和 S(i, j) 定义为 nums 的坐标[i,j] 闭区间的数值的和，其中 i <= j
暴力的解法就是遍历整型数组，对每个元素都作为一个开始的 i，向后遍历求范围和然后计算总数，时间复杂度是 O(N^2)
前缀和 pre[i] 表示从开始 0 到 i 下标的数组的和。范围和可以使用前缀和计算得到
S(i, j) = pre[j] - pre[i-1]
这里 S(i, j) = pre[j] - pre[i-1] 的成立条件是前缀和需要是单调递增的(比如 [0,-3,1,4] pre[1] = -3, pre[2] = -2, pre[2]-pre[0]=1)，
这里的元素可能是负数，前缀和不一定是单调递增的
这里可以手动维护为单调递增，但是会丢失索引信息，因此适用于不需要知道具体子序列的情况

比如当前的前缀和是 cur，那么前缀和小于等于 cur - lower 有多少个，就说明以当前结 尾的区间和大于等于 lower 的有多少个。
类似地，前缀和小于等于 cur - upper 有多少个 ，就说明以当前结尾的区间和大于等于 upper 的有多少个。
这里的意思基于 S(i, j) = pre[j] - pre[i-1]，如果前缀和小于等于 cur - lower 则 cur - 前缀和得到的范围和大于等于 lower，
直接使用前者-后者就是答案

基于这个想法，我们可使用二分在 log(N) 的时间快速求出这两个数字，使用平衡二叉树代替数组可使得插入的时间复杂度降低到 O(logN)。
Python 可使用 SortedList 来实现， Java 可用 TreeMap 代替。

时间复杂度和空间复杂度都是 O(logN)

TODO: 需要重做

*/
func countRangeSum(nums []int, lower int, upper int) int {
	var ans, cur int
	var pre []int
	for _, v := range nums {
		// 当前的前缀和
		cur += v
		sort.Ints(pre)
		// >= lower - >= upper 就是两者之间的，同时增加一个元素之后，前面的区间其实都会增加
		// 最右满足 <= target 的位置的右邻居 寻找最左满足 >= target 的位置
		ans += bsr(pre, cur-lower) - bsl(pre, cur-upper)
		// 加入前缀和
		pre = append(pre, cur)
	}
	return ans
}

func bsl(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func bsr(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

/*
	官方的解法
*/

type node struct {
	ch       [2]*node
	priority int
	key      int
	dupCnt   int
	sz       int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) maintain() {
	o.sz = o.dupCnt + o.ch[0].size() + o.ch[1].size()
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap struct {
	root *node
}

func (t *treap) _insert(o *node, key int) *node {
	if o == nil {
		return &node{priority: rand.Int(), key: key, dupCnt: 1, sz: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.ch[d] = t._insert(o.ch[d], key)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.dupCnt++
	}
	o.maintain()
	return o
}

func (t *treap) insert(key int) {
	t.root = t._insert(t.root, key)
}

// equal=false: 小于 key 的元素个数
// equal=true: 小于或等于 key 的元素个数
func (t *treap) rank(key int, equal bool) (cnt int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.ch[0]
		case c > 0:
			cnt += o.dupCnt + o.ch[0].size()
			o = o.ch[1]
		default:
			cnt += o.ch[0].size()
			if equal {
				cnt += o.dupCnt
			}
			return
		}
	}
	return
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	t := &treap{}
	for _, sum := range preSum {
		left, right := sum-upper, sum-lower
		cnt += t.rank(right, true) - t.rank(left, false)
		t.insert(sum)
	}
	return
}
