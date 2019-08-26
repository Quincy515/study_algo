package _5_DFS_and_Components

import (
	"fmt"
	"testing"
)

func TestReadGraph(t *testing.T) {
	// 使用两种图的存储方式读取testG1.txt文件
	fmt.Println("=======使用两种图的存储方式读取testG1.txt文件======")
	filename := "testG1.txt"
	g1 := NewSparseGraph(13, false)
	readGraph := &ReadGraph{}

	if err := readGraph.Read(g1, filename); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("TestG1.txt, Using Sparse Graph, Component Count: ", NewComponents(g1).Count())
	//g1.Show()

	//g2 := NewDenseGraph(13, false)
	//if err := readGraph.Read(g2, filename); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("TestG1.txt, Using Dense Graph, Component Count: ", NewComponents(g2).Count())
	//g2.Show()

	// 使用两种图的存储方式读取testG2.txt文件
	fmt.Println("=======使用两种图的存储方式读取testG2.txt文件======")
	filename = "testG2.txt"
	g3 := NewSparseGraph(7, false)

	if err := readGraph.Read(g3, filename); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("TestG2.txt, Using Sparse Graph, Component Count: ", NewComponents(g3).Count())
	//g3.Show()

	//g4 := NewDenseGraph(7, false)
	//if err := readGraph.Read(g4, filename); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("TestG3.txt, Using Dense Graph, Component Count: ", NewComponents(g4).Count())
	//g2.Show()
}
