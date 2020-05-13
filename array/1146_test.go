package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1146. Snapshot Array", func(t *testing.T) {
		input := []int{4, 3, 2, 1, 0}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里有四种操作，然后只需要在每个 snap 之间保存中间操作
	改变了的元素的内容就好了，查找的时候，先去中间的结果中找，找不到
	就直接查找原来的初始化的东西

	思路可以是每个元素的数值有多个，如果有多次 set，
	则对应的数值存储为 {{snap_id, value}, ...}

	下面的思路是 set 的修改都一直保存下来，然后每次 snap 就分配空间保存下来本次的
	所有的元素的数值。保存的一份数据是一直改变的
	如果 snap 操作很多的话，这种方法很浪费空间
*/
// index, value
type MapArray map[int]int

type SnapshotArray struct {
	// snap_id, map
	snapshots map[int]MapArray
	self      MapArray
}

func Constructor(length int) SnapshotArray {
	snapshots := make(map[int]MapArray)
	self := make(MapArray)
	snapshotArray := SnapshotArray{snapshots, self}
	return snapshotArray
}

// get, set 一直都对同一个 mapArray 操作
func (this *SnapshotArray) Set(index int, val int) {
	this.self[index] = val
	return
}

// 调用的时候，将当前的所有的元素的数值保存
func (this *SnapshotArray) Snap() int {
	snapID := len(this.snapshots)
	this.snapshots[snapID] = make(MapArray)
	for k, v := range this.self {
		this.snapshots[snapID][k] = v
	}
	return snapID
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return this.snapshots[snap_id][index]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

/*
	another way
	这种思路是每个元素的数值有多个，如果有多次 set，
	则对应的数值存储为 {{snap_id, value}, ...}
*/

type Item struct {
	snapId, value int
}
type SnapshotArray struct {
	currSnapId int
	arr        [][]Item
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		currSnapId: 0,
		arr:        make([][]Item, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	n := len(this.arr[index])
	if n > 0 {
		last := &this.arr[index][n-1]
		if last.snapId == this.currSnapId {
			last.value = val
			return
		}
	}
	this.arr[index] = append(this.arr[index], Item{this.currSnapId, val})
}

func (this *SnapshotArray) Snap() int {
	this.currSnapId++
	return this.currSnapId - 1
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	a := this.arr[index]
	n := len(a)
	if n == 0 {
		return 0
	}
	if a[0].snapId > snapId {
		return 0
	}
	if a[n-1].snapId <= snapId {
		return a[n-1].value
	}
	lo, hi := 0, len(a)-1
	for hi-lo > 1 {
		mid := (lo + hi) / 2
		if a[mid].snapId <= snapId {
			lo = mid
		} else {
			hi = mid
		}
	}
	return a[lo].value
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
