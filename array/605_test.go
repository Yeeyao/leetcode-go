package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("605. Can Place Flowers", func(t *testing.T) {
		input := []int{1, 0, 0, 0, 1}
		n := 1
		want := true
		got := solution(input, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("605. Can Place Flowers2", func(t *testing.T) {
		input := []int{1, 0, 0, 0, 1}
		n := 2
		want := false
		got := solution(input, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("605. Can Place Flowers3", func(t *testing.T) {
		input := []int{0, 0, 0, 1}
		n := 2
		want := false
		got := solution(input, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("605. Can Place Flowers4", func(t *testing.T) {
		input := []int{1, 0, 0, 0, 1, 0, 0}
		n := 2
		want := true
		got := solution(input, n)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	遍历下去，如果当前的上一个是 1，表示需要找到 3 个连续的 0，如果是结尾，则只需要找到 2 个
	如果上一个不是 1，表示需要找到两个连续的 0

	还是说，可以直接就找连续 0 的数量，
	如果是开头则需要找到 2 个就行
	如果是结尾则一样找 2 个
	中间则需要找 3 个
*/

/*
	这里直接计算 0 的个数来处理
	遇到 0 直接递增当前的 0 计数
	遇到 1 直接判断增加的空位，并将计数置 0
 */
func solution(input []int, n int) bool {
	inputLen := len(input)
	// 开头处理
	Count := 1
	ZeroCount := 0
	for i := 0; i < inputLen; i++ {
		if input[i] == 0 {
			Count++
		} else {
			ZeroCount += (Count - 1) / 2
			Count = 0
		}
	}
	// 结尾处理
	if Count != 0 {
		ZeroCount += Count / 2
	}
	return ZeroCount >= n
}

/*
	更快的方法 判断和修改
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		// 当前元素为 0
		if flowerbed[i] == 0 {
			// 第一个元素或者上一个元素为 0 同时最后元素或者下一个元素为 0
			if (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 >= length || flowerbed[i+1] == 0) {
				// 这里会将遍历过的元素置 1，其实就是模拟把花盘放置在该位置
				flowerbed[i] = 1
				n--
			}
			// 当前元素非 0 非两头的元素，同时自己和前后元素都是 1 这里没有什么必要吧
		} else {
			if (i-1 >= 0 && flowerbed[i-1] == 1) || (i+1 < length && flowerbed[i+1] == 1) {
				return false
			}
		}
	}
	if n >= 1 {
		return false
	}

	return true

}
