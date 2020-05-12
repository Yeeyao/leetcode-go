package array

import (
	"math"
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("954. Array of Doubled Pairs", func(t *testing.T) {
		A := []int{3, 1, 3, 6}
		want := false
		got := solution(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("954. Array of Doubled Pairs2", func(t *testing.T) {
		A := []int{4, -2, 2, -4}
		want := true
		got := solution(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	Lee 的方法
	这里的思路是统计所有元素出现的数量，然后进行数量的对比处理
	根据绝对值的大小顺序来遍历元素
	若 c[i] > c[2 * i] 则直接返回 false
	若 c[i] <= c[2 * i] 则将 c[2 * i] 的数目减少 c[i]
	对于 0 如果数量是奇数，则可以直接返回 false，偶数则满足
	遍历计数数组，如果当前元素数量为 0 则跳过
	当前元素数值为正数，则需要找 i * 2 的元素数量，为负数，则需要找 i / 2 的元素数量
	如果元素数量

	先将 A 按照绝对值大小排序，这样，负数 a 的目标数值也是 a / 2
	遍历 A 的元素，
		如果是偶数同时 a / 2的数量大于 0 就将 a / 2 数量递减
			同时如果 a / 2 的数量是 0 就直接从计数 map 中删除
		否则，将 a 的计数递增
	最后，如果计数 map 长度是 0 表示全部元素都统计并消耗完了，就返回 true

*/
func solution(A []int) bool {
	aLen := len(A)
	if aLen == 0 {
		return true
	}
	countMap := make(map[int]int)
	// 直接将 A 按照绝对值大小排序
	sort.SliceStable(A, func(i, j int) bool {
		return abs(A[i]) < abs(A[j])
	})
	for _, v := range A {
		// 偶数且目标数值数量大于 0 才处理
		if v%2 == 0 && countMap[v/2] > 0 {
			countMap[v/2]--
			if countMap[v/2] == 0 {
				delete(countMap, v/2)
			}
		} else {
			countMap[v]++
		}
	}
	return len(countMap) == 0
}

/*
	上面的改进形式，先计数，然后判断数量
	当计数大于 0 时，判断是否大于目标的数量，是则返回 false 因为已经不够了
	否则，目标需要减去当前元素的数量表示消耗了，然后当前元素的数量置 0
*/
func solution2(A []int) bool {
	// 绝对值排序
	sort.Slice(A, func(i, j int) bool { return math.Abs(float64(A[i])) < math.Abs(float64(A[j])) })
	m := make(map[int]int)
	// 先计数
	for _, k := range A {
		m[k]++
	}
	// 遍历并消耗处理
	for _, v := range A {
		if m[v] > 0 {
			// 比目标还多则直接返回 false
			if m[v] > m[2*v] {
				return false
			}
			// 消耗目标的数量
			m[2*v] -= m[v]
			m[v] = 0
		}
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

/*
	判断给定的数组是否可以将其重排，使得从 0 到 len(A) / 2 的 i
	A 长度是偶数
	满足 A[2 * i + 1] = 2 * A[2 * i]
	i = 0 需要 A[1] = 2A[0] i = 1 A[3] = 2A[2] i = 2 A[5] = 2A[4]
	看例子的话，应该是需要排序，然后对负数进行处理一下
	重复元素也要处理 deprecated
*/
func solution3(A []int) bool {
	aLen := len(A)
	if aLen == 0 {
		return true
	}
	sort.Ints(A)
	for i := 0; i < aLen; i += 2 {
		if A[i] > 0 {
			if A[i+1] != 2*A[i] {
				return false
			}
		}
		if A[i] < 0 {
			if 2*A[i+1] != A[i] {
				return false
			}
		}
	}
	return true
}
