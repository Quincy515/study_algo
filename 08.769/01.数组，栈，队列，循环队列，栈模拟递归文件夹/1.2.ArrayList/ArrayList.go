package ArrayList

// 接口
type List interface {
	Size() int                                 // 数组大小
	Get(index int) (interface{}, error)        // 获取索引位置的元素
	Set(index int, newValue interface{}) error // 修改索引位置的元素
	Insert(index int, item interface{}) error  // 在索引位置插入一个元素
	Append(item interface{}) error             // 追加一个元素
	Clear()                                    // 清空数组
	Delete(index int) error                    // 删除指定索引的元素
	String() string                            // 返回字符串
}

// 数据结构
type ArrayList struct {
	dataStore []interface{} // 存储的元素
	theSize   int           // 存储元素的个数
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 对数组开辟内存
	list.theSize = 0
	return list
}
