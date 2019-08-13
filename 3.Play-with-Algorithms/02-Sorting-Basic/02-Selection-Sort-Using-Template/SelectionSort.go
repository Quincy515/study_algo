package Selection_Sort

import (
	"reflect"
	"sort"
)

func SelectionSort(arr []interface{}, n int) {
	for i := 0; i < n-1; i++ {
		// 寻找[i,n)区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if compare(arr[j], arr[minIndex]) < 0 {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// slice 是引用类型
func swap(arr []interface{}, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func compare(a, b interface{}) int {
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
	default:
		panic("unsupported type params")
	}
}

func StudentsSort(s sort.Interface) {
	length := s.Len()
	for i := 0; i < length; i++ {
		// 寻找[i,n)区间里的最小值
		minIndex := i
		for j := i + 1; j < length; j++ {
			if s.Less(j, i) {
				minIndex = j // j < i
			}
		}
		s.Swap(minIndex, i)
	}
}
