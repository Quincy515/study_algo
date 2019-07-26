package _9_Unweighted_Shortest_Path

import (
	"container/list"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Graph struct {
	V   int
	E   int
	adj []*list.List
}

func NewGraph(name string) (*Graph, error) {
	// 1. 读取文件
	result := make([]string, len(name))
	if contents, err := ioutil.ReadFile(name); err == nil {
		result = strings.Fields(string(contents))
	}

	// 2. 根据顶点V声明二维矩阵
	V, err := strconv.Atoi(result[0])
	if V < 0 || err != nil {
		return nil, errors.New("获取顶点V失败，确保V为非负整数\n")
	}
	adj := make([]*list.List, 0) // adj是一个链表数组
	for i := 0; i < V; i++ {
		adj = append(adj, list.New()) // 第i个元素表示的是第i个顶点和它相邻的顶点
	}

	// 3. 根据边E来构造二维矩阵
	E, err := strconv.Atoi(result[1])
	if E < 0 || err != nil {
		return nil, errors.New("获取顶点V失败，确保V为非负整数\n")
	}
	for i := 0; i < 2*E; i++ {
		if i%2 == 0 {
			a, _ := strconv.Atoi(result[2+i])
			ValidateVertex(a, V) // 判断顶点是否合法
			b, _ := strconv.Atoi(result[3+i])
			ValidateVertex(b, V)

			if a == b {
				return nil, errors.New("检测到自环边!\n")
			}
			if contain(adj[a], b) {
				return nil, errors.New("检测到平行边!\n")
			}

			adj[a].PushBack(b)
			adj[b].PushBack(a)
		}
	}

	return &Graph{
		V:   V,
		E:   E,
		adj: adj,
	}, nil
}

// 返回图有几个顶点
func (a *Graph) Vertex() int {
	return a.V
}

// 返回图有几个边
func (a *Graph) Edge() int {
	return a.E
}

// 判断两个顶点之间是否存在一条边
func (a *Graph) HasEdge(v, w int) bool {
	ValidateVertex(v, a.V)
	ValidateVertex(w, a.V)
	return contain(a.adj[v], w)
}

// 返回v顶点的相邻的边
func (a *Graph) Adjacency(v int) *list.List {
	ValidateVertex(v, a.V) // 判断用户传入的顶点v是合法的
	return a.adj[v]
}

// 求一个顶点的度degree
func (a *Graph) Degree(v int) int {
	return a.Adjacency(v).Len()
}

// 打印出邻接表表示的图
func (a *Graph) Print() {
	fmt.Printf("V: %d, E: %d\n", a.V, a.E)
	for i := 0; i < a.V; i++ {
		fmt.Printf("%d: ", i)
		for w := a.adj[i].Front(); w != nil; w = w.Next() {
			p := w.Value.(int)
			fmt.Printf("%d ", p)
		}
		fmt.Println()
	}
}

func ValidateVertex(v, V int) {
	if v < 0 || v >= V {
		panic("文件内容行数列数与定义不相符\n")
	}
}

func contain(v *list.List, w int) bool {
	for i := v.Front(); i != nil; i = i.Next() {
		if i.Value == w {
			return true
		}
	}
	return false
}
