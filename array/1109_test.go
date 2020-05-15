package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("59. Spiral Matrix II", func(t *testing.T) {
		bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
		n := 5
		want := []int{10, 55, 45, 25, 25}
		got := solution(bookings, n)
		if !IntSliceEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	航班预约 b[i] = [i, j, k] 表示我们预定了从 i 到 j 航班中的 k 个座位
	给定 n，代表航班的长度，需要返回 n 中每个航班的座位预定总数
*/
func solution(bookings [][]int, n int) []int {
	// 根据 n 构造，然后遍历每个元素，把预定的加到总和上面
	retArr := make([]int, n)
	for _, v := range bookings {
		begin, end, seat := v[0], v[1], v[2]
		end2 := end
		// 超过要求的航班数量
		if begin > n {
			continue
		}
		// 最后的航班编号
		if end2 > n {
			end2 = n
		}
		for i := begin; i <= end2; i++ {
			retArr[i-1] += seat
		}
	}
	return retArr
}

/*
	对于 [i, j, k]
	我们知道在第 i 天需要 k 个座位，然后在 j + 1 天已经不需要了
	单独看一个人的 bookings 就可以理解，循环中只是将所有人的累加起来而已
*/
func solution2(bookings [][]int, n int) []int {
	retArr := make([]int, n)
	// 这里将所有的人的数据累计而已
	// 单独看的话，对于一个人，i 到 j 天都要相同的座位数量，然后对 j + 1 天，
	// 需要判断跟 n 的关系。
	// 因为已经不在预定范围内了，所有后面就不能累加了，所以在这天就需要减去座位数量
	for _, v := range bookings {
		begin, end, seat := v[0], v[1], v[2]
		retArr[begin-1] += seat
		// 预定天数小于 n 需要在最后减去座位数量
		if end < n {
			retArr[end] -= seat
		}
	}
	for i := 1; i < n; i++ {
		retArr[i] += retArr[i-1]
	}
	return retArr
}

func IntSliceEqual(a, b []int) bool {
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
