package leetcode

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

// Time: O(n) Space: O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	// 处理边界情况，链表为空，或只有一个节点，或k<=0
	if head == nil || head.Next == nil || k <= 0 {
		return head // 直接返回原链表
	}
	end := head           // 否者定义指向链表结束节点的指针end
	n := 1                // 定义链表长度n
	for end.Next != nil { // 遍历一次链表计算链表长度
		n++
		end = end.Next
	}
	// 然后让end的Next指向头节点，形成一个环形链表
	end.Next = head

	k %= n         // 把k更新为对n取模的结果
	newEnd := head // 定义newEnd指针
	for i := 0; i < n-k-1; i++ {
		newEnd = newEnd.Next // 遍历n-k-1后，指向新的结束节点
	}
	newHead := newEnd.Next // newEnd的下一个节点，作为链表新的头节点
	newEnd.Next = nil      // newEnd的Next指针置空nil
	return newHead         // 最后返回新的头节点
}
