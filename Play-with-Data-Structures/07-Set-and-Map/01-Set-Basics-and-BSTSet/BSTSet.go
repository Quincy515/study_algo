package _1_Set_Basics_and_BSTSet

type BSTSet struct {
	bst *BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{NewBST()}
}

func (bs *BSTSet) GetSize() int {
	return bs.bst.Size()
}

func (bs *BSTSet) IsEmpty() bool {
	return bs.bst.IsEmpty()
}

func (bs *BSTSet) Add(e interface{}) {
	bs.bst.Add(e)
}

func (bs *BSTSet) Remove(e interface{}) {
	bs.bst.Remove(e)
}

func (bs *BSTSet) Contains(e interface{}) bool {
	return bs.bst.Contains(e)
}
