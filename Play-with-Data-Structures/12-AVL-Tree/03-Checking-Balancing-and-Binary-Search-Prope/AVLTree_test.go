package AVLTree

import (
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	fmt.Println("pride-and-prejudice")
	words := ReadFile("pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words))

	customMap := NewAVLTree()
	for _, word := range words {
		if customMap.Contains(word) {
			customMap.Set(word, customMap.Get(word).(int)+1)
		} else {
			customMap.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", customMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", customMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", customMap.Get("prejudice"))

	fmt.Println("is BST: ", customMap.IsBST())           // 还是二分搜索树
	fmt.Println("is Balanced: ", customMap.IsBalanced()) // 不是平衡二叉树
}
