package stack

import "sort"

/*
给定下雨的数组 rains, rains[i] 表示今天第 i 个湖泊下雨情况，i 大于 0 表示第 i 个湖泊今天下雨。
i = 0 表示今天湖泊没有下雨，可以选择一个湖泊抽空
需要返回 ans 数组，如果 rains[i] > 0 则 ans[i] = -1 否则 ans[i] 等于抽空的湖泊编号
下雨时湖泊会被充满，此时继续在这个湖泊下雨则会发生洪水，此时直接返回空数组

这里存在的情况就是下雨之后，需要看选择抽空哪个湖泊

这里大概是判断后面下雨的湖泊是否已经正在前面出现过，如果已经出现过则抽水的时候优先选择这些湖泊。因为第一次遍历的时候是没有办法提前知道这个信息的
因此应该是第一次遍历得到上面的信息，然后使用上面的信息来确定抽空哪个湖泊。这里的信息需要保存什么？
如果同一个湖泊连续下雨（这里不需要相邻，指的是两次下雨之间没有可以抽空的时候，当然，这个信息不确定是第一次还是第二次遍历的时候判断）则直接返回空数组

- 这里需要抽空的是之前已经下雨的湖泊
- 如果多个湖泊都满了，则选择这些湖泊中后面最早下雨的一个（这个信息怎么记录，直接使用日期（下标），湖泊编号）
- 有湖泊下雨就将后面存在这个湖泊下雨的信息放到堆里，应该使用多个堆，每个堆保存每个湖泊的下雨情况

题解

	遍历 rains 数组，模拟每天变化
	如果 rains 0 表示是晴天，不抽干任何湖泊，将当天记录到 sunny 数组
	如果 rains 大于 0 表示有一个湖泊下雨，看它是否已经洪水泛滥（之前是否下雨）。使用 lake 记录每个湖泊情况，0 表示没水 1 表示有水，lake[i] = 1 表示已经下雨
	如果当前湖泊下雨了，则到 sunny 数组找一个晴天抽干它，只需要保持 lakes[i] = 1 就可以了

	这里的力扣加加题解有问题，0,1,1 例子，第一个晴天不能被后面使用的，本质问题是第二次降水的时候只能用两次之间的晴天来抽干第一次的降水
	因此这里湖泊记录下上次的降水的日期，然后需要抽空的时候找这个日期之后的晴天才行

*/

func avoidFlood(rains []int) []int {
	lakeRainMap := make(map[int]int)
	// 记录晴天的天数
	var ans, sunny []int
	for i, lakeNum := range rains {
		if lakeNum > 0 {
			// 这里可以第 0 天下雨，所以这个判断有问题
			//lakeRainDay := lakeRainMap[lakeNum]
			//if lakeRainDay > 0 {
			// 当前的湖泊已经有水了，找之前的晴天中可以抽水的一天
			if lakeRainDay, ok := lakeRainMap[lakeNum]; ok {
				sunnyLen := len(sunny)
				findIndex := sort.SearchInts(sunny, lakeRainDay)
				// 找不到下一天
				if findIndex == sunnyLen {
					return []int{}
				}
				// 这一个晴天用掉了
				ans[sunny[findIndex]] = lakeNum
				copy(sunny[findIndex:sunnyLen-1], sunny[findIndex+1:sunnyLen])
				sunny = sunny[:sunnyLen-1]
			}
			// 将湖泊记录为水满，这里记录天数
			lakeRainMap[lakeNum] = i
			ans = append(ans, -1)
		} else {
			// 晴天默认使用 1 号
			ans = append(ans, 1)
			sunny = append(sunny, i)
		}
	}
	return ans
}

func getNextDay(day int, sunny []int) int {
	left, right := 0, len(sunny)-1
	for left < right {
		mid := left + (right-left)/2
		if day > sunny[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
