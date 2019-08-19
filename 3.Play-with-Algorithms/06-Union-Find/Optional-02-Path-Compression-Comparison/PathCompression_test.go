package Path_Compression

import (
	"fmt"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Optional-02-Path-Compression-Comparison/UF5"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Optional-02-Path-Compression-Comparison/UF6"
	"testing"
)

func TestUnionFind(t *testing.T) {
	// 为了能够方便地看出两种路径压缩算法的不同,只使用有5个元素的并查集进行试验
	n := 5

	uf5 := UF5.NewUnionFind(n)
	uf6 := UF6.NewUnionFind(n)

	// 我们将我们的并查集初始设置成如下的样子
	//            0
	//           /
	//          1
	//         /
	//        2
	//       /
	//      3
	//     /
	//    4
	for i := 1; i < n; i++ {
		uf5.Parent[i] = i - 1
		uf6.Parent[i] = i - 1
	}

	// 现在,对两个并查集调用find(4)操作
	uf5.Find(n - 1)
	uf6.Find(n - 1)

	// 通过show,可以看出,使用迭代的路径压缩,并查集变成这个样子:
	//     0
	//    / \
	//   1   2
	//  / \
	// 3  4
	fmt.Println("UF implements Path Compression without recursion: ")
	uf5.Show()
	fmt.Println()

	// 使用递归的路径压缩, 我们的并查集变成这个样子:
	//     0
	//  / / \ \
	// 1 2   3 4
	fmt.Println("UF implements Path Compression by recursion: ")
	uf6.Show()
}
