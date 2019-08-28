package Dijkstra

import (
	"bytes"
	"fmt"
)

// 边
type Edge struct {
	a, b int    // 边的两个端点
	w    weight // 边的权值
}

// 构造函数
func NewEdge(a, b int, w weight) *Edge {
	return &Edge{a, b, w}
}

// 空的构造函数,所有成员变量都取默认值
func NewEdgeNP() *Edge        { return &Edge{} }
func (this *Edge) V() int     { return this.a } // 返回第一个顶点
func (this *Edge) W() int     { return this.b } // 返回第二个顶点
func (this *Edge) Wt() weight { return this.w } // 返回权值
func (this *Edge) Other(x int) int { // 给定一个顶点,返回另一个顶点
	assert(x == this.a || x == this.b)
	if x == this.a {
		return this.b
	} else {
		return this.a
	}
}

// 输出边的信息
func (this *Edge) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(this.a, "-", this.b, ": ", this.w))
	return buffer.String()
}

// 边的大小比较,是对边的权值的大小比较
func (this *Edge) compareTo(that Edge) int {
	if compare(this.w, that.Wt()) < 0 {
		return -1
	} else if compare(this.w, that.Wt()) > 0 {
		return 1
	} else {
		return 0
	}
}
