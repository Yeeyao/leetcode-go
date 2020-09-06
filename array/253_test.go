package array

import "sort"

/*
253 Meeting Rooms II 类似 56. Merge Intervals
	Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
	find the minimum number of conference rooms required.
	给定开始和结束时间对的数组数组，找到需要的最小会议室数量

Example 1:
Input: [[0, 30],[5, 10],[15, 20]]
Output: 2
Example 2:
Input: [[7,10],[2,4]]
Output: 1
*/

/*
	将开始时间和结束时间分别保存到两个 slice 中并分别排序
	使用 ptr 作为结束时间数组的下标
	循环遍历开始时间数组，如果当前开始时间小于等于结束时间，就需要房间 + 1 否则，结束时间下标 + 1
*/
func minMeetingRooms2(interval [][]int) int {
	iLen := len(interval)
	if iLen == 0 {
		return 0
	}
	startTime := make([]int, iLen)
	endTime := make([]int, iLen)
	for i, inv := range interval {
		startTime[i] = inv[0]
		endTime[i] = inv[1]
	}
	sort.Ints(startTime)
	sort.Ints(endTime)
	count, endPtr := 0, 0
	for _, s := range startTime {
		if s < endTime[endPtr] {
			count++
		} else {
			endPtr++
		}
	}
	return count
}

/*
	先按照开始时间排序，然后用一个栈维护所有的结束时间
	栈中的元素是从栈顶到栈底递减
	每个元素开始时间需要和栈顶元素比较
		如果栈为空或者开始时间小于栈顶元素的结束时间，会议室 + 1
*/
func minMeetingRooms(interval [][]int) int {
	sort.Slice(interval, func(i, j int) bool { return interval[i][0] < interval[j][0] })
	stack := make([]int, 0)
	count := 0
	for _, inv := range interval {
		stackLen := len(stack)
		if stackLen == 0 || inv[1] < stack[stackLen-1] {
			count++
			continue
		} else {
			for len(stack) > 0 && inv[1] > stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, inv[1])
		}
	}
	return count
}
