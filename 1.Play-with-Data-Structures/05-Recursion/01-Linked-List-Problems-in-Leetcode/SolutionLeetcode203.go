package _1_Linked_List_Problems_in_Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}
	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			cur := prev.Next
			prev.Next = cur.Next
			cur.Next = nil
		} else {
			prev = prev.Next
		}
	}

	return head
}

func removeElements1(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummyHead.Next
}
