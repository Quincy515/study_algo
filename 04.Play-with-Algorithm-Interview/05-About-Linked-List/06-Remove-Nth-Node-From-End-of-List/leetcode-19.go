package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(k) k为链表长度 Space: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // 定义虚拟头结点，并指向链表头节点
	p, q := dummy, dummy           // 定义游标p，q初始都指向dummy
	for n > 0 && q.Next != nil {   // 当n大于0，并且q还没有到尾节点时
		q = q.Next // q不断向后移动
		n--
	}
	if n != 0 { // 移动结束后n不等于0，说明n大于链表长度
		return dummy.Next // 直接返回原链表
	}
	for q.Next != nil { // 只要q没有到尾节点
		p = p.Next // 不断移动p，q
		q = q.Next
	}
	p.Next = p.Next.Next // 循环结束后，跳过要删除的节点
	return dummy.Next    // 返回dummy.Next
}
