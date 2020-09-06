package string

import (
	"reflect"
	"testing"
)

/*
438. Find All Anagrams in a String
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
Strings consists of lowercase English letters only and the length of
both strings s and p will not be larger than 20,100.
The order of output does not matter.
在一个给定的字符串中找到所有的易位构词 p 是所求的字符串 s 是目标字符串，返回每个易位构词的开始
字符串中的字母都是小写的同时长度小于 20100
[ref](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/solution/hua-dong-chuang-kou-tong-yong-si-xiang-jie-jue-zi-/)
TODO 其他滑动窗口题目
易位构词 直接统计 p 和 s 中每个字符的出现次数，然后直接进行各个字符的出现次数比较
滑动窗口
首先判断两个字符串的长度
*/
func TestPro(t *testing.T) {
	t.Run("438. Find All Anagrams in a String", func(t *testing.T) {
		input := "cbaebabacd"
		p := "abc"
		want := []int{0, 6}
		got := findAnagrams(input, p)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	优化版本， 直接使用一个 map 统计好目标串每个字符出现次数
	计算 s, p 长度并 sLen < pLen 直接返回
	统计 p 的字母出现频率数组 pFreq
	使用 left,right,count 表示左右指针位置以及未匹配的字符数量
	循环条件使右边没有到达 s 的结尾
		取右边的元素并判断 pFreq 对应字母是否大于 0 是就 count--
		将 pFreq 该字母的频率减小 1
		右指针右移一位
		如果 count == 0 将左指针加入到结果
		如果左右指针的字符等于 p 的长度
			如果左边字符频率大于等于 0 count++
			pFreq 该字符对应频率加 1
			左指针右移

	首先右边的先遍历到，然后不管是否存在于 p 都将对应的 pFreq 频率减小，然后左边遍历的时候将频率增加
	是否保存结果的关键是 count 的数值，只有在右边遍历的时候，频率大于 0 时 count--，左边再遍历的时候，
	如果频率大于等于 0 就 count++
*/
func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return []int{}
	}
	pFreq := make([]int, 26)
	for _, ps := range p {
		pFreq[ps-'a']++
	}
	res := make([]int, 0)
	left, right, count := 0, 0, pLen
	// 右指针需要小于 s 的长度
	for right < sLen {
		// 取右边的字符，如果匹配，将未匹配字符数量减少，对应的匹配字符频率减小，右指针右移一位
		cur := s[right]
		if curFreq := pFreq[cur-'a']; curFreq > 0 {
			count--
		}
		// 这里和左边的配合，将会把 pFreq 变回开始的时候 如果右边的字母不是 p 中的，这里将会变成负值，left 部分先判断是否 count 递增
		// 然后才还原
		pFreq[cur-'a']--
		right++
		// 如果未匹配字符数量为 0 表示全部都匹配了，直接保存左指针作为开始位置
		if count == 0 {
			res = append(res, left)
		}
		// 如果左右指针之间的字符数量等于 p 的长度因为循环开始直接判断的，所以这里需要提前处理左边指针的字符
		// 同时左指针位置移动一位
		// 这里如果右边的部分有匹配的，这里左边重新遇到直接会将 count 还原
		if right-left == pLen {
			if pFreq[s[left]-'a'] >= 0 {
				count++
			}
			pFreq[s[left]-'a']++
			left++
		}
	}
	return res
}

/*
	直接使用一个 map 统计好目标串每个字符出现次数
	然后使用临时字符串保存当前的字符串以及临时 map 保存当前字符串字符出现次数
*/
func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	bs, ps := []byte(s), []byte(p)
	bsLen, psLen := len(bs), len(ps)
	pMap := make(map[int]int, 26)
	for _, ps := range ps {
		pMap[int(ps-'a')]++
	}
	// 这里直接使用 int 保存了，将 byte 转换为 int
	tempInt := make([]int, 0)
	tempMap := make(map[int]int, 26)
	i := 0
	res := make([]int, 0)
	for i < bsLen {
		// 如果当前的字母不存在于目标中，则直接将临时字符串设置为空并跳过 用字母计数判断
		if pMap[int(bs[i]-'a')] == 0 {
			tempInt = []int{}
			tempMap = map[int]int{}
			i++
		} else {
			// 先将临时字符串字母计数器递增
			addInt := int(bs[i] - 'a')
			tempMap[addInt]++
			tempInt = append(tempInt, addInt)
			tempIntLen := len(tempInt)
			// 长度还不够，不需要比较
			if tempIntLen < psLen {
				i++
				continue
			} else {
				// 长度长一个，则移动窗口
				if tempIntLen > psLen {
					tempMap[tempInt[0]]--
					tempInt = tempInt[1:tempIntLen]
				}
				// 如果当前满足则保存索引
				if meet(tempMap, pMap) {
					res = append(res, i-psLen+1)
				}
			}
			i++
		}
	}
	return res
}

func meet(a, b map[int]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
