package hash_table

import "testing"

/*
771. Jewels and Stones
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.
Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
The letters in J are guaranteed distinct, and all characters in J and S are letters.
Letters are case sensitive, so "a" is considered a different type of stone from "A".
给定字符串 J 表示宝石的石头类型，S 表示你拥有的石头类型，求 S 中宝石的数量
直接建立一个 map 保存每个宝石的类型，然后遍历 S 每个石头，如果是宝石类型（map 中找到该类型）就计数
下面的方法是遍历 j
*/

func TestPro(t *testing.T) {
	t.Run("leetcode 771 Jewels and Stones", func(t *testing.T) {
		inputone := "aA"
		inputtwo := "aAAbbbb"
		want := 3
		got := solution(inputone, inputtwo)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func solution(j, s string) int {
	var sum = 0
	for _, v := range s {
		if charinj(j, v) {
			sum += 1
		}
	}
	return sum
}

func charinj(j string, c rune) bool {
	for _, v := range j {
		if v == c {
			return true
		}
	}
	return false
}
