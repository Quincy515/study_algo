package Prim

import (
	"fmt"
	"testing"
	"time"
)

// 测试实现的两种Prim算法的性能差距
// 可以看出使用索引堆实现的Prim算法优于 Lazy Prim 算法
func TestGraphPerformance(t *testing.T) {
	fmt.Println("===读取 testG1.txt 生成 Weighted Sparse Graph===")
	filename1 := "testG1.txt"
	filename2 := "testG2.txt"
	filename3 := "testG3.txt"
	filename4 := "testG4.txt"
	v1, v2, v3, v4 := 8, 250, 1000, 10000
	readGraph := &ReadGraph{}

	// 文件读取
	g1 := NewSparseGraph(v1, false)
	if err := readGraph.Read(g1, filename1); err != nil {
		fmt.Println(err)
		return
	}
	//g1.Show()
	fmt.Println("testG1.txt load successfully.")

	g2 := NewSparseGraph(v2, false)
	if err := readGraph.Read(g2, filename2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("testG2.txt load successfully.")

	g3 := NewSparseGraph(v3, false)
	if err := readGraph.Read(g3, filename3); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("testG3.txt load successfully.")

	g4 := NewSparseGraph(v4, false)
	if err := readGraph.Read(g4, filename4); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("testG4.txt load successfully.")

	// Test Lazy Prim MST
	fmt.Println("Test Lazy Prim MST: ")
	startTime := time.Now()
	NewLazyPrimMST(g1)
	fmt.Println("Test for G1: ", time.Now().Sub(startTime))
	startTime = time.Now()
	NewLazyPrimMST(g2)
	fmt.Println("Test for G2: ", time.Now().Sub(startTime))
	startTime = time.Now()
	NewLazyPrimMST(g3)
	fmt.Println("Test for G3: ", time.Now().Sub(startTime))
	startTime = time.Now()
	NewLazyPrimMST(g4)
	fmt.Println("Test for G4: ", time.Now().Sub(startTime))

	// Test Prim MST
	fmt.Println("Test Prim MST: ")
	startTime = time.Now()
	NewPrimMST(g1)
	fmt.Println("The for G1: ", time.Now().Sub(startTime))
	startTime = time.Now()
	NewPrimMST(g2)
	fmt.Println("The for G2: ", time.Now().Sub(startTime))
	startTime = time.Now()
	NewPrimMST(g3)
	fmt.Println("The for G3: ", time.Now().Sub(startTime))
	startTime = time.Now()
	NewPrimMST(g4)
	fmt.Println("The for G4: ", time.Now().Sub(startTime))
}
