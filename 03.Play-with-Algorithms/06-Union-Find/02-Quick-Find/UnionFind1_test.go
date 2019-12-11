package Quick_Find

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	// 使用10000的数据规模
	n := 10000

	// 虽然isConnected只需要O(1)的时间,但由于union操作需要O(n)的时间
	// 总体测试过程的算法复杂度是O(n^2)
	TestUF1(n)
}
