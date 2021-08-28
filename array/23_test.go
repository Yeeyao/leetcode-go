package array

import (
	"container/heap"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("23. 合并K个升序链表", func(t *testing.T) {
		input := []*ListNode{}
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
	给定一个链表数组，每个链表都已经按照元素大小升序排列，将所有的链表合并到一个链表上

	方法 1：直接顺序合并，使用一个结果链表保存最终的结果。假设输入的链表平均长度是 n，第一次合并需要的时间是 O(1)，结果链表长度是 n，第二次合并，需要时间 O(2n),
           结果链表长度是 2n，第 i 次合并后，结果链表长度是 i)n，所需时间 O((i-1)n+n)=O(in)，因此总时间 O(k^2n) k 是输入链表的数量
	方法 2：分治合并，直接将一对对链表合并，第一次排序，剩余链表数量是 k/2，第二次 k/4 ... 最后是 1。第一次需要时间是 O(kn)，第二次需要 k/2 * O(2kn)
			总时间 O(kn * logk)，使用递归，空间复杂度 O(logk)
	方法 3：优先队列，这个方法和前两种方法的思路有所不同，我们需要维护当前每个链表没有被合并的元素的最前面一个，k 个链表就最多有 k 个满足这样条件的元素，
			每次在这些元素里面选取 val 属性最小的元素合并到答案中。在选取最小元素的时候，我们可以用优先队列来优化这个过程。就是很直观，每次从所有的链表头获取
		    最小的一个。优先队列中元素不超过 k 个，插入删除时间为 O(logk)，最多 kn 个元素，每个元素被插入和删除一次，总时间 O(kn * logk)，优先队列使用的
			空间 O(logk)
			实现上，使用一个结构体保存当前的节点的数值以及当前的节点，然后一开始先将所有的链表的头部元素放到优先队列。
			构造两个指针 head, tail 来保存结果的链表，其中 head 作为结果返回，tail 移动。当优先队列非空的时候进行循环
				获取队列头部元素并弹出队列，将 tail.Next 指向当前的头部元素，然后 tail 向后移动，接着如果头部元素的下一个元素非空，
				则将下一个元素加入到优先队列（表示当前的这个输入链表后面还有元素）
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	分割函数不断分割，然后合并函数每次将分割的进行合并
*/
func solution(lists []*ListNode) *ListNode {
	return splitFunc(lists, 0, len(lists)-1)
}

// 切割函数
func splitFunc(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	return mergeFunc(splitFunc(lists, left, mid), splitFunc(lists, mid+1, right))
}

/*
	head 指向结果链表的头部，tail 指向插入位置的前一个位置
	其实这里 head 只是记录一下头部，实际的移动是 tail 处理的，因此这里 head 初始化后，tail 需要指向 head
*/
func mergeFunc(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	head := &ListNode{}
	tail := head
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a == nil {
		tail.Next = b
	} else {
		tail.Next = a
	}
	return head.Next
}

// 优先队列方法 需要注意，这里使用 heap 的库函数，则 Pop Push 以及 Init 都需要调用 heap 提供的
type ListNodePq struct {
	val  int
	node *ListNode
}

type ListNodeHeap []ListNodePq

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(ListNodePq))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	nodeHeap := &ListNodeHeap{}
	for _, v := range lists {
		if v != nil {
			nodeHeap.Push(ListNodePq{val: v.Val, node: v})
		}
	}
	heap.Init(nodeHeap)
	head := &ListNode{}
	tail := head
	for nodeHeap.Len() > 0 {
		t := heap.Pop(nodeHeap).(ListNodePq)
		tail.Next = t.node
		tail = tail.Next
		if t.node.Next != nil {
			heap.Push(nodeHeap, ListNodePq{
				val:  t.node.Next.Val,
				node: t.node.Next,
			})
		}
	}
	return head.Next
}
