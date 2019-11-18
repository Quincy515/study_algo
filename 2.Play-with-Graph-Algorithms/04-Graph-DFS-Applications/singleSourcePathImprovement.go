package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/2.Play-with-Graph-Algorithms/utils"
)

// 单源路径优化
type SingleSourcePath struct {
	V      int
	G      *Graph.AdjList
	source int   // 源 顶点
	pre    []int // 记录前一个顶点
}

func NewSingleSourcePath(G *Graph.AdjList, source int) *SingleSourcePath {
	Graph.ValidateVertex(source, G.V) // 确保用户传入的顶点s是合法的

	pre := make([]int, G.V)    // 图中有多少个顶点，pre中就有多少个元素
	for i := 0; i < G.V; i++ { // 给pre赋默认值-1
		pre[i] = -1 // -1 表示未被访问，没有前一顶点
	}

	dfs(source, source, G, pre) // 单源路径，所以dfs遍历源顶点的路径就可以了

	return &SingleSourcePath{
		V:      G.V,
		G:      G,
		source: source,
		pre:    pre,
	}
}

func dfs(v, parent int, G *Graph.AdjList, pre []int) {
	pre[v] = parent

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int) // 遍历顶点v的所有相邻点
		if pre[ww] == -1 {  // 如果没有遍历过，就递归调用
			dfs(ww, v, G, pre)
		}
	}

}

// 源s能否到达目标target
func (s *SingleSourcePath) isConnectedTo(target int) bool {
	Graph.ValidateVertex(target, s.V) // 验证用户传递的目标节点是否合法
	return s.pre[target] != -1        // 查看目标节点是否已经被遍历到
}

func (s *SingleSourcePath) path(target int) []int {
	res := []int{}
	if !s.isConnectedTo(target) {
		return res // 源s到目标t没有路径,返回空数组
	}

	cur := target // 目标顶点往回走到源s
	for cur != s.source {
		res = append(res, cur) // 把cur保存到路径中
		cur = s.pre[cur]       // cur赋值为前一个顶点
	}
	res = append(res, s.source) // 把源s添加到路径

	reverse(res) // 反转路径res
	return res
}

// 反转函数
func reverse(l []int) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
}

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		sspath := NewSingleSourcePath(g, 0)
		fmt.Println("顶点0 -> 6的路径: ", sspath.path(6)) // 顶点0->6路径: [0 1 3 2 6]
		fmt.Println("顶点0 -> 5的路径: ", sspath.path(5)) // 顶点0->5路径: []
	}
}
