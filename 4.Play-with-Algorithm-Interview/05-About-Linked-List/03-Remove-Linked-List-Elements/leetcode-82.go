package leetcode

// Time Complexity: O(n), Space Complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	// 如果整个链表的值都是重复的，也可以方便的返回空链表
	dummy := &ListNode{Next: head} // 统一链表的处理
	pre, cur := dummy, dummy.Next
	for cur != nil { // 只要当前节点的值和下一个节点的值相同
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next // 就一直移动到最后一个相同的数上
		}
		// update pre
		if pre.Next != cur { // 跨过相同的数字
			pre.Next = cur.Next
		} else { // 没有相同的数字就移动一位
			pre = pre.Next
		}
		// update cur
		cur = pre.Next
	}
	return dummy.Next
}
