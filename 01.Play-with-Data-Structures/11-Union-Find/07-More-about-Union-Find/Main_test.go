package UnionFind6

import (
	"fmt"
	UnionFind3 "github.com/custergo/study_algo/01.Play-with-Data-Structures/11-Union-Find/04-Optimized-by-Size"
	UnionFind4 "github.com/custergo/study_algo/01.Play-with-Data-Structures/11-Union-Find/05-Optimized-by-Rank"
	UnionFind5 "github.com/custergo/study_algo/01.Play-with-Data-Structures/11-Union-Find/06-Path-Compression"
	"math/rand"
	"testing"
	"time"
)

func testUF(uf UF, m int) time.Duration {
	size := uf.GetSize()
	rand.Seed(time.Now().Unix())

	startTime := time.Now()
	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.UnionElements(a, b)
	}

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.IsConnected(a, b)
	}
	return time.Now().Sub(startTime)
}

func TestUnionFind(t *testing.T) {
	size := 10000000
	m := 10000000

	//uf1 := UnionFind1.NewUnionFind(size)
	//fmt.Println("UnionFind1: ", testUF(uf1, m)) // 105.206651ms
	//
	//uf2 := UnionFind2.NewUnionFind2(size)
	//fmt.Println("UnionFind2: ", testUF(uf2, m)) // 1.273354105s

	uf3 := UnionFind3.NewUnionFind(size)
	fmt.Println("UnionFind3: ", testUF(uf3, m)) // 12.956873ms

	uf4 := UnionFind4.NewUnionFind(size)
	fmt.Println("UnionFind4: ", testUF(uf4, m)) // 12.956873ms

	uf5 := UnionFind5.NewUnionFind(size)
	fmt.Println("UnionFind4: ", testUF(uf5, m)) // 路径压缩非递归

	uf6 := NewUnionFind(size)
	fmt.Println("UnionFind4: ", testUF(uf6, m)) // 路径压缩递归方式

}
