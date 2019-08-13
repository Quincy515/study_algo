package _9_Non_Recursion_Preorder_Traverse_in_BST

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
	fmt.Println()

	bst.PreOrderNR()

}
