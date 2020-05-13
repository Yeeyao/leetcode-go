package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 188. Best Time to Buy and Sell Stock IV ", func(t *testing.T) {
		input := []int{2, 4, 1}
		k := 2
		want := 2
		got := solution(k, input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	这里给定任意的 k
	每次价格中都要遍历 k 次
	T[i][k][1] = T[i-1][k-1][1]
	T[i][k][0] = T[i-1][k-1][0]

	T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	T[i][k][1] = max(T[i-1][k][1], T[i-1][k][0] - prices[i])
*/
func solution(k int, prices []int) int {
	// 这里一个优化，如果 k 的长度大于 len(prices) / 2，问题直接转换为 k = +INF
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	tik0, tik1 := 0, intMin
	pricesLen := len(prices)
	if k >= pricesLen>>1 {
		for _, p := range prices {
			if tik0 < tik1+p {
				tik0 = tik1 + p
			}
			if tik1 < tik0-p {
				tik1 = tik0 - p
			}
		}
		return tik0
	}

	// 这里需要存储每次 k 交易的结果，同时初始化为最小
	tik0Arr, tik1Arr := make([]int, pricesLen+1), make([]int, pricesLen+1)
	for i, _ := range tik1Arr {
		tik1Arr[i] = intMin
	}
	for _, p := range prices {
		// 处理 k 次交易 需要理解啊 同时注意处理顺序TODO:
		for j := k; j > 0; j-- {
			if tik0Arr[j] < tik1Arr[j]+p {
				tik0Arr[j] = tik1Arr[j] + p
			}
			// 这里需要 j - 1
			if tik1Arr[j] < tik0Arr[j-1]-p {
				tik1Arr[j] = tik0Arr[j-1] - p
			}
		}
	}
	return tik0Arr[k]
}
