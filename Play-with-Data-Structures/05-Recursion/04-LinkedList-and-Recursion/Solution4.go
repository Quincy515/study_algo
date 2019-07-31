package _4_LinkedList_and_Recursion

// 递归实现
func removeElements4(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	res := removeElements4(head.Next, val)
	if head.Val == val {
		return res
	} else {
		head.Next = res
		return head
	}
}
