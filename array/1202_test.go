package array

import (
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1202. Smallest String With Swaps", func(t *testing.T) {
		input := "dcab"
		pairs := [][]int{{0, 3}, {1, 2}}
		want := "bacd"
		got := solution(input, pairs)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
[ref](https://leetcode.com/problems/smallest-string-with-swaps/discuss/525386/go-union-find)
给定字符串和可交换的位置，返回最小字典序的字符串
union-find
*/
// 找触点
func find(ds []int, u int) int {
	if ds[u] == u {
		return u
	}
	// 向上
	ds[u] = find(ds, ds[u])
	return ds[u]
}

// 合并
func union(ds []int, u, v int) {
	u = find(ds, u)
	v = find(ds, v)
	if u != v {
		ds[u] = v
	}
}

/*
	将可交换的所有索引存放到一个子集里面
	然后根据索引子集获取对应的字符，将字符排序后放回
*/
func solution(s string, pairs [][]int) string {
	sLen := len(s)
	ret := []byte(s)
	// 存放每个集合所在位置的信息
	ds := make([]int, sLen)
	// 初始化所有元素触点为自身
	for i := 0; i < sLen; i++ {
		ds[i] = i
	}
	// 可交换的索引保存在 slice
	m := make([][]int, sLen)
	// 将所有可以交换索引的合并
	for _, p := range pairs {
		union(ds, p[0], p[1])
	}
	// 将所有的可交换的索引位置保存下来
	for i, _ := range s {
		u := find(ds, i)
		m[u] = append(m[u], i)
	}
	for _, ids := range m {
		// 根据索引位置集合获取对应的字符
		bs := []byte{}
		for _, id := range ids {
			bs = append(bs, s[id])
		}
		// 每个集合按照字母序降序排列
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		// 每个集合字符保存到对应的位置
		for i, id := range ids {
			ret[id] = bs[i]
		}
	}
	return string(ret)
}
