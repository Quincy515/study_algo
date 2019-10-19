package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements1(head *ListNode, val int) *ListNode {
	dummy := new(ListNode) // 新建一个虚拟头节点
	dummy.Next = head      // 虚拟头节点指向head
	notEqual := dummy      // 新建一个游标
	for notEqual.Next != nil {
		if notEqual.Next.Val == val {
			notEqual.Next = notEqual.Next.Next
		} else {
			notEqual = notEqual.Next
		}
	}
	return dummy.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{} // 新头节点
	newHead.Next = head
	cur := newHead // 游标
	for p := head; p != nil; p = p.Next {
		if p.Val == val {
			cur.Next = p.Next
		} else {
			cur = p
		}
	}
	return newHead.Next
}
