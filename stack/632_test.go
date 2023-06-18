package stack

import (
	"container/heap"
)

/*
632. Smallest Range Covering Elements from K Lists
你有 k 个 非递减排列 的整数列表。找到一个 最小 区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
找交叉？需要注意这些数组元素是非递减的。

这道题本质上就是在 m 个一维数组中各取出一个数字，重新组成新的数组 A，使得新的数组 A 中最大值和最小值的差值（diff）最小。

这里的多路在于多个输入的数组，关键需要找到堆保存什么信息，当前的最小区间吧。然后每个数组同样是移动指针
这里同样是记录下当前的最小 diff，然后每个数组一个指针来移动，类似 1439，每一次所有数组的指针位置结果和当前的最小值保存到堆，使用另外一个变量保存当前堆的最大值
更新的时候选择其中一个数组指针前进

使用小顶堆获取堆中最小值，进而通过一个变量记录堆中的最大值，这样就知道了 diff，每次更新指针都会产生一个新的 diff，不断重复这个过程并维护全局最小 diff 即可。
这种算法的成立的前提是 k 个列表都是升序排列的，这里需要数组升序原理和上面题目是一样的，有序之后就可以对每个列表维护一个指针，进而使用上面的思路解决。

选择最小值移动是为了让区间的开始数值变大，让区间变小，当前的最大值则是为了满足题目要求让所有列表至少有元素在区间内
这里如果有一行的列指针达到结尾就表示这一行当前的最后的元素已经是区间的最小值了，可以提前返回
这里的最大值的计算是只需要计算当前遇到的元素的最大值，因为后面的会更大。而题目要求 diff（区间最大值-区间最小值）尽可能小
这里堆保存当前的元素数值以及对应的行列序号，对最小值所在的行的列指针进行移动，加入堆后更新堆当前的最大值

times: 1
*/

func smallestRange(nums [][]int) []int {
	h := eleWithPosSp{epList: []*eleWithPos{}}
	// 返回的最大值
	max := 100000
	min := -100000

	nowMax := min
	for row, v := range nums {
		heap.Push(&h, &eleWithPos{
			value: v[0],
			row:   row,
			col:   0,
		})
		if v[0] > nowMax {
			nowMax = v[0]
		}
	}
	for {
		minEle := heap.Pop(&h).(*eleWithPos)
		// 判断 diff 并更新区间
		if nowMax-minEle.value < max-min {
			min, max = minEle.value, nowMax
		}
		// 如果本行已经到达最后一列，则直接返回
		if minEle.col == len(nums[minEle.row])-1 {
			return []int{min, max}
		}
		// 否则，本行指针移动到下一列
		addValue := nums[minEle.row][minEle.col+1]
		heap.Push(&h, &eleWithPos{
			value: addValue,
			row:   minEle.row,
			col:   minEle.col + 1,
		})
		// 因为所有行都会前进，所以这里需要更新最大值，但是最终还是比较 diff
		if addValue > nowMax {
			nowMax = addValue
		}
	}
}

type eleWithPosSp struct {
	epList []*eleWithPos
}

type eleWithPos struct {
	value int
	row   int
	col   int
}

func (sp *eleWithPosSp) Less(i, j int) bool {
	return sp.epList[i].value < sp.epList[j].value
}

func (sp *eleWithPosSp) Swap(i, j int) {
	temp := sp.epList[j]
	sp.epList[j] = sp.epList[i]
	sp.epList[i] = temp
}

func (sp *eleWithPosSp) Len() int {
	return len(sp.epList)
}

func (sp *eleWithPosSp) Push(v interface{}) {
	sp.epList = append(sp.epList, v.(*eleWithPos))
}

func (sp *eleWithPosSp) Pop() interface{} {
	top := sp.epList[len(sp.epList)-1]
	sp.epList = sp.epList[:len(sp.epList)-1]
	return top
}
