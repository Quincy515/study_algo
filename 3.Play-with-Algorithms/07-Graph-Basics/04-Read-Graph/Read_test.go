package Graph_Representation

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
	fmt.Println("test G1 in Sparse Graph: ")
	g1.Show()

	g2 := NewDenseGraph(13, false)
	if err := readGraph.Read(g2, filename); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test G1 in Dense Graph: ")
	g2.Show()

	// 使用两种图的存储方式读取testG2.txt文件
	fmt.Println("=======使用两种图的存储方式读取testG2.txt文件======")
	filename = "testG2.txt"
	g3 := NewSparseGraph(6, false)

	if err := readGraph.Read(g3, filename); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test G2 in Sparse Graph: ")
	g3.Show()

	g4 := NewDenseGraph(6, false)
	if err := readGraph.Read(g4, filename); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test G2 in Dense Graph: ")
	g2.Show()
}
