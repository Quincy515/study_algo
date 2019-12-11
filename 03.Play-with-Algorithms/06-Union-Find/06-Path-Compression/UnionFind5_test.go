package Path_Compression

import (
	Optimize_by_Size "github.com/custergo/study_algo/03.Play-with-Algorithms/06-Union-Find/04-Optimize-by-Size"
	Optimize_by_Rank "github.com/custergo/study_algo/03.Play-with-Algorithms/06-Union-Find/05-Optimize-by-Rank"
	"testing"
)

func TestUnionFind(t *testing.T) {
	// 使用10000的数据规模
	n := 1000000

	// 虽然isConnected只需要O(1)的时间,但由于union操作需要O(n)的时间
	// 总体测试过程的算法复杂度是O(n^2)
	//Quick_Find.TestUF1(n)

	// 对于UF2来说,其时间性能是O(n*h),h为并查集表达的树的最大高度
	// 这里严格来讲,h和logn没有关系,不过大家课件简单这么理解
	// 后序内容会对h进行优化,总体而言,这个h是远小于n的
	// 所以实现的UF2测试结果远远好于UF1,n越大越明显
	//Quick_Union.TestUF2(n)

	// 对于UF3来说,其时间性能依然是O(n*h),h为并查集表达的树的最大高度
	// 但由于UF3能更高概率的保证树的平衡,所以性能更优
	Optimize_by_Size.TestUF3(n)

	// UF4虽然相对UF3进行了优化,但优化的地方出现的情况较少
	// 所以性能更优表现的不明显,甚至在一些数据下性能会更差
	Optimize_by_Rank.TestUF4(n)

	// UF5虽然相对UF4进行了优化,但优化的地方出现的情况较少,
	// 所以性能更优表现的不明显,甚至在一些数据下性能会更差
	TestUF5(n)
}
