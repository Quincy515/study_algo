package AVLMap

type AVLSet struct {
	avl AVLTree
}

func NewAVLSet() *AVLSet {
	return &AVLSet{}
}

func (this *AVLSet) GetSize() int {
	return this.avl.GetSize()
}

func (this *AVLSet) IsEmpty() bool {
	return this.avl.IsEmpty()
}

func (this *AVLSet) Add(e interface{}) {
	this.avl.Add(e, nil)
}

func (this *AVLSet) Contains(e interface{}) bool {
	return this.avl.Contains(e)
}

func (this *AVLSet) Remove(e interface{}) {
	this.avl.Remove(e)
}
