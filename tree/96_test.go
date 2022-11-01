package tree

/*
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
给定 n 计算有多少个从 1 到 n 的二叉搜索树
[ref](https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/)

更像是数学题
遍历每个数字 i，将该数字作为根，将 1...(i - 1) 作为左子树，(i+1)...n 作为右子树，然后同样方式递归创建左右子树
定义两个函数
	G(n) 长度为 n 的序列能构成的不同二叉搜索树的个数
	F(i,n) 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1≤i≤n)
可以知道
	G(n) = 求和(F(i,n)) 其中 i 从 1 到 n
	G(0) = 1 G(1) = 1
选择 i 作为根，则根 i 的所有 BFS 集合是左子树的集合和右子树集合的笛卡尔积(因为这里左子树和右子树可以任意匹配)
	反过来 F(i, n) = G(i-1) * G(n-i) 因为是 BFS 左子树和右子树数值有限制 G(n) 与具体的树无关，只是表示数量
因此 G(n) = 求和(G(i - 1) * G(n - i)) 其中 i 从 1 到 n
即 G(n) = G(0) * G(n) + G(1) + G(n-1) + ... + G(n-1) * G(0)
递归计算所有的数值

构造长度为 n + 1 的数组 G
G[0], G[1] 初始化为 1
第一层循环 i 从 2 到 n 因为 0, 1 已经计算了
	第二次循环 j 从 1 到 i
		G[i] += G[j-1]*G[i-j]

*/
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	// 想要计算 G[i] = G[i-1] * G[n-i]，这里 G[n-i] 没法直接计算
	// 因此需要往下计算然后向上加
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

/*
	数学方法  卡塔兰数 Cn
	C0 = 1 Cn+1 = 2 * (2 * n + 1) / (n + 2) * Cn
*/
func numTrees(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}
