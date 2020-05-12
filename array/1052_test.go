package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1052. Grumpy Bookstore Owner", func(t *testing.T) {
		customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
		grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
		X := 3
		want := 16
		got := solution(customers, grumpy, X)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	书店将开启 customers 长度的分钟时间，每分钟，有 customers[i]
	进入书店。同时，老板在 grumpy[i] 为 1 时是 grumpy 的，
	然后顾客 customers[i] 将不满意。
	老板可以让自己在 X 分钟内不会 grumpy，但是只能使用一次
	返回一天内使最多顾客满意的分钟时间长度

	这里 X 可以直接将 grumpy 连续几个元素设置成 0
	暴力方法是对所有的顺序都尝试一次，然后计算每个顺序的时间，但是很耗时
	对应的 grumpy[i] 为 0 就可以使对应的 customers[i] 时间的所有顾客满意
	需要注意这里求得是最大满意顾客数量

	如果遍历，移动一位，其实只有首尾数量变化
	一开始就全部计算，然后，不断移动 x 区间来处理

*/
func solution(customers, grumpy []int, X int) int {
	// 总共服务时间
	serverMinutes := len(customers)
	if serverMinutes == 0 {
		return 0
	}
	// 先全部加起来
	totalSum := 0
	maxSum := 0
	for i := 0; i < serverMinutes; i++ {
		if grumpy[i] == 0 {
			totalSum += customers[i]
		}
	}
	// 移动 x 的区间 第一次
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			totalSum += customers[i]
		}
	}
	if totalSum > maxSum {
		maxSum = totalSum
	}

	// x 区间范围处理
	for i := 0; i < serverMinutes-X; i++ {
		// 移动前 x 第一个位置为 1 移动后需要减去
		if grumpy[i] == 1 {
			totalSum -= customers[i]
		}
		// 移动后 x 最后一个位置为 1 的 移动后需要加上去
		if grumpy[i+X] == 1 {
			totalSum += customers[i+X]
		}
		if totalSum > maxSum {
			maxSum = totalSum
		}
	}
	return maxSum
}
