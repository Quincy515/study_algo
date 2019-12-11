package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/02.Play-with-Graph-Algorithms/utils"
)

// 两个顶点之间的路径
type Path struct {
	G       *Graph.AdjList
	source  int    // 源 顶点
	target  int    // 目标顶点
	pre     []int  // 记录前一个顶点
	visited []bool // 是否被遍历过

}

func NewPath(G *Graph.AdjList, source, target int) *Path {
	Graph.ValidateVertex(source, G.V) // 确保用户传入的顶点s是合法的
	Graph.ValidateVertex(target, G.V) // 确保用户传入的顶点s是合法的

	visited := make([]bool, G.V) // 记录每个节点是否被遍历过
	for i := 0; i < G.V; i++ {   // 给visited赋默认值false
		visited[i] = false // false 表示未被访问
	}

	pre := make([]int, G.V)    // 图中有多少个顶点，pre中就有多少个元素
	for i := 0; i < G.V; i++ { // 给pre赋默认值-1
		pre[i] = -1 // -1 表示未被访问，没有前一顶点
	}

	dfs(source, source, target, G, pre, visited) // 单源路径，所以dfs遍历源顶点的路径就可以了

	for _, e := range visited {
		fmt.Print(e, " ")
	}
	fmt.Println()

	return &Path{
		G:       G,
		source:  source,
		target:  target,
		pre:     pre,
		visited: visited,
	}
}

func dfs(v, parent, target int, G *Graph.AdjList, pre []int, visited []bool) bool {
	visited[v] = true
	pre[v] = parent

	if v == target {
		return true
	}

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int) // 遍历顶点v的所有相邻点
		if !visited[ww] {   // 如果没有遍历过，就递归调用
			if dfs(ww, v, target, G, pre, visited) {
				return true
			}
		}
	}

	return false
}

// 源s能否到达目标target
func (s *Path) isConnected() bool {
	return s.visited[s.target] // 查看目标节点是否已经被遍历到
}

func (s *Path) path() []int {
	res := []int{}
	if !s.isConnected() {
		return res // 源s到目标t没有路径,返回空数组
	}

	cur := s.target // 目标顶点往回走到源s
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
		path := NewPath(g, 0, 6)                  // true true true true false false true
		fmt.Println("顶点0 -> 6的路径: ", path.path()) // 顶点0->6路径: [0 1 3 2 6]

		path2 := NewPath(g, 0, 1)                  // true true false false false false false
		fmt.Println("顶点0 -> 1的路径: ", path2.path()) // 顶点0->1路径: [0 1]

		path3 := NewPath(g, 0, 5)                  // true true true true true false true
		fmt.Println("顶点0 -> 5的路径: ", path3.path()) // 顶点0->5路径: []
	}
}
