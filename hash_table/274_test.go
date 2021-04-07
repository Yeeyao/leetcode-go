package hash_table

import "testing"

/*
[ref](https://leetcode.com/problems/h-index/)

Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith
paper, return compute the researcher's h-index.
According to the definition of h-index on Wikipedia: A scientist has an index h if h of their n papers have at least h
citations each, and the other n − h papers have no more than h citations each.
If there are several possible values for h, the maximum one is taken as the h-index

这里 h 指的是计算研究者的 h 指数，它是一个科学家的 n 篇 paper 的每个都至少有 h 次引用，其他 paper 少于 h 次。如果存在多个可能的 h 值
则取最大的
*/

func TestPro(t *testing.T) {
	t.Run("274. H-Index", func(t *testing.T) {
		citations := []int{3, 0, 6, 1, 5}
		got := solution(citations)
		want := 3
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("274. H-Index2", func(t *testing.T) {
		citations := []int{1, 3, 1}
		got := solution(citations)
		want := 1
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("274. H-Index3", func(t *testing.T) {
		citations := []int{}
		got := solution(citations)
		want := 0
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("274. H-Index4", func(t *testing.T) {
		citations := []int{4}
		got := solution(citations)
		want := 1
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	简单的思路是先将每个论文的引用次数降序排序，然后用一个计数器 h 表示当前的最大的 h，初始化是 1
	遍历排序后数组然后判断
		如果当前数值大于等于 h，则 h 递增，遍历下一个
		如果当前数值小于 h 则直接返回 h
	但是这里的标签是 哈希表，排序后判断是不是效率太低了，是的

这里的问题是，最大值将会被引用次数以及该次数的引用文章数量限制

基于计数排序的解法
其实是利用计数数组，然后从最大的反向遍历判断，这里有当前的 h 和当前数值的关系
这里遍历每个数值和计数然后判断
	如果当前 h 大于当前数值，则直接返回 h 了
	如果当前 h 小于当前数值，则需要 h 增加，但是这里最多可以增加到 h 和数值的最大值
		如果 h + 当前数值数量 > val 则返回 val
		如果 h + 当前数值数量 = val 则返回 val
        如果 h + 当前数量 < val h 递增计数 继续遍历
	如果当前 h 等于当前数值，直接返回 h 了

*/
func solution(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	minNum, maxNum := citations[0], citations[0]
	for _, v := range citations {
		if v < minNum {
			minNum = v
		}
		if v > maxNum {
			maxNum = v
		}
	}
	caLen := maxNum - minNum + 1
	countArr := make([]int, caLen)
	for _, v := range citations {
		countArr[v-minNum]++
	}
	h := 0
	for i := caLen - 1; i >= 0; i-- {
		val := minNum + i
		if h >= val {
			return h
		} else {
			count := countArr[i]
			if h+count >= val {
				return val
			} else {
				h += count
			}
		}
	}
	return h
}
