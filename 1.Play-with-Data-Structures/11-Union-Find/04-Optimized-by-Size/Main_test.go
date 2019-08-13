package UnionFind3

import (
	"fmt"
	UnionFind1 "github.com/custergo/study_algo/Play-with-Data-Structures/11-Union-Find/02-Quick-Find"
	UnionFind2 "github.com/custergo/study_algo/Play-with-Data-Structures/11-Union-Find/03-Quick-Union"
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
	size := 10000
	m := 100000

	uf1 := UnionFind1.NewUnionFind(size)
	fmt.Println("UnionFind1: ", testUF(uf1, m)) // 105.206651ms

	uf2 := UnionFind2.NewUnionFind2(size)
	fmt.Println("UnionFind2: ", testUF(uf2, m)) // 1.273354105s

	uf3 := NewUnionFind(size)
	fmt.Println("UnionFind3: ", testUF(uf3, m)) // 12.956873ms
}
