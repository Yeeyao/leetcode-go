package array

import "sort"

/*
56. Merge Intervals

判断当前左端点大小和结果的右端点大小
先排序 初始化第一个间隔起始 begin end
先按照开始时间进行排列，然后将第一个区间加入结果数组中，按顺序依次考虑之后的每个区间
	如果当前区间的左端点在结果数组最后一个区间的右端点之后，则不会重合，
		直接将这个区间加入结果数组末尾
	否则，它们重合，使用当前区间的右端点更新最后一个区间的右端点，设置为两者较大的值
*/
func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(a, b int) bool { return intervals[a][0] < intervals[b][0] })
	iLen := len(intervals)
	if iLen == 0 {
		return [][]int{}
	}
	res := [][]int{}
	for _, intv := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < intv[0] {
			res = append(res, intv)
		} else {
			if intv[1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intv[1]
			}
		}
	}
	return res
}
