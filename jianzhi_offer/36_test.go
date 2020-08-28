package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("36  两个链表的第一个公共结点 ", func(t *testing.T) {
		nums := []int{7, 5, 6, 4}
		get := solution(nums)
		want := 5
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*

 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
[ref](https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/solution/shuang-zhi-zhen-fa-lang-man-xiang-yu-by-ml-zimingm/)
	两个链表的第一个公共节点 值相等，后节点相同
	brute force O(nm) 一个个比较
	双指针 两个都分别逐个遍历，到达尾部重新从头部开始遍历 如果相遇就是第一个公共节点
	类似小学的追逐问题 leetcode 142
	O(m + n) 时间
*/
func solution(headA, headB *ListNode) *ListNode {
	headAP, headBP := headA, headB
	for headAP != headBP {
		if headAP != nil {
			headAP = headAP.Next
		} else {
			headAP = headB
		}
		if headBP != nil {
			headBP = headBP.Next
		} else {
			headBP = headA
		}
	}
	return headAP
}
