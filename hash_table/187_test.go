package hash_table

import (
	"math"
	"reflect"
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("187. Repeated DNA Sequences", func(t *testing.T) {
		s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
		want := []string{"AAAAACCCCC", "CCCCCAAAAA"}
		got := solution(s)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
187. Repeated DNA Sequences
All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T', for example: "ACGAATTCCG".
When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.
Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

所有的 DNA 是由 A C G T 四种碱基组成，编写一个函数，找到在 DNS 碱基序列中出现重复出现的 10 个字母长的碱基序列
这里有点类似重复子串的问题，找多个重复子串

这里是需要 10 个一组作为重复的计算，其实可以推广到一般性的长度 L 来计算
这里题解使用滑动窗口以及 hashset
线性时间获得滑动窗口的方法比较笨，同时消耗大量时间和空间，都是 O((N-L)L)
两种改进的方法：Rabin-Karp 算法 = 使用旋转哈希算法实现常数窗口切片 位操作 = 使用掩码实现常数窗口切片

Rabin-Karp 算法
	思想是对字符串进行切片并在滑动窗口中计算序列的哈希值，两者都是在常数时间内进行
	首先将字符串转换为整数数组，计算序列的哈希值。这里使用 4 进制计算，后续的哈希值从上一个哈希值中计算出来，所以叫旋转哈希
	每次计算哈希值将整体左移，然后将上一个的哈希值最高位去掉，最后加上新遍历的数值

*/

// Rabin Karp 时间福啊咋读 O(N - L) 空间复杂度 O(N)
func solution(s string) []string {
	L, n := 10, len(s)
	if n <= L {
		return []string{}
	}
	a := 4
	// 看作一个 4 进制的数字，这里是最高位的基数
	aL := int(math.Pow(float64(a), float64(L)))
	// 将所有的字符串转换为对应的数字
	numMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = numMap[s[i]]
	}

	h := 0
	seen := make(map[int]bool)
	output := make(map[string]bool)

	// 遍历所有的字符 使用旋转哈希函数同时记录下已经出现过的字符串的哈希值
	for start := 0; start < n-L+1; start++ {
		// 不是从头开始计算 h 向左移动一位，减去最高位，然后加上新的一位
		if start != 0 {
			h = h*a - nums[start-1]*aL + nums[start+L-1]
		} else {
			// 第一次计算 h
			for i := 0; i < L; i++ {
				h = h*a + nums[i]
			}
		}
		// 已经出现过
		if _, ok := seen[h]; ok {
			output[s[start:start+L]] = true
		}
		seen[h] = true
	}
	res := make([]string, 0)
	for rs, _ := range output {
		res = append(res, rs)
	}
	return res
}

/*
	位操作 对字符串进行切片，在滑动窗口中计算序列的腌码，两者在恒定时间内进行
	和 Rabin-Karp 一样，先将字符串转换为两个比特位整数数组
	思路类似旋转哈希，只是使用掩码来处理

	A -> 0 = 00(2) C -> 1 = 01(2) G -> 2 = 10(2) T -> 3 = 11(2)
	这里二进制序列中每个数字占用位不超过 2 位，因此可以在循环中计算掩码
		左移释放最后两位 bitmask <<= 2
		将当前数字存储到移动后的后两位 bitmask ｜= nums[i]
	这里 Java 的运算符优先级，3 << 2 * L 先计算乘法，再计算移位

*/
func solution2(s string) []string {
	// 前面一样将字符串转换为数字数组
	L, n := 10, len(s)
	if n <= L {
		return []string{}
	}
	// 将所有的字符串转换为对应的数字
	numMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = numMap[s[i]]
	}
	// 掩码
	bitMask := 0
	seen := make(map[int]bool)
	output := make(map[string]bool)

	for start := 0; start < n-L+1; start++ {
		if start != 0 {
			// 先左移释放最后的两位
			bitMask <<= 2
			// 向最后两位添加新的两个二进制
			bitMask |= nums[start+L-1]
			// 重置最前面的两位
			bitMask &= ^(3 << (2 * L))
		} else {
			for i := 0; i < L; i++ {
				bitMask <<= 2
				bitMask |= nums[i]
			}
		}
		// 已经出现过
		if _, ok := seen[bitMask]; ok {
			output[s[start:start+L]] = true
		}
		seen[bitMask] = true
	}
	res := make([]string, 0)
	for rs, _ := range output {
		res = append(res, rs)
	}
	return res
}

// 线性滑动窗口
func solution3(s string) []string {
	L, N := 10, len(s)
	hashSet := make(map[string]bool)
	res := make([]string, 0)
	for i := 0; i < N-L+1; i++ {
		tmp := s[i : i+L]
		if _, ok := hashSet[tmp]; ok {
			res = append(res, tmp)
		} else {
			hashSet[tmp] = true
		}
	}
	return res
}
