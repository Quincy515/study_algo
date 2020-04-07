# 知识点

重点数组和切片的不同

- 数组的数据类型
- 数组的创建
- 数组的遍历
- Golang数组与切片的区别
- 切片的扩容规律

# 数组基本操作

## 数组的声明

Go语言中一个数组的声明方式主要有以下几种[File 01.go]()

```go
package main

func main() {
	// 第一种，在初始化时只声明数组长度，不声明数组内容
	var arr1 [5]int
	// 第二种，知道数组很多，不想自己写长度的时候可以用这种方法
	// 声明之后由编译器自己推算数组长度
	arr2 := [...]int{1,3,5,7,9}
	// 第三种，声明的时候长度和初值一起声明
	arr3 := [3]int{2,4,6}
	// 二维数组的声明，其意义是三行五列
	var Block [3][5]int
}
```

这里需要注意的是Go语言中数组的初始值如果不做声明的话默认是全部有初值的。

比如arr1这个数组虽然只声明了长度为5，但是Go的编译器也会把这5个元素全部初始化为0。

而对于`bool`值类型的数组，如果不做赋值操作，则初始值全为`false`。

## 数组的遍历

Go的数组遍历主要有以下两种方式。 [File 02.go]()

```go
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
```

接下来，在终端执行:

`go run 02.go`

结果如下:

```shell
> go run 02.go
0
0
0
0
0
索引: 0, 值: 0
索引: 1, 值: 0
索引: 2, 值: 0
索引: 3, 值: 0
索引: 4, 值: 0
false false false false false
false false false false false
false false false false false
```

其中Go语言官方更加提倡的是第二种以`range`的方式进行遍历，这样写会让代码更加优雅，而且绝对不会越界。

那么，如果我们只想要数组里的index不想要value时怎么range呢？

> 答：只需要 `for i := range arr` 就可以了。

如果只想要value不想要索引的时候怎么写呢？

> 答：`for _, value := range arr` 注意这里的下划线不能省略。

## 封装一个数组打印函数

现在封装一个用来打印数组的函数并对其进行测试。 [File 03.go]()

```go
func printArr(arr [5]int) {
	// 第二种 数组的遍历
	for index, value := range arr {
		fmt.Printf("索引：%的， 值：%d\n", index, value)
	}
}
```

分别将arr1、2、3传入打印，先猜测以下会发生什么结果呢？

```go 
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
```

结果是程序在打印arr3时抛出了如下异常

```shell
> go run 03.go
# command-line-arguments
.\03.go:15:10: cannot use arr3 (type [3]int) as type [5]int in argument to printArr
```

这个就要牵扯出一个概念了，Go语言中数组是值类型。

也就是说 [3]int, 和 [5]int 在Go中会认为是两个不同的数据类型。

同样地，你在 PrintArr 中改变数组中的值也不会改变原数组的值。

> 注意在使用数组时，既要数据类型统一又要长度统一才能传递。

所以在Go中一般不直接使用数组，而是使用切片。