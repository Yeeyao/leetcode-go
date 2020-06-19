package jianzhi_offer

import (
	"testing"
)

func TestPro(t *testing.T) {
	t.Run("17 树的子结构", func(t *testing.T) {
		A := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
		B := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
		get := solution(A, B)
		want := true
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
	输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）
	B 是不是 A 的一部分
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	B 是 A 的子结构，则子结构的根节点可以是 A 的任意节点
		先序遍历 A 的每个节点 nA
		判断 A 中以 nA 为根节点的子树是否包含树 B

	最外层函数是用 A 的所有节点作为开始的节点跟 B 来比较
	里层是对每个节点进行继续遍历判断


	A 或者 B 是空则直接返回 false 外层循环函数递归调用 A 左右两个子节点
		这里返回是取或，只需要一边 true 就行
	内层循环的函数就进行节点的数值判断
		B 为空，表示完成匹配，返回 true;
		A 为空表示遍历完了，没有匹配，返回 false
		A,B的值 不同，返回 false
		然后递归判断 A，B 的对应左右子树
	返回值 判断 A B 的左右子树分别是否相等

*/
func solution(A, B *TreeNode) bool {
	return !(A == nil || B == nil) &&
		// 这里类似先序遍历，先根节点作为开始的比较节点，然后左右两个字节点并递归下去
		// 同时注意这里的或的判断
		(solutionHelper(A, B) || solutionHelper(A.Left, B) || solutionHelper(A.Right, B))
}

func solutionHelper(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	// 以当前为开始节点继续进行左右子树的遍历 这里作为开始，然后分别递归判断 A B 的左右子树，需要 与 判断
	return solutionHelper(A.Left, B.Left) && solutionHelper(A.Right, B.Right)
}
