package tree

/*
236. Lowest Common Ancestor of a Binary Tree
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T
that has both p and q as descendants (where we allow a node to be a descendant of itself).”
Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]

给定一棵二叉树，找到给定树中两个节点的最小共同祖先
找到第一个共同祖先，该祖先的上面节点直到根节点都是共同祖先，每个都遍历然后判断大小
节点自己可以是自己的祖先

直接对两个指定的元素从根开始找两次，然后将找到该元素之前的节点按顺序保存，最后保存节点本身
然后变成从两个节点列表找第一个公共节点
优化一下，先保存其中一个的所有祖先节点列表，然后在找另一个节点的祖先节点中找公共的
[ref](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2/)
这里先从根节点开始遍历整个二叉树，哈希表记录m每个节点的父节点指针
p 节点开始不断往祖先移动，记录下已经访问过的祖先节点
q 节点不断往祖先移动，如果祖先已经访问过，就是最近公共祖先
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	visited := make(map[int]bool)
	dfs(root, parent)
	// 自己也算
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func dfs(root *TreeNode, parent map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		parent[root.Left.Val] = root
		dfs(root.Left, parent)
	}
	if root.Right != nil {
		parent[root.Right.Val] = root
		dfs(root.Right, parent)
	}
}
