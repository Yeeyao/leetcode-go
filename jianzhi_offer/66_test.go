package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("66 机器人的运动范围", func(t *testing.T) {
		m := 2
		n := 3
		k := 1
		get := solution(m, n, k)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	count = 0
	t.Run("66 机器人的运动范围2", func(t *testing.T) {
		m := 3
		n := 1
		k := 0
		get := solution(m, n, k)
		want := 1
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})

	t.Run("66 机器人的运动范围3", func(t *testing.T) {
		m := 2
		n := 3
		k := 1
		get := solution2(m, n, k)
		want := 3
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
	count = 0
	t.Run("66 机器人的运动范围4", func(t *testing.T) {
		m := 3
		n := 1
		k := 0
		get := solution2(m, n, k)
		want := 1
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

也是 dfs 加上剪枝 这里需要额外的空间保存是否访问过
终止条件是 坐标的和大于 k 以及访问过该节点
将大于 k 的节点看作是障碍就可以简化模型，同时，可以朝下和右移动来优化

初始化一个一已遍历二维 slice visit 调用 dfs 参数是 m, n, 0, 0, k, visit
dfs
	元素已经越界或者已经访问或者坐标和超过 k 返回
	否则，总计数器 +1 记录已经访问并递归向下和向右调用
*/

var count int

func solution(m, n, k int) int {
	visit := make([][]int, m)
	for i, _ := range visit {
		visit[i] = make([]int, n)
	}
	dfs(m, n, 0, 0, k, visit)
	return count
}

func dfs(m, n, i, j, k int, visit [][]int) {
	if i < 0 || i >= m || j < 0 || j >= n || visit[i][j] != 0 || calcSum(i)+calcSum(j) > k {
		return
	}
	count++
	visit[i][j] = 1
	//dfs(m, n, i-1, j, k, visit)
	dfs(m, n, i+1, j, k, visit)
	//dfs(m, n, i, j-1, k, visit)
	dfs(m, n, i, j+1, k, visit)
}

func calcSum(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum

}

/*
	递推处理
	visit[i][j] 定义 (i, j) 的可达性，可达得 1，不可达得 0
	可以知道 visit[i][i] = visit[i-1][j] | visit[i][j-1]
*/
func solution2(m, n, k int) int {
	if k == 0 {
		return 1
	}
	res := 1
	visit := make([][]int, m)
	for i, _ := range visit {
		visit[i] = make([]int, n)
	}
	// 初始化
	visit[0][0] = 1
	// 需要跳过边界或者不满足的总和
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || calcSum(i)+calcSum(j) > k {
				continue
			}
			if i-1 >= 0 {
				visit[i][j] |= visit[i-1][j]
			}
			if j-1 >= 0 {
				visit[i][j] |= visit[i][j-1]
			}
			// 累加和
			res += visit[i][j]
		}
	}
	return res
}
