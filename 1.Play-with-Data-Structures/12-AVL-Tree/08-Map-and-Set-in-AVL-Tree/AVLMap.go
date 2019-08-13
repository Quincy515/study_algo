package AVLMap

type AVLMap struct {
	avl AVLTree
}

func NewAVLMap() *AVLMap {
	return &AVLMap{}
}

func (am *AVLMap) GetSize() int {
	return am.avl.GetSize()
}

func (am *AVLMap) IsEmpty() bool {
	return am.avl.IsEmpty()
}

func (am *AVLMap) Add(k interface{}, v interface{}) {
	am.avl.Add(k, v)
}

func (am *AVLMap) Contains(key interface{}) bool {
	return am.avl.Contains(key)
}

func (am *AVLMap) Get(key interface{}) interface{} {
	return am.avl.Get(key)
}

func (am *AVLMap) Set(key, newValue interface{}) {
	am.avl.Set(key, newValue)
}

func (am *AVLMap) Remove(key interface{}) interface{} {
	return am.avl.Remove(key)
}
