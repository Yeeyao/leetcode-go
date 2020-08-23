package tree

/*
Diameter of Binary Tree
Given a binary tree, you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

任意两个节点的最长路径，返回边的数量，其实就是求最后一层的节点数量，如果大于 1 就是任意两个节点的
*/
/*
一条路径的长度为该路径经过的节点数减一，所以求直径（即求路径长度的最大值）等效于求路径经过节点数的最大值减一。
以该节点为起点的路径经过节点数的最大值即为 L+R+1
	左儿子向下遍历经过最多的节点数 L 右儿子向下遍历经过最多的节点数 R

其实就是计算当前节点的左右子树深度之和 + 1，最后返回的是每个节点的这个结果的最大值 -1
	在深度计算过程中更新最大值
*/

/*
	题目所求路径是节点的左右子树节点数 + 1 然后 - 1 得到结果
	先初始化 res = 1
	递归调用函数后返回 res - 1
	然后递归计算高度
		如果当前节点为 nil 返回 0
		否则，递归调用得到左右节点的高度
		更新 res
		返回高度
*/
*/
var res int
func diameterOfBinaryTree(root *TreeNode) int {
	res = 1
	helper(root)
	return res - 1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	// 注意这里先计算了最大值，再计算高度
	if left + right + 1 > res {
		res = left + right + 1
	}
	if left > right {
		return left + 1
	} else {
		return  right + 1
	}
}