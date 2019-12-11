package main

/**
思路：开一个栈存放链表前半段
时间复杂度: O(N)
空间复杂度: O(N)
*/

func isPalindrome1(l *LinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return false
	}
	s := make([]string, 0, lLen/2)
	cur := l.head
	for i := uint(1); i <= lLen; i++ {
		cur = cur.next
		// 如果链表有奇数个结点，中间的直接忽略
		if lLen%2 != 0 && i == (lLen/2+1) {
			continue
		}
		// 前一半入栈
		if i <= lLen/2 {
			s = append(s, cur.GetValue().(string))
		} else {
			// 后一半与前一半进行对比
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}
	return true
}

/**
思路：找到链表中间结点，将前半部分转置，再从中间向左右遍历对比
时间复杂度: O(N)
空间复杂度: O(N)
*/
func isPalindrome2(l *LinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	var isPalindrome = true
	step := lLen / 2
	var pre *ListNode = nil
	cur := l.head.next
	next := l.head.next.next
	for i := uint(1); i <= step; i++ {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.next
	}
	mid := cur

	var left, right *ListNode = pre, nil
	if lLen%2 != 0 {
		right = mid.next
	} else {
		right = mid
	}

	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.next
		right = right.next
	}

	// 复原链表
	cur = pre
	pre = mid
	for nil != cur {
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre
	return isPalindrome
}
