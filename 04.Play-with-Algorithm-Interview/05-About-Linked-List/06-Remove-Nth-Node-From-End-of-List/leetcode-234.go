package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用辅助栈 Time: O(n), Space: O(n)
func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)               // 定义一个存储整数的栈
	for p := head; p != nil; p = p.Next { // 遍历单链表
		stack = append(stack, p.Val) // 把数字入栈
	}

	for p := head; p != nil; p = p.Next { // 遍历单链表
		// 对比链表当前节点的数和栈顶的数
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.Val != val { // 如果不相等
			return false // 则返回false
		}
	}
	return true // 循环结束后没有返回false则返回true
}

// 翻转一半链表的方法 Time: O(n), Space: O(1)
func isPalindrome(head *ListNode) bool {
	length := 0 // 定义链表长度为0
	for p := head; p != nil; p = p.Next {
		length++ // 遍历链表，当p不为空时长度+1
	}

	cur := head       // 定义指向当前节点的cur
	var pre *ListNode // 定义前一节点的pre
	for i := 0; i < length/2; i++ {
		next := cur.Next // 先存储cur的下一节点
		cur.Next = pre   // cur指向它的前一节点pre
		pre = cur        // 更新cur和pre的指向，pre移动到cur的位置
		cur = next       // cur移动到next的位置
	}

	if length%2 == 1 { // 判断链表长度是否为奇数
		cur = cur.Next // 是的话，cur还要向后再移动一位
	}
	// pre和cur各自向两边移动
	for ; pre != nil && cur != nil; pre, cur = pre.Next, cur.Next {
		if pre.Val != cur.Val { // 当它们指向的节点不相等时
			return false // 直接返回false,说明并不是回文链表
		}
	}
	return true // 循环结束后还没返回，则最后返回true
}
