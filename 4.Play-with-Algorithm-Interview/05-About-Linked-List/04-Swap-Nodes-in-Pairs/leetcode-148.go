package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 单链表快速排序Time Complexity: O(nlogn), Space: O(n)
func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head // 返回排序后的链表
}

// 用于交换两个节点上的数值
func swap(a, b *ListNode) {
	a.Val, b.Val = b.Val, a.Val
}

// 快速排序的主体函数
func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return // 头节点已经是结束节点或要排序只有一个头节点直接返回
	}
	pivot := head.Val             // 头节点的值作为pivot
	slow, fast := head, head.Next // 定义快慢指针
	for fast != end {             // 当快指针不等于结束指针时，不断执行以下操作
		if fast.Val <= pivot { // 如果快指针小于pivot
			slow = slow.Next // 不断向后移动慢指针
			swap(slow, fast) // 交换快慢指针指向的节点值
		}
		fast = fast.Next // 快指针向后移动一位
	}
	// 循环结束后交换头节点和慢指针指向的值,把pivot放在大小两堆数值的中间
	swap(head, slow)
	quickSort(head, slow) // 递归处理pivot左右两边的链表
	quickSort(slow.Next, end)
}

// 单链表归并排序 Time: O(n*logn) Space: O(logn)
func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head // 快慢指针初始化在链表头
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next      // 慢指针一次走一步
		fast = fast.Next.Next // 快指针一次走两步
	}
	// 循环结束后慢指针指向的是前一半链表的最后一个元素
	right := mergeSortList(slow.Next)       // 先排序后一段链表
	slow.Next = nil                         // 然后把慢指针的Next置为空指针
	left := mergeSortList(head)             // 再排序前一段链表
	return mergeTwoSortedLists(left, right) // 合并排序后的两段链表
}

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}
