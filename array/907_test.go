package array

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run(" 907. Sum of Subarray Minimums ", func(t *testing.T) {
		A := []int{3, 1, 4, 2}
		want := 17
		got := solution20(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" 907. Sum of Subarray Minimums2 ", func(t *testing.T) {
		A := []int{3, 1, 2, 4}
		want := 17
		got := solution20(A)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
	[lee](https://leetcode.com/problems/sum-of-subarray-minimums/discuss/170750/C%2B%2BJavaPython-Stack-Solution)
	[另一个解释](https://leetcode.com/problems/sum-of-subarray-minimums/discuss/178876/stack-solution-with-very-detailed-explanation-step-by-step)

	res = sum(A[i] * f(i)) A[i] 是最小值，f(i) 是子数组的数量
	计算 f(i)，left[i] 表示左边的所有元素都小于 A[i] 的元素数量
	right[i] 表示右边比 A[i] 大的元素数量
	left[i] + 1 表示以 A[i] 结尾的所有左边子数组数量，其中 A[i] 是最小值
	right[i] + 1 表示以 A[i] 为开头的右边子数组数量，其中 A[i] 为第一个最小值
	f(i) = (left[i] + 1) * (right[i] + 1)
	使用两个 stack 来计算 left[i], right[i] 单调递增栈
	注意这里要求子数组需要连续

	[3, 1, 2, 4] 例子
*/
// func solution10(A[]int) int {
// 	res, n, modNum := 0, len(A), 1000000007
// 	// 这里记录 元素数值，当前小于该元素数值的前面的元素数量
// 	// 两个都是单调递增栈
// 	s1, s2 := Stack{}, Stack{}
// 	left, right := make([]int, ALen), make([]int, ALen)
// 	// left array 0 到 n 
// 	for i := 0; i < n; i++ {
// 		// 栈中的元素大于当前元素的数量计数
// 		count := 1
// 		for !s1.empty() && s1.top().first > A[i] {
// 			count += s1.top().second
// 			s1.pop()
// 		}
// 		s1.push({A[i], count})
// 		left[i] = count
// 	}
// 	// right array n - 1 到 0 这边直接反序处理
// 	for i := n - 1; i >= 0; i-- {
// 		count := 1
// 		for !s2.empty() && s2.top().first > A[i] {
// 			count += s2.top().second
// 			s2.pop()
// 		}
// 		s1.push({A[i], count})
// 		right[i] = count
// 	}
	
// 	for i := 0; i < n; i++ {
// 		res = (res + A[i] * left[i] * right[i]) % modNum
// 	}
// 	return res
// }


/*
	找所有子数组的最小值的和
	类似 78 不对，这里要求所有的子数组的需要需要是连续的
	TLE
*/
// func solution00(A []int) int {
// 	sum := 0
// 	ALen := len(A)
// 	solutionHelper(A, &sum, []int{}, 0, ALen)
// 	return sum % (1000000000 + 7)
// }

// func solutionHelper(A []int, sum *int, tempArr []int, start, ALen int) {
// 	tempArrLen := len(tempArr)
// 	if tempArrLen > 0 {
// 		tempMin := 30001
// 		for _, v := range tempArr {
// 			if v < tempMin {
// 				tempMin = v
// 			}
// 		}
// 		*sum += tempMin
// 	}
// 	for i := start; i < ALen; i++ {
// 		if tempArrLen == 0 {
// 			tempArr = append(tempArr, A[i])
// 			solutionHelper(A, sum, tempArr, i+1, ALen)
// 			tempArr = tempArr[:tempArrLen]
// 		} else {
// 			if i+1 <= ALen {
// 				tempArr = append(tempArr, A[i])
// 				solutionHelper(A, sum, tempArr, i+1, ALen)
// 			}
// 			break
// 		}
// 	}
// }

/*
	直接 brute force?
*/
// func solution000(A []int) int {
// 	ALen := len(A)
// 	sum := 0
// 	for i := 0; i < ALen; i++ {
// 		sum += A[i]
// 		tempMin := A[i]
// 		for j := i + 1; j < ALen; j++ {
// 			if A[j] < tempMin {
// 				tempMin = A[j]
// 			}
// 			sum += tempMin
// 		}
// 	}
// 	return sum % (1000000000 + 7)
// }

/*
	使用数组来模拟栈 计算参考 solution40
*/
func solution20(A []int) int {
	ln := len(A)
	sum := 0
	mod := 1000000007
	if ln == 0 {
		return 0
	}
	if ln == 1 {
		return A[0]
	}
	// 模拟栈
	st := make([]int, ln)
	// 保存 A[i] 为结尾的当前最小元素的和
	sums := make([]int, ln)
	stLn := 0
	for i := 0; i < ln; i++ {
		// 非空找到第一个小于 A[i] 的位置
		for stLn > 0 &&  A[i] <= A[st[stLn-1]] {
			stLn--
		}
		// 当前元素的索引入栈
		st[stLn] = i
		// 当前元素小于全部的栈元素 注意这里是 sums[stLn]
		if stLn == 0 {
			sums[stLn] = (i + 1) * A[i]
		} else {
			sums[stLn] = sums[stLn-1] + (i-st[stLn-1])*A[i]
		}
		sum += sums[stLn]
		sum %= mod
		// 注意这里的自增，每次遍历完，栈顶的元素位置
		stLn++
	}
	return sum
}

// func solution3(A []int) int {
// 	stack := make([]int, 0)
// 	dp := make([]int, len(arr))
// 	res := 0
// 	for index, ele := range arr {
// 		//last which less than arr[index]
// 		k := -1
// 		//大的出栈
// 		for len(stack) > 0 && arr[stack[len(stack)-1]] >= ele {
// 			stack = stack[:len(stack)-1]
// 		}
// 		if len(stack) > 0 && arr[stack[len(stack)-1]] < ele {
// 			k = stack[len(stack)-1]
// 		}
// 		stack = append(stack, index)
// 		dp[index] += (index - k) * ele
// 		if k != -1 {
// 			dp[index] += dp[k]
// 		}
// 		dp[index] %= MOD
// 		res += dp[index]
// 		res %= MOD
// 	}
// 	return res
// }

/*
	可以使用 dp，观察到我们只需要关心以某个数字结尾时的子数组的最小值之和
	dp[i] 表示以 A[i] 结尾的所有子数组的最小值之和
	对于元素 A[i]，
	如果 A[i] 大于等于 A[i-1]，因为不会影响到 A[i]前的子数组
	的最小值，所以 A[i] 结尾所组成的子数组最小值之和为
	前面原来的最小值 dp[i-1] + A[i]
	如果 A[i] 小于 A[i-1]，则需要向前找到第一个比 A[i] 大的位置 j
		如果 j < 0 表示找不到该位置，则 A[i] 比前面所有元素都小，此时
		dp[i] = (i + 1) * A[i]
		如果 j >= 0，则有 dp[i] = dp[j] + (i - j) * A[i]
	注意这里的 dp[i] 不是累加下来的，只是对应的位置的最小值总和

	这里的问题是，每次遇到 A[i] 小于 A[i-1] 时就要向前遍历判断
	[参考](https://www.cnblogs.com/grandyang/p/11273330.html)

*/
func solution40(A []int) int {
	ALen := len(A)
	dp := make([]int, ALen)
	// 注意这里的 sum 以及 dp[0] 初始值
	sum, modNum := A[0], 1000000007
	dp[0] = A[0]
	// 从 1 开始遍历
	for i := 1; i < ALen; i++ {
		if A[i] >= A[i-1] {
			dp[i] = dp[i-1] + A[i]
		} else {
			// 这里 j 从 i - 1 开始 向前找第一个小于 A[i] 的位置索引
			j := i - 1
			for ; j >= 0 && A[j] > A[i]; j-- {
			}
			if j < 0 {
				dp[i] = (i + 1) * A[i]
			} else {
				dp[i] = dp[j] + (i-j)*A[i]
			}
		}
		sum = (sum + dp[i]) % modNum
	}
	return sum
}

/*
	使用单调栈
	将 dp 长度初始化为 len(A) + 1，栈先存放 -1
	dp[i] 表示 A[i - 1] 结尾的所有子数组最小值之和
	遍历数组，栈顶元素不是 -1 且 A[i] 小于等于栈顶元素，则将栈顶元素移除，这样
	栈顶元素就是前面第一个比 A[i] 小的数字（知道 j ？）。
	此时 dp[i + 1] 更新同上文。这个是单调递减的栈
	这里 如果 A[i] >= A[i - 1]，栈顶元素不需要移除，此时 dp[i] = dp[i-1] + A[i]
	索引以及栈的元素数量变化代替了上面的向前遍历查找
	注意这里的 stack 只是
*/
// func solution50(A []int) int {
// 	sum, modNum := 0, 1000000007
// 	ALen := len(A)
// 	dp := make([]int, ALen+1)
// 	stack := Stack{}
// 	stack.push(-1)
// 	// 从 0 开始
// 	for i := 0; i < ALen; i++ {
// 		// 先将所有的小于等于 A[i] 的出栈 这里 -1 用来方便计算
// 		for stack.top != -1 && A[i] <= A[stack.top] {
// 			stack.pop
// 		}
// 		dp[i+1] = (dp[stack.top+1] + (i-stack.top)*A[i]) % modNum
// 		stack.push(A[i])
// 		sum = (sum + dp[i+1]) % modNum
// 	}
// 	return sum
// }

/*
	含有 n 个元素的集合，非空子集数量是 2^n - 1 加入一个新的元素 i，
	以 i 结尾的元素数量是 n + 1
	对于每个数字，只需要知道前面第一个小于它的数字的位置以及后面第一个小于它的数字位置，
	就能知道当前数字是多少个子数组的最小值

	stack 保存序对，数值以及在数组中的索引，两个 stack st_pre, st_next 都是单调递减栈
	left, right，left[i] 表示以 A[i] 为结束且 A[i] 是最小值的子数组的数量；
	right[i] 表示以 A[i] 开始且 A[i] 是最小值的子数组数量
	st_pre 就是小于当前 A[i] 的索引位置
	right[i] 先初始化为 n - i，st_next 非空且栈顶元素 t > A[i]
	可以将 right[t.second] 更新为 i - t.second 
		这里的意思是，出栈的元素的 right[i] 更新
	最后，每个的 sum 是 left[i] * A[i] * right[i]
*/
// func solution60(A []int) int {
// 	sum, modNum := 0, 1000000007
// 	ALen := len(A)
// 	left, right := make([]int, ALen), make([]int, ALen)
// 	// 这里的栈保存的是序对 {value,index}
// 	st_pre, st_next := Stack{}, Stack{}
// 	// 从 0 开始
// 	for i := 0; i < ALen; i++ {
// 		// 非空，将大于 A[i] 的元素出栈
// 		for !st_pre.empty() && st_pre.top().first > A[i] {
// 			st_pre.pop()
// 		}
// 		// A[i] 左边的元素数值都大于 A[i]
// 		if st_pre.empty() {
// 			// A[i] 左边有 i + 1 个子数组的最小值是 A[i]
// 			left[i] = (i + 1)
// 	 	} else {
// 			 left[i] = i - st_pre.top().second
// 		}
// 		// 将 A[i] 入栈
// 		st_pre.push({A[i], i})
// 		// 先初始化
// 		right[i] = ALen - i
// 		// 非空，将大于 A[i] 的元素出栈
// 		for !st_next.empty() && st_next.top().first > A[i] {
// 			t := st_next.top()
// 			st_next.pop()
// 			// 更新栈顶的数值
// 			right[t.second] = i - t.second
// 		}
// 		// 将 A[i] 入栈
// 		st_next.push({A[i], i})
// 	}
// 	// 遍历处理
// 	for i := 0; i < ALen; i++ {
// 		sum = (sum + A[i] * left[i] * right[i]) % modNum
// 	}
// 	return sum
// }

// /*
// 	solution6 优化 只使用一个单调栈，记录当前数字之前的第一个小于它的数字的位置

// */
// func solution70(A[]int) int {
// 	sum, modNum := 0, 1000000007
// 	ALen := len(A)
// 	stack := Stack{}
// 	for i := 0; i < ALen + 1; i++ {
// 		cur := 0
// 		if i == ALen {
// 			cur = A[i]
// 		}
// 		// 非空且栈顶大于 0
// 		for !stack.empty() && cur < A[stack.top()] {
// 			index := stack.top()
// 			stack.pop()
// 			left := st.top()
// 			if stack.empty() {
// 				left = -1
// 			}
// 			right := i - index
// 			sum = (sum + A[index] * left * right) % modNum
// 		}
// 		stack.push(i)
// 	}
// 	return sum
// }