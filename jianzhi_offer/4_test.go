package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("4 重建二叉树", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		value := 4
		want := true
		got := solution(nums, value)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

/*
[LeetCode 105](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
前序遍历 根 左子树 右子树 中序遍历 左子树 根 右子树 后序遍历 左子树 右子树 根
			1
		2      3
	4       5    6
		7      8

前序遍历第一个节点是根节点，只要找到根节点在中序遍历的位置，在根节点前遍历的是左子树的节点，在根节点后遍历的是右子树的节点。
由此可知左右子树有多少节点
由于树中节点的数量和遍历方式无关，通过中序遍历得到左右子树节点数量之后，可以根据节点数量得到前序遍历的左右子树分界，因此可以
进一步得到左右子树各自的前序和中序遍历，可以使用递归方式，重建左右子树最后整个二叉树

递归
因为这里的元素都是唯一的，所有可以通过保存 map 来根据元素数值获得索引信息
使用一个 map 存储中序遍历每个元素以及其下标。调用递归方法，对于前序遍历和中序遍历，下标从 0 到 n - 1
递归方法：判断前序遍历下标范围，若开始大于结束，则当前二叉树没有节点，返回 null。若开始等于结束，当前二叉树只有一个根节点
若开始小于结束，二叉树有多个节点，在中序遍历中得到根节点索引，从而得到左右子树各自下标范围和数量，知道数量后，前序遍历中可以得到
左右子树各自下标范围，然后递归重建左右子树，将当前根节点作为左右子树根节点
*/

/*
	如果 preOrder 长度为 0 直接返回 nil
	创建一个 map 并将中序遍历的每个元素保存 key 是 inorder[i], value 是 i
	初始化 root 然后等于一个回调递归函数，最后返回 root
*/

type TreeNode struct {
	val         int
	left, right TreeNode
}

func solution(preOrder, inOrder []int) TreeNode {
	preOrderLen := len(preOrder)
	if preOrder == nil || preOrderLen == 0 {
		return nil
	}
	// 中序遍历的结果保存到 map
	indexMap := make(map[int]int)
	for i := 0; i < preOrderLen; i++ {
		indexMap[inOrder[i]] = i
	}
	root := buildTree(preOrder, inOrder, 0, preOrderLen-1, 0, preOrderLen-1, indexMap)
	return root
}

/*
因为这里遍历 preOrder 所以 poStart > poEnd 就直接返回 nil poStart == poEnd 就直接返回 root
根据 preOrder 当前元素值去 inOrder 中找到对应的索引，然后将根索引值减去当前中序索引的开始值表示当前的根左子树的节点数量，
将中序索引的结束值减去根索引值表示当前根右子树的节点数量，然后递归调用生成左右子树，最后将左右子树赋值给根的左右子树，返回 root
*/
func buildTree(preOrder, inOrder []int, poStart, poEnd, ioStart, ioEnd int, indexMap map[int]int) TreeNode {
	if poStart > poEnd {
		return nil
	}
	rootVal := preOrder[poStart]
	root := TreeNode{rootVal, nil, nil}
	if poStart == poEnd {
		return root
	} else {
		rootIoIndex := indexMap[rootVal]
		leftNodeNum := rootIoIndex - ioStart
		rightNodeNum := ioEnd - rootIoIndex
		// 注意这里的元素索引变化 posOrder 每次遍历了一个就需要跳到下一个，同时左右的边界处理，
		// inOrder 每次遍历处理了一个根，就需要跳过并递归处理左右
		// TODO: circle one: rii = 3 pos = 0, poe = 7, lnn = 3, rnn = 4 lst bt(p, i, 1, 3, 0, 2, im) rst bt(p, i, 4, 7, 4, 7, im)
		// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
		leftSubTree := buildTree(preOrder, inOrder, poStart + 1, poStart + leftNodeNum, ioStart, rootIoIndex - 1, indexMap)
		rightSubTree := buildTree(preOrder, inOrder, poEnd - rightNodeNum + 1, poEnd, rootIoIndex + 1, ioEnd, indexMap)
		root.left = leftSubTree
		root.right = rightSubTree
		return root
	}
}

/*
	迭代方法，使用栈 TODO:
preorder = [3,9,8,5,4,10,20,15,7]
inorder = [4,5,8,10,9,3,15,20,7]
 */
