package _1_Weighted_Graph

import "reflect"

// 图的接口
type Graph interface {
	V() int
	E() int
	AddEdge(v, w int, weight weight)
	Show()
}

func assert(exp bool) {
	if !exp { // 表示式为false
		panic("断言失败,发生错误!")
	}
}

type weight interface{}

func compare(a, b weight) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		panic("cannot compare different type params")
	}
	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}
