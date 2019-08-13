package LinkedListSet

type LinkedListSet struct {
	list *LinkedList
}

func NewLinkedListSet() *LinkedListSet {
	return &LinkedListSet{NewLinkedList()}
}

func (ls *LinkedListSet) GetSize() int {
	return ls.list.GetSize()
}

func (ls *LinkedListSet) IsEmpty() bool {
	return ls.list.IsEmpty()
}

func (ls *LinkedListSet) Contains(e interface{}) bool {
	return ls.list.Contains(e)
}

func (ls *LinkedListSet) Add(e interface{}) {
	if !ls.list.Contains(e) { // 集合不会添加重复元素
		ls.list.AddFirst(e) // 从链表头添加元素，O(1)时间复杂度
	}
}

func (ls *LinkedListSet) Remove(e interface{}) {
	ls.list.RemoveElement(e)
}
