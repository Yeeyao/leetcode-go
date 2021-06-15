package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	给定一个链表，需要移除特定数值的节点，最后返回新的链表头
	可能不存在节点
	首先使用一个新的节点，将 head 赋值过去 然后需要判断每个节点
	这里需要记录下当前的节点，以及当前节点的上一个节点，因为当前节点可能会被移除，所以上一个节点就需要指向下一个位置
	prev, cur 分别表示上一个节点和当前节点，这两个怎么更新呢？
		正常的话，prev = cur cur = cur.Next
		如果 cur.Val == val 当前 cur 需要跳过
			cur = cur.Next.Next 不对的，这样是跳过了下一个，而不是当前这个
			需要跳过当前这个，则 prev.Next = cur.Next
	判断当前节点数值，每次判断结束将当前节点指向下一个节点来判断下一个节点
		如果是需要删除的数值，prev.Next = cur.Next 跳过该节点来实现删除

	新的头节点，将它的 next 指向 head，这样，不管 head 是不是 nil 都可以返回到新的节点，这是一般化的处理
	这里的问题还是，head 本身如果数值是需要删除的，它本身也不会从链表移除，所以需要有机制可以将它从链表删除
	这样就不能从 head 开始判断，而是需要从上一个节点开始？这样如果 head 需要删除，就可以跳过 head

	初始化，cur 指向头节点，prev 应该是空的节点
	循环处理，这里 cur 和 prev 都会变化，所以新的 head 就不能使用这两个的信息
		cur 一直向后是合理的
		prev 则根据 cur 数值进行处理

看了一下题解，不需要用 pred，所以不要上来就思维定势，就是自己想得太复杂了

[ref](https://leetcode.com/problems/remove-linked-list-elements/discuss/57308/Concise-C%2B%2B-solution-with-pseudo-ListHead)
*/
func removeElements(head *ListNode, val int) *ListNode {
	pHead := &ListNode{}
	pHead.Next = head
	// cur 指向 pHead
	cur := pHead
	for cur != nil {
		// 下一个的数值等于 val，当前的下一个节点就跳过下一个
		if cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return pHead.Next
}
