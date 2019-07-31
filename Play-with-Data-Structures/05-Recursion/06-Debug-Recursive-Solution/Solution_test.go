package _6_Debug_Recursive_Solution

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := NewListNode(nums)
	fmt.Println(head)

	res := removeElements(head, 6, 0)
	fmt.Println(res)
}

/**
=== RUN   TestListNode
1->2->6->3->4->5->6->NULL
Call: remove  6  in  1->2->6->3->4->5->6->NULL
--Call: remove  6  in  2->6->3->4->5->6->NULL
----Call: remove  6  in  6->3->4->5->6->NULL
------Call: remove  6  in  3->4->5->6->NULL
--------Call: remove  6  in  4->5->6->NULL
----------Call: remove  6  in  5->6->NULL
------------Call: remove  6  in  6->NULL
--------------Call: remove  6  in  NULL
--------------Return:  NULL
------------After remove  6 :  NULL
------------Return:  NULL
----------After remove  6 :  NULL
----------Return:  5->NULL
--------After remove  6 :  5->NULL
--------Return:  4->5->NULL
------After remove  6 :  4->5->NULL
------Return:  3->4->5->NULL
----After remove  6 :  3->4->5->NULL
----Return:  3->4->5->NULL
--After remove  6 :  3->4->5->NULL
--Return:  2->3->4->5->NULL
After remove  6 :  2->3->4->5->NULL
Return:  1->2->3->4->5->NULL
1->2->3->4->5->NULL
--- PASS: TestListNode (0.00s)
PASS
*/
