package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1414. Find the Minimum Number of Fibonacci Numbers Whose Sum Is K", func(t *testing.T) {
		input := 8
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1414. Find the Minimum Number of Fibonacci Numbers Whose Sum Is K2", func(t *testing.T) {
		input := 7
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1414. Find the Minimum Number of Fibonacci Numbers Whose Sum Is K3", func(t *testing.T) {
		input := 10
		want := 2
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("1414. Find the Minimum Number of Fibonacci Numbers Whose Sum Is K4", func(t *testing.T) {
		input := 19
		want := 3
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	需要找到斐波那契数列的数来组成给定的和的斐波那契数列数值的最小数量
	需要构造斐波那契数列，数量是刚好小于给到的数值
	然后利用找到的数组数值，转变成求和

	求和，从最大的数值开始遍历

		剩余的数值等于 k - 当前的数值

			剩余数值大于 0 继续看上一个元素，数量 + 1

			剩余数值小于 0 直接跳过当前元素

			剩余数值等于 0 表示找到了数量了，直接比较并修改

	能不能利用斐波那契数列的特点 后一项等于前面两项之和
*/
func solution(k int) int {
	fa := genFa(k)
	// 求数量
	faLen := len(fa)
	minCount := faLen
	for i := faLen - 1; i >= 0; i-- {
		left := k
		count := 0
		for j := i; j >= 0; j-- {
			thisLeft := left - fa[j]
			if thisLeft > 0 {
				count++
				left = thisLeft
			}
			if thisLeft == 0 {
				count++
				if count > 0 && count < minCount {
					minCount = count
				}
			}
		}
	}
	return minCount
}

/*
	生成斐波那契数列，小于等于 k
*/
func genFa(k int) []int {
	retArr := []int{1, 1}
	if k == 1 {
		return []int{1}
	}
	if k == 2 {
		return []int{1, 1}
	}
	for i := 2; retArr[i-1]+retArr[i-2] <= k; i++ {
		retArr = append(retArr, retArr[i-1]+retArr[i-2])
	}
	return retArr
}
