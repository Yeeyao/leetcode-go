package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 1014. Best Sightseeing Pair ", func(t *testing.T) {
		input := []int{8, 1, 5, 2, 6}
		want := 11
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 1014. Best Sightseeing Pair2 ", func(t *testing.T) {
		input := []int{8, 1, 5, 2, 6}
		want := 11
		got := solution2(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	求数组中最大的两数以及距离之差(小的索引减去大的索引)
	因为这里没有说明所有元素只出现一次

	将 a[i] + a[j] + i - j 看成(a[i] + i) + (a[j] - j) 左右两个部分
	初始化 ret = a[0] + a[1] - 1
	循环中，先将当前的元素作为 right 部分，更新总和
	然后判断当前元素是否适合作为 left
*/
func solution(A []int) int {
	left := 0
	ret := A[0] + A[1] - 1
	aLen := len(A)
	for i := 1; i < aLen; i++ {
		retT := A[left] + left + A[i] - i
		if retT > ret {
			ret = retT
		}
		if A[i]+i > A[left]+left {
			left = i
		}
	}
	return ret
}

/*
	这里更新 cur 跟上面的是一样的道理，选择 left
	res 也是一样地先用 cur 计算新的总和
	每次循环的 -1 是跟上面一样的，只不过是代码简化了。
	说实话，那么难理解的代码真是。但是简洁啊
*/
func solution2(A []int) int {
	res, cur := 0, 0
	for _, v := range A {
		if cur+v > res {
			res = cur + v
		}
		if v > cur {
			cur = v
		}
		cur--
	}
	return res
}
