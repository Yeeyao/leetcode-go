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

	先序遍历树 AAA 中的每个节点 nAn_AnA​ ；（对应函数 isSubStructure(A, B)）
	判断树 AAA 中 以 nAn_AnA​ 为根节点的子树 是否包含树 BBB 。（对应函数 recur(A, B)）

isSubStructure(A, B) 函数：

    特例处理： 当 树 AAA 为空 或 树 BBB 为空 时，直接返回 falsefalsefalse ；
    返回值： 若树 BBB 是树 AAA 的子结构，则必满足以下三种情况之一，因此用或 || 连接；
        以 节点 AAA 为根节点的子树 包含树 BBB ，对应 recur(A, B)；
        树 BBB 是 树 AAA 左子树 的子结构，对应 isSubStructure(A.left, B)；
        树 BBB 是 树 AAA 右子树 的子结构，对应 isSubStructure(A.right, B)；
终止条件：
    当节点 BBB 为空：说明树 BBB 已匹配完成（越过叶子节点），因此返回 true；
    当节点 AAA 为空：说明已经越过树 AAA 叶子节点，即匹配失败，返回 false；
    当节点 AAA 和 BBB 的值不同：说明匹配失败，返回 false；
返回值：
    判断 AAA 和 BBB 的左子节点是否相等，即 recur(A.left, B.left) ；
    判断 AAA 和 BBB 的右子节点是否相等，即 recur(A.right, B.right) ；
*/
func solution(A, B *TreeNode) bool {
	return !(A == nil || B == nil) &&
		// 这里类似先序遍历，先根节点作为开始的比较节点，然后左右两个字节点并递归下去
		// 同时注意这里的或的判断
		(solutionHelper(A, B) || solution(A.Left, B) || solution(A.Right, B))
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
