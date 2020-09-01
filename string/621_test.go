package string

import (
	"sort"
	"testing"
)

/*
Given a characters array tasks, representing the tasks a CPU needs to do, where each letter represents a different task.
Tasks could be done in any order. Each task is done in one unit of time.
For each unit of time, the CPU could complete either one task or just be idle.
However, there is a non-negative integer n that represents the cooldown period between two same tasks
(the same letter in the array), that is that there must be at least n units of time between any two same tasks.
Return the least number of units of times that the CPU will take to finish all the given tasks.

给定要给字符数组任务，不同的字母表示不同的任务，CPU 需要处理它们。每个任务需要一单位时间处理，可以以任意顺序完成任务
每个单元时间，CPU 可以完成任务或者保持空闲
两个任务之间有一个整型 n 表示它们之间的冷却时间 n 即同一个任务间至少需要间隔 n 时间
返回 CPU 完成全部任务所需的最少时间
*/

func TestPro(t *testing.T) {
	t.Run("621. Task Scheduler", func(t *testing.T) {
		task := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
		n := 2
		want := 8
		got := leastInterval(task, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	尽量在相同的任务中插入其他任务，感觉有点像高矮重排
	n 为 0 直接返回长度了
	n 非 0 排序字母后遍历，两个相同的字母间插入不同的字母
		插入的字母个数最大是 n，如果不够就需要 idle
	先遍历，统计每个字母的出现次数，然后按出现次数降序排列，
	n + 1 个任务为一轮，这样同一轮任务最多只能安排一次，选择剩余次数最多的 n + 1 个任务依次执行
	如果任务种类 t 小于 n + 1 则选择全部 t 个任务，其余时间空闲

	建立 map 来统计每个字符出现次数，然后排序
	循环条件是 map[25] > 0
		初始化 i 为 0
		循环条件是 i <= n(n + 1 一轮)
			如果 map[25] == 0 直接 break
			如果 i < 26 同时 map[25 - i] > 0 即后面的字符次数大于 0
				将次数减少 1
			总时间递增
			i 递增取下一个元素
		重新进行排序
	最后返回时间
	这里排序效率也太低了
*/

/*
[ref](https://leetcode-cn.com/problems/task-scheduler/solution/ren-wu-diao-du-qi-by-leetcode/)
记任务 A 出现次数最多为 countMax，如果有其他任务 B ... 出现次数也是 countMax，
则直接 ABXXXABXXXABXXX 这样安排任务顺序，所以其中一个结果
	(counter[25] - 1) *(n + 1) 是最大次数 - 1的任务加上后面的空格总时间，然后加上补全的任务时间
对于 AAABBB 0 这种情况，就需要再和任务长度比较取较大值

统计并排序，初始化 i = 25 然后如果其他任务出现次数也是最大次数，i 递减
最后返回  max(len(task) (counter[25] - 1) * (n + 1) + 25 - i)
*/
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	counter := make([]int, 26)
	for _, t := range tasks {
		counter[t-'A']++
	}
	sort.Ints(counter)
	i := 25
	for i >= 0 && counter[i] == counter[25] {
		i--
	}
	return max(len(tasks), (counter[25]-1)*(n+1)+25-i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func leastInterval2(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	counter := make([]int, 26)
	time := 0
	for _, t := range tasks {
		counter[t-'A']++
	}
	// 先进行排序
	sort.Ints(counter)
	for counter[25] > 0 {
		i := 0
		for i <= n {
			// 如果最后的出现次数等于 0 跳过该遍历
			if counter[25] == 0 {
				break
			}
			if i < 26 && counter[25-i] > 0 {
				counter[25-i]--
			}
			time++
			i++
		}
		sort.Ints(counter)
	}
	return time
}

/*
	优先队列 选择每一轮任务时，使用优先队列替代排序
	一开始，我们将所有任务加入优先队列，每一轮，从其中选出最多的 n + 1 个任务，将它们数量 - 1
	再放回去，直到队列为空
*/

/*
[ref](https://leetcode.com/problems/task-scheduler/discuss/104496/concise-Java-Solution-O(N)-time-O(26)-space)
	设计
	A 为出现次数最多的任务，假设出现 p 次，则执行完所有任务的时间至少是
		(p - 1) * (n + 1) + 1
	同时 CPU 产生了 (p - 1) * n 个空闲时间，这里就是时间间隔
	考虑将剩余的任务安排到空闲时间当中，按照任务的出现次数排序，降序
		某个任务和 A 出现次数相同，只能让 B 占据 p - 1 个空闲时间，非空闲时间需要额外安排一个时间给 B 执行
		某个任务 C 比 A 出现次数少 1，直接让 C 占据 p - 1 个空闲时间就行
		某个任务 D 比 A 出现次数少 2 或者更多，按照列优先顺序将 D 填入空闲空间
	将所有任务安排完成，还有剩余空闲时间，则总时间是 任务总数 + 剩余空闲时间

使用 map 统计每个字符出现次数然后进行排序，选中最大的次数并计算空闲位置
从最大次数开始，空闲位置减去最大的次数，最后判断空闲位置数量
*/
func leastInterval3(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	counter := make([]int, 26)
	for _, t := range tasks {
		counter[t-'A']++
	}
	sort.Ints(counter)
	maxVal := counter[25] - 1
	idleSlots := maxVal * n
	for i := 24; i >= 0; i-- {
		idleSlots -= min(counter[i], maxVal)
	}
	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}
	return len(tasks)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
