package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("54 第一个只出现一次的字符", func(t *testing.T) {
		s := "abaccdeff"
		get := solution(s)
		want := byte('b')
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("54 第一个只出现一次的字符2", func(t *testing.T) {
		s := ""
		get := solution(s)
		want := byte(' ')
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	t.Run("54 第一个只出现一次的字符3", func(t *testing.T) {
		s := "abaccdeff"
		get := solution2(s)
		want := byte('b')
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("54 第一个只出现一次的字符4", func(t *testing.T) {
		s := ""
		get := solution2(s)
		want := byte(' ')
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
	相对第一个出现的重现字符使用 map，找到第一个重复的可以很快，这里采用这种则需要遍历整个字符
	如果这里也这样遍历，则需要两次遍历，第一次统计出现次数，第二次找到第一个出现次数是 1 的
	只包含小写字母 这里用 int slice
*/
//func solution(s string) byte {
//	byteS := []byte(s)
//	charTimes := make([]int, 26)
//	for _, c := range byteS {
//		charTimes[c-'a']++
//	}
//	for k, v := range charTimes {
//		if v == 1 {
//			return byte(k + 'a')
//		}
//	}
//	return ' '
//}

/*
	使用 map 默认是 false，遍历字符串，出现了一次就设置为 true，出现第二次就设置为 false
	然后重新遍历 string 取判断 map 对应的每个字符，返回第一个 true 的字符
*/
func solution(s string) byte {
	byteS := []byte(s)
	charTimes := make(map[byte]bool, 26)
	for _, c := range byteS {
		if _, ok := charTimes[c-'a']; ok {
			charTimes[c-'a'] = false
		} else {
			charTimes[c-'a'] = true
		}
	}
	for _, c := range byteS {
		if charTimes[c-'a'] == true {
			return c
		}
	}
	return ' '
}
