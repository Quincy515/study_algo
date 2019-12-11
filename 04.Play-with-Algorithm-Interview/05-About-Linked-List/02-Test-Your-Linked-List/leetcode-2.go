package leetcode

// Time Complexity: O(max(l1,l2)), Space Complexity: O(max(l1,l2)) max(l1,l2)链表最大的长度
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode) // 定义虚拟头节点
	p := dummy             // 游标
	carry := 0             // 进位变量
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry // sum 初始化为 carray，首先计算进位
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 以sum对10取模的结果来创建新的节点
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next       // p向后移动一个节点，指向新创建的当前节点
		carry = sum / 10 // 更新进位
	}
	return dummy.Next // 返回结果链表的头节点
}
