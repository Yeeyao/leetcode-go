package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1122. Relative Sort Array", func(t *testing.T) {
		input := []string{"bella", "label", "roller"}
		want := []string{"e", "l", "l"}
		got := solution(input)
		if !StrSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1122. Relative Sort Array2", func(t *testing.T) {
		input := []string{"cool", "lock", "cook"}
		want := []string{"c", "o"}
		got := solution(input)
		if !StrSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1122. Relative Sort Array3", func(t *testing.T) {
		input := []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}
		want := []string{}
		got := solution(input)
		if !StrSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	单独统计每个单词中的字母出现数量，使用map并记录在 wordsCounter[i]
	将所有出现的字母记录下来 记录在 charCounter[i]
	当统计完了，针对每个 wordsCounter[i] 遍历
	其中，当出现在 charCounter 中的字母，
		没有出现在 wordsCounter[i] 则需要在 charCounter 中删除该字母
		如果已经出现了，则需要比较次数，取较小值作为 charCounter[i] 的值

	初始化一个最大值的每个单词出现次数的 int 数组或者 slice，key 是单词，value 是出现数量
	作为总体的次数数组
	遍历每个输入的单词，统计每个字母出现次数，更新总体次数数组，用较小值
	最后直接输出

*/
func solution(input []string) []string {
	// 最小数量数组，初始化为最大值:
	charCounter := ['z' - 'a' + 1]int{}
	var output []string
	for i := range charCounter {
		charCounter[i] = 1000
	}
	for _, str := range input {
		// 统计一个单词的字母出现数量
		strCounter := ['z' - 'a' + 1]int{}
		for _, char := range str {
			strCounter[char-'a']++
		}
		// 这里计算最小的出现数量
		for i, times := range strCounter {
			if times < charCounter[i] {
				charCounter[i] = times
			}
		}
	}
	for char, times := range charCounter {
		for i := 0; i < times; i++ {
			output = append(output, string(char+'a'))
		}
	}
	return output
}

/*
其他方法 基本思路是一样的，但是我的代码比较啰嗦
这里更好的是，首先初始化所有字母的出现次数列表
之后遍历每个字符串，首先对每个字符串都统计该字符串的字母出现数量
之后，比较全部出现次数并取较小值
最后根据所有字母出现次数列表来生成结果
*/
func commonChars(A []string) []string {
	// min count of letter among the words
	// 初始化所有字母的出现次数
	var minCount = ['z' - 'a' + 1]int{}
	for i := range minCount {
		minCount[i] = 1000
	}

	for _, str := range A {
		// per word
		var counts = ['z' - 'a' + 1]int{}
		for _, ch := range str {
			counts[ch-'a']++
		}

		for i, count := range counts {
			if count < minCount[i] {
				minCount[i] = count
			}
		}
	}

	var result = make([]string, 0)
	for i, count := range minCount {
		if count == 1000 {
			continue
		}

		for j := 0; j < count; j++ {
			result = append(result, string(i+'a'))
		}
	}

	return result
}

func StrSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/*
自己的旧代码 对每个单词都保存一个map，冗余。
同时，创建多个map还需要去除空的
最后，检查个数的处理也很啰嗦
*/
func commonChars(input []string) []string {
	wordCounterNum := 0
	charCounter := make(map[string]int)
	wordsCounter := make([]map[string]int, 100)
	var output []string
	// 这里每个单词自己统计出现的字母次数，同时记录总体的数量
	for i := 0; i < len(input); i++ {
		wordsCounter[i] = make(map[string]int, 100)
		for j := 0; j < len(input[i]); j++ {
			char := input[i][j]
			wordsCounter[i][string(char)]++
			charCounter[string(char)]++
		}
	}
	for i := 0; i < len(wordsCounter); i++ {
		if len(wordsCounter[i]) != 0 {
			wordCounterNum++
		}
	}

	for i := 0; i < wordCounterNum; i++ {
		for char, times := range charCounter {
			if wordCharTimes, ok := wordsCounter[i][char]; ok {
				if wordCharTimes < times {
					charCounter[char] = wordCharTimes
				}
			} else {
				delete(charCounter, char)
			}
		}
	}
	for char, times := range charCounter {
		for i := 0; i < times; i++ {
			output = append(output, char)
		}
	}
	return output
}
