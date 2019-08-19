package UnionFind

import (
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Chapter-06-Completed-Code/UF1"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Chapter-06-Completed-Code/UF2"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Chapter-06-Completed-Code/UF3"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Chapter-06-Completed-Code/UF4"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Chapter-06-Completed-Code/UF5"
	"github.com/custergo/study_algo/3.Play-with-Algorithms/06-Union-Find/Chapter-06-Completed-Code/UF6"
	"testing"
)

func TestUnionFind(t *testing.T) {
	// 使用10000的数据规模
	n := 10000

	// 虽然isConnected只需要O(1)的时间,但由于union操作需要O(n)的时间
	// 总体测试过程的算法复杂度是O(n^2)
	uf1 := UF1.NewUnionFind(n)
	TestUF("UF1", uf1, n)

	// 对于UF2来说,其时间性能是O(n*h),h为并查集表达的树的最大高度
	// 这里严格来讲,h和logn没有关系,不过大家课件简单这么理解
	// 后序内容会对h进行优化,总体而言,这个h是远小于n的
	// 所以实现的UF2测试结果远远好于UF1,n越大越明显
	uf2 := UF2.NewUnionFind(n)
	TestUF("UF2", uf2, n)

	// 对于UF3来说,其时间性能依然是O(n*h),h为并查集表达的树的最大高度
	// 但由于UF3能更高概率的保证树的平衡,所以性能更优
	uf3 := UF3.NewUnionFind(n)
	TestUF("UF3", uf3, n)

	// UF4虽然相对UF3进行了优化,但优化的地方出现的情况较少
	// 所以性能更优表现的不明显,甚至在一些数据下性能会更差
	uf4 := UF4.NewUnionFind(n)
	TestUF("UF4", uf4, n)

	// UF5虽然相对UF4进行了优化,但优化的地方出现的情况较少,
	// 所以性能更优表现的不明显,甚至在一些数据下性能会更差
	uf5 := UF5.NewUnionFind(n)
	TestUF("UF5", uf5, n)

	// UF65虽然相对UF4进行了优化,但优化的地方出现的情况较少,
	// 所以性能更优表现的不明显,甚至在一些数据下性能会更差
	uf6 := UF6.NewUnionFind(n)
	TestUF("UF6", uf6, n)
}
