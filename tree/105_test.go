package tree

/*
[ref](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9/)
105. Construct Binary Tree from Preorder and Inorder Traversal
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
    3
   / \
  9  20
    /  \
   15   7

通过树的前序和中序遍历重建二叉树
前序遍历中，第一个元素是根节点 中序遍历则是左子树，根，右子树
对具体节点两种遍历进行映射
针对前序遍历的第一个根节点，在中序遍历中找到它的位置，其左边是该节点的左子树节点，右边是右子树节点

这里先判断先序遍历列表为空直接返回 nil
用先序遍历的第一个节点初始化 root，然后找到该节点在中序遍历中的索引位置
root 的左右节点就使用递归处理，这里变化的是两个 slice 都是根据 inorder 长度变化
	跳过当前的根节点
最后返回 root
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

/*
	另外一个递归的版本
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderNumToIndex := make(map[int]int)
	for i, num := range inorder {
		inorderNumToIndex[num] = i
	}
	preorderIndex := 0
	return build(preorder, inorder, &preorderIndex, 0, len(inorder)-1, inorderNumToIndex)
}

func build(preorder, inorder []int, preorderIndex *int, inorderStart, inorderEnd int, inorderNumToIndex map[int]int) *TreeNode {
	if *preorderIndex == len(inorder) || inorderStart > inorderEnd {
		return nil
	}
	curNum := preorder[*preorderIndex]
	*preorderIndex += 1
	node := &TreeNode{
		Val: curNum,
	}
	inorderIndex := inorderNumToIndex[curNum]
	node.Left = build(preorder, inorder, preorderIndex, inorderStart, inorderIndex-1, inorderNumToIndex)
	node.Right = build(preorder, inorder, preorderIndex, inorderIndex+1, inorderEnd, inorderNumToIndex)
	return node
}
