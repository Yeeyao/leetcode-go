package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("34 第一个只出现一次的字符", func(t *testing.T) {
		s := "abaccdeff"
		get := solution(s)
		want := byte('b')
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
	单纯统计每个字母出现的次数然后再遍历判断第一个只出现一次的感觉
	直接使用 map 同时需要注意的是 map 的顺序遍历问题，没法顺序遍历，所以只能用 int slice
*/
func solution(s string) byte {
	num := 'z' - 'a'
	countMap := make([]int, num)
	for _, str := range s {
		countMap[str-'a']++
	}
	for k, v := range countMap {
		if v == 1 {
			return byte(k + 'a')
		}
	}
	return ' '
}
