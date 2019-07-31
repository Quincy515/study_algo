package _6_PreOrder_Traverse_in_BST

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := NewBST()
	nums := []int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}

	//////////////////
	//       5      //
	//     /  \     //
	// 	  3    6    //
	//   / \    \   //
	//  2  4    8   //
	//////////////////
	bst.PreOrder()
	fmt.Println(bst)
}

/**
5
--3
----2
------nil
------nil
----4
------nil
------nil
--6
----nil
----8
------nil
------nil
*/
