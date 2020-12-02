package hash_table

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("166. Fraction to Recurring Decimal", func(t *testing.T) {
		numerator, denominator := 1, 2
		want := "0.5"
		got := solution(numerator, denominator)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
166. Fraction to Recurring Decimal
给定分数的分子和分母，返回他们的小数形式。小数中有循环部分，使用（）将循环部分包含。
如果有多个结果，返回其中一个
[ref](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/solution/fen-shu-dao-xiao-shu-by-leetcode/)

思路是长除法，核心思想是余数出现循环时，对应的商也会出现循环
算法：需要用一个哈希表来记录余数出现在小数部分的位置，当发现已经出现的余数，就可以将重复出现的小数部分用括号括起来
除法过程中出现 0 就停止程序
需要考虑各种边界

直接上述的思路
*/

func solution(numerator int, denominator int) string {
	if denominator == 0 {
		return "NAN"
	}
	// 结果
	var tmp string
	// 其中一个是负数，需要加上负号
	if numerator*denominator < 0 {
		tmp += "-"
	}
	// 都转为负数
	numerator, denominator = abs(numerator), abs(denominator)
	// 先加上当前结果
	tmp += strconv.Itoa(numerator / denominator)
	// 余数部分
	num := numerator % denominator
	// 余数是 0 直接返回
	if num == 0 {
		return tmp
	}
	tmp += "."

	// 保存余数出现的位置
	hm := make(map[int]int)
	rpPos := -1
	for {
		// 余数处理
		num *= 10
		// 出现重复余数
		if pos, ok := hm[num]; ok {
			// 记录重复的余数位置
			rpPos = pos
			break
		} else {
			hm[num] = len(tmp)
		}
		tmp += strconv.Itoa(num / denominator)
		num %= denominator
		// 出现 0 要退出循环
		if num == 0 {
			break
		}
	}
	// 没有重复出现的余数
	if rpPos == -1 {
		return tmp
	}
	return fmt.Sprintf("%s(%s)", tmp[:rpPos], tmp[rpPos:])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
