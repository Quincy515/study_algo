package Bellman_Ford

import (
	"reflect"
	"strconv"
)

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
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
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
	case *Edge:
		aTemp, err := strconv.ParseFloat(a.(*Edge).Wt().(string), 32)
		if err != nil {
			panic(err)
		}
		bTemp, err := strconv.ParseFloat(b.(*Edge).Wt().(string), 32)
		if err != nil {
			panic(err)
		}
		if aTemp > bTemp {
			return 1
		} else if aTemp < bTemp {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params: ")
	}
}

func add(a, b weight) float64 {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		panic("cannot add different type params")
	}
	switch a.(type) {
	case int:
		return float64(a.(int) + b.(int))
	case float64:
		return a.(float64) + b.(float64)
	default:
		panic("unsupported type params: ")
	}
}
