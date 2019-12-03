package array

import (
	"fmt"
	"testing"
)

// 每个字母只能使用一次
func TestPro(t *testing.T) {
	t.Run("leetcode 1160  Find Words That Can Be Formed by Characters", func(t *testing.T) {
		words := []string{"cat", "bt", "hat", "tree"}
		chars := "atach"
		want := 6
		got := solution(words, chars)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("leetcode 1160  Find Words That Can Be Formed by Characters", func(t *testing.T) {
		words := []string{"hello", "world", "leetcode"}
		chars := "welldonehoneyr"
		want := 10
		got := solution(words, chars)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func solution(words []string, chars string) int {
	length := 0
	for _, v := range words {
		byteChars := []byte(chars)
		if isGoods(v, byteChars) {
			length = length + len(v)
		}
	}
	return length
}

func isGoods(word string, byteChars []byte) bool {
	// 每次遍历需要重置是否找到标记
	for i := 0; i < len(word); i++ {
		isFind := false
		for j := 0; j < len(byteChars); j++ {
			if word[i] == byteChars[j] {
				isFind = true
				byteChars[j] = ' '
				break
			}
		}
		if !isFind {
			return false
		}
		isFind = false
		fmt.Println(byteChars)
	}
	return true
}
