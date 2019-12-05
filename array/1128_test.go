package array

import "testing"

func TestPro(t *testing.T) {
	t.Run("1128. Number of Equivalent Domino Pairs", func(t *testing.T) {
		input := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
		want := 1
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	brute force

	先从保存的数组中找是否存在对应的元素
	存在则数量 + 1
	不存在跳过
	保存当前的数组

	下面是小改进，通过将序对进行特征计算来计数
*/

func solution(input [][]int) int {
	count := make(map[int]int)
	sum := 0
	for _, v := range input {
		var num int
		if v[0] > v[1] {
			num = v[0] + v[1]*10
		} else {
			num = v[0]*10 + v[1]
		}
		count[num]++
	}
	for _, v := range count {
		sum += v * (v - 1) / 2
	}
	return sum
}

func numEquivDominoPairs(dominoes [][]int) int {
	var res int
	hash := make(map[[2]int]int)
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		tmp := [2]int{v[0], v[1]}
		if v, _ := hash[tmp]; v > 0 {
			res += v
		}
		hash[tmp]++
	}
	return res
}
