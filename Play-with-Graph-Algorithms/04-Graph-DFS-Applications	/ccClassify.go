package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/graph/utils"
)

type CC struct {
	cccount int
	visited []int
	V       int
}

// 判断两个顶点是否属于同一个联通分量
func CCClassify(G *Graph.AdjList) *CC {
	cccount := 0                // 联通分量的个数
	visited := make([]int, G.V) // 记录每个节点是否被遍历过
	for i := 0; i < G.V; i++ {  // 给visited赋默认值-1
		visited[i] = -1 // -1 表示未被访问
	}
	for v := 0; v < G.V; v++ { // 遍历每个顶点，避免忽略联通分量
		if visited[v] == -1 {
			dfs(v, G, cccount, visited)
			cccount++
		}
	}

	return &CC{
		cccount: cccount,
		visited: visited,
		V:       G.V,
	}
}

// 图中联通分量的个数
func (c *CC) count() int {
	// for k, v := range c.visited {
	// 	fmt.Printf("k: %d, v: %d\n", k, v)
	// }
	return c.cccount
}

// 图中的两个顶点是否属于同一个联通分量
func (c *CC) isConnected(v, w int) bool {
	Graph.ValidateVertex(v, c.V)
	Graph.ValidateVertex(w, c.V)
	return c.visited[v] == c.visited[w]
}

// 图中有多少联通分量，每个联通分量包含多少顶点
func (c *CC) components() [][]int {
	res := make([][]int, 0)
	for i := 0; i < c.cccount; i++ {
		tmp := make([]int, 0)
		res = append(res, tmp)
	}
	for v := 0; v < c.V; v++ {
		res[c.visited[v]] = append(res[c.visited[v]], v)
	}
	return res
}

func dfs(v int, G *Graph.AdjList, ccid int, visited []int) {
	visited[v] = ccid

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int)    // 遍历顶点v的所有相邻点
		if visited[ww] == -1 { // 如果没有遍历过，就递归调用
			dfs(ww, G, ccid, visited)
		}
	}

}

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		cc := CCClassify(g)
		fmt.Println("联通分量的个数: ", cc.count())                   // 联通分量的个数: 2
		fmt.Println("两个顶点是否属于同一个联通分量: ", cc.isConnected(0, 6)) // true
		fmt.Println("两个顶点是否属于同一个联通分量: ", cc.isConnected(0, 5)) // false
		comp := cc.components()
		for ccid := 0; ccid < len(comp); ccid++ {
			fmt.Print(ccid, " : ")         // 0: 0 1 2 3 4 6
			for _, w := range comp[ccid] { // 1 : 5
				fmt.Print(w, " ")
			}
			fmt.Println()
		}
	}
}
