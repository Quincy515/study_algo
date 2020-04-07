package main

import "fmt"

func main() {
	// 第一种，在初始化时只声明数组长度，不声明数组内容
	var arr1 [5]int
	// 第二种，知道数组很多，不想自己写长度的时候可以用这种方式
	// 声明之后由编译器自己推算数组长度
	arr2 := [...]int{1, 3, 5, 7, 9}
	// 第三种，声明的时候长度和初值一起声明
	arr3 := [3]int{2, 4, 6}
	printArr(arr1)
	printArr(arr2)
	printArr(arr3)
}

func printArr(arr [5]int) {
	// 第二种 数组的遍历
	for index, value := range arr {
		fmt.Printf("索引：%的， 值：%d\n", index, value)
	}
}