package leetcode

// Time Complexity: O(n) Space Complexity: O(1)
func partition(head *ListNode, x int) *ListNode {
	if head == nil { // 处理边界条件，head为空返回nil
		return nil
	}
	smaller := &ListNode{} // 小于给定值的节点链接在smaller后面
	greater := &ListNode{} // 等于等于给定值的节点链接在greater后
	s := smaller
	g := greater // 定义指针在链表中使用
	// 遍历原始链表
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x { // 如果当前节点值小于给定值
			s.Next = cur // 就链接到smaller后
			s = s.Next   // 并移动相应指针
		} else {
			g.Next = cur
			g = g.Next
		}
	}
	s.Next = greater.Next // 把两个链表连接起来
	g.Next = nil          // 并把链表最后节点指向置空
	return smaller.Next
}
