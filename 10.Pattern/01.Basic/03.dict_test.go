package main

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	// 创建
	m := make(map[string]int)
	// 设置 key value
	m["hello"] = 1
	// 删除 key
	delete(m, "hello")
	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 在访问的key不存在时，仍会返回默认零值，不能通过返回nil来判断元素是否存在
	if value, ok := m["hello"]; !ok {
		fmt.Println(value)
	}
}

/**
注意点
- map 键需要可比较的，不能为 slice、map、function，struct 类型不包含上述字段，也可以作为 key
- map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由0变为1
- 比较两个 map 需要遍历，其中的 key value 是否相同，因为有默认值关系，所以需要检查 value 和 ok 两个值
*/

/**
Map与工厂模式
• Map的value可以是一个方法
• 与Go的Dock type接口方式一起，可以方便的实现单一方法对象的工厂模式
*/

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

/**
Go的内置集合中没有Set实现，可以map[type]bool
1. 元素的唯一性
2. 基本操作
	a. 添加元素
	b. 判断元素是否存在
	c. 删除元素
	d. 元素个数
*/

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] { // 判断一个元素是否存在
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	// 添加元素要设置为 true
	mySet[3] = true
	t.Log(len(mySet)) // 此时set中有 1 和 3
	delete(mySet, 1) // 删除 set 中的 1
	n = 1
	if mySet[n] { // 判断 1 是否存在
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet)) // 返回此时元素个数
}
