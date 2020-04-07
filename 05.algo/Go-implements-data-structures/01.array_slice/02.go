package main

import "fmt"

func main() {
	// 数组的声明
	// 第一种，在初始化时只声明数组长度，不声明数组内容
	var arr1 [5]int
	// 第二种，知道数组很多，不想自己写长度的时候可以用这种方法
	// 声明之后由编译器自己推算数组长度
	// arr2 := [...]int{1, 3, 5, 7, 9}
	// 第三种，声明的时候长度和初值一起声明
	// arr3 := [3]int{2, 4, 6}
	// 二维数组的声明，其意义是三行五列
	var Block [3][5]bool

	// 数组的遍历
	// 第一种
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d\n", arr1[i])
	}
	// 第二种
	for index, value := range arr1 {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}

	// 以第二种方式遍历二维数组，只取值，也就是取出一个数组
	for _, v := range Block {
		// 再对这个数组取值
		for _, value := range v {
			fmt.Printf("%v ", value)
		}
		fmt.Printf("\n")
	}
}
