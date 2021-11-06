package array

import (
	"math"
	"sort"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("1996. The Number of Weak Characters in the Game", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		want := 4
		got := solution(input)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

/*
	玩游戏，包含多个字母，每个字母有攻击力和防御力两个数值。给定二维整型数组 properties，properties[i]=[attacki,defensei] 表示第 i 个字母。
	一个字母是 weak，如果任何其他的字母的攻击和防御力都比这个字母对应的攻击力和防御力高。一般化，如果字母 j attackj > attacki defensej > defensei
	则表示字母 i 是 weak。

	要求返回 weak 的字母的数量，这里需要一个字母的两个数值都比较小才是 weak
	这里还是需要排序的，只是排序之后怎么处理？比如按照攻击力排序，如果攻击力相同，则防御力降序

	暴力解法：直接每个元素都遍历一次整个数组，看是否有大于当前元素的，有就可以提前返回并加入计数
	排序优化：按照攻击力升序，如果相同就防御力降序

	主要是没领会 354 的意思，不然这里应该可以自己想出来
	类似 354 因为需要找到是否 weak，因此要先确定当前的最大的攻击力和防御力数值，从后面向前遍历，然后不断更新当前最大防御力
		攻击力按照升序排列了，所以就不需要判断了，只需要判断防御力
			如果攻击力不同，且前面的防御力小于后面的，就表示后面的攻击力和防御力都比前面的高，因此总数递增
			如果攻击力相同，排序后，先遍历的（后面的）防御力小于后遍历的（前面的），总数不递增
*/
func solution(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		a, b := properties[i], properties[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	res := 0
	mtn := math.MinInt32
	for i := len(properties) - 1; i >= 0; i-- {
		if properties[i][1] < mtn {
			res++
		}
		mtn = max(mtn, properties[i][1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
