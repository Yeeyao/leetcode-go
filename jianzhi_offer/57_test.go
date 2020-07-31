package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("57 二叉树的下一个结点  ", func(t *testing.T) {
		root := Node{1, nil}
		get := solution(&root)
		want := nil
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

/*
	给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。
	注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。

	给出给定节点的中序遍历的下一个节点 左中右
	如果节点有右子树，则需要遍历右子树，如果没有，则需要向上遍历
	节点的形式
		单独的左边叶子节点，下一个节点是父节点
		单独的右边叶子节点，下一个节点是父节点的父节点
		中间节点，有右子树，右子树的左子树的最后一个子树的左节点或者左子树的根
		中间节点，没有右子树
			如果是根节点的左子树部分，下一个节点是根节点
			如果是根节点的右子树部分，下一个节点是根节点的父节点
*/
func solution(node *Node) *Node {
	// 当前节点有右子树，直接找右子树的最左子节点
	if node.Left == nil && node.Right != nil {
		res := node.Right
		for res.Left != nil {
			res = res.Left
		}
		return res
	} else {
		/*
			当前节点没有右子树
			如果当前节点是根节点的左子树部分，返回根节点
			如果是右子树部分，则需要向上找根节点的左边部分，这里需要循环向上找左子树的根
		*/

		for node.Parent != nil {
			if node.Parent.Left == node {
				return node.Parent
			}
			node = node.Parent
		}
	}
	return nil
}
