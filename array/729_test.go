package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("729. My Calendar I", func(t *testing.T) {
		customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
		grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
		X := 3
		want := 16
		got := solution(customers, grumpy, X)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	book(start, end) 将会产生 [start, end)
	可以添加事件，要求事件不会导致双重的 booking，
	即两个事件的时间区间不会相交
	可以直接使用二维数组保存每个插入的区间，然后每次插入新的区间
	遍历已经插入的区间来判断是否可以插入
*/
type MyCalendar struct {
	intervals [][]int
}

func Constructor() MyCalendar {
	inter := make([][]int, 0)
	return MyCalendar{inter}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, v := range this.intervals {
		// 这里比较开始的较大者和结束的较小者
		// max(start, v[0]) < min(end, v[1]) return false
		a, b := 0, 0
		if start > v[0] {
			a = start
		} else {
			a = v[0]
		}
		if end < v[1] {
			b = end
		} else {
			b = v[1]
		}
		if a < b {
			return false
		}
	}
	this.intervals = append(this.intervals, []int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

/*
	改进的方法 use binary search tree
*/
type MyCalendar struct {
	tree *Tree
}

func Constructor() MyCalendar {
	return MyCalendar{new(Tree)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	return this.tree.Insert(start, end)
}

type Node struct {
	Start int
	End   int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(start, end int) bool {
	inNode := &Node{start, end, nil, nil}
	// 根节点
	if t.root == nil {
		t.root = inNode
		return true
	}
	for node := t.root; node != nil; {
		// 需要向左边找
		if node.Start >= end {
			// 左子树是空的，插入
			if node.Left == nil {
				node.Left = inNode
				return true
			}
			// 左子树非空，继续向左边找
			node = node.Left
			continue
		}
		// 需要向右边找
		if node.End <= start {
			if node.Right == nil {
				node.Right = inNode
				return true
			}
			node = node.Right
			continue
		}
		// 两边都不满足表示有区间重叠了
		return false
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
