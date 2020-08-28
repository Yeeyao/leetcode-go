package array

import (
	"fmt"
	"sort"
)

/*
Suppose you have a random list of people standing in a queue.
Each person is described by a pair of integers (h, k), where h is the height of the person and
k is the number of people in front of this person who have a height greater than or equal to h.
Write an algorithm to reconstruct the queue.

给定表示人随机站立的列表
队列是每个人使用一个序对 (h,k) 表示，h 是人的高度，k 是在这个人之间高度大于或者等于 h 的人数
这里队列元素给出来了，编写算法重建队列

先选出 k 是 0 的元素，表示这个元素之前没有大于等于对应 h 的元素
这里 k 是 0 的元素需要 h 按照升序排列
按照高度降序排列，然后对相同高度的人中，按照 k 值升序排列

将最高的人按照 k 值升序排序，然后将它们放置到输出队列中与 k 值相等的索引位置上。
按降序取下一个高度，同样按 k 值对该身高的人升序排序，然后逐个插入到输出队列中与 k 值相等的索引位置上。
直到完成为止

[ref](https://leetcode-cn.com/problems/queue-reconstruction-by-height/solution/gen-ju-shen-gao-zhong-jian-dui-lie-by-leetcode/)

输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

[7,0][7,1][6,1][5,1][5,2][4,4]

*/
func reconstructQueue(people [][]int) [][]int {
	// 先按照升高 h 排序，身高相同就按照人数排序
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] ||
			(people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	// 元素插入到合适位置将元素都插入到 k 所在的为止
	// from 为当前索引 to 是元素的目标索引
	for from, p := range people {
		// 取 k 找到需要插入的为止
		to := p[1]
		// 向后移动
		copy(people[to+1:from+1], people[to:from])
		fmt.Println(people[to+1 : from+1])
		fmt.Println(people[to:from])
		fmt.Println(p)
		// 插入
		people[to] = p
	}
	return people
}
