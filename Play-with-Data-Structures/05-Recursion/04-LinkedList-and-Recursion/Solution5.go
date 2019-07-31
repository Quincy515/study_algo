package _4_LinkedList_and_Recursion

// 优化代码
func removeElements5(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements5(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
