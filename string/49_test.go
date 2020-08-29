package string

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("49. Group Anagrams", func(t *testing.T) {
		strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		want := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
		got := groupAnagrams(strs)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.

给定字符串数组，返回 anagrams 组，anagrams 是通过对单词进行重排序得到
只要组成的字符相同的字符串就直接归到同一组

1. 将字符串都按照字典序进行排序后，排序后相同的字符串放在一起
2. 直接统计每个字母的出现次数然后相同的放在一起
*/

// 1
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string][]string, 0)
	for _, s := range strs {
		sp := strings.Split(s, "")
		sort.Strings(sp)
		sc := strings.Join(sp, "")
		resMap[sc] = append(resMap[sc], s)
	}
	res := make([][]string, 0)
	for _, v := range resMap {
		res = append(res, v)
	}
	return res
}

// 2 [26]int 作为 key
func groupAnagrams2(strs []string) [][]string {
	m := make(map[[26]int][]string)
	// 遍历每个字符串然后计算 key 并保存
	for _, s := range strs {
		key := count(s)
		store, ok := m[key]
		if !ok {
			store = make([]string, 0)
		}
		store = append(store, s)
		m[key] = store
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

/*
	统计每个单词的计数
*/
func count(s string) [26]int {
	res := [26]int{}
	for _, v := range s {
		res[v-'a']++
	}
	return res
}
