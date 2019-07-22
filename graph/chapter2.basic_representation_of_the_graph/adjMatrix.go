package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type AdjMatrix struct {
	V   int
	E   int
	adj [][]int
}

func NewAdjMatrix(name string) (*AdjMatrix, error) {
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
	adj := make([][]int, 0)
	for i := 0; i < V; i++ {
		tmp := make([]int, V)
		adj = append(adj, tmp)
	}

	// 3. 根据边E来构造二维矩阵
	E, err := strconv.Atoi(result[1])
	if E < 0 || err != nil {
		return nil, errors.New("获取顶点V失败，确保V为非负整数\n")
	}
	for i := 0; i < 2*E; i++ {
		if i%2 == 0 {
			a, _ := strconv.Atoi(result[2+i])
			validateVertex(a, V) // 判断顶点是否合法
			b, _ := strconv.Atoi(result[3+i])
			validateVertex(b, V)

			if a == b {
				return nil, errors.New("检测到自环边!\n")
			}
			if adj[a][b] == 1 {
				return nil, errors.New("检测到平行边!\n")
			}

			adj[a][b] = 1
			adj[b][a] = 1
		}
	}

	return &AdjMatrix{
		V:   V,
		E:   E,
		adj: adj,
	}, nil
}

// 返回图有几个顶点
func (a *AdjMatrix) Vertex() int {
	return a.V
}

// 返回图有几个边
func (a *AdjMatrix) Edge() int {
	return a.E
}

// 判断两个顶点之间是否存在一条边
func (a *AdjMatrix) HasEdge(v, w int) bool {
	validateVertex(v, a.V)
	validateVertex(w, a.V)
	return a.adj[v][w] == 1
}

// 返回v顶点的相邻的边
func (a *AdjMatrix) Adjacency(v int) []int {
	validateVertex(v, a.V) // 判断用户传入的顶点v是合法的
	res := make([]int, 0)
	for i := 0; i < a.V; i++ {
		if a.adj[v][i] == 1 {
			res = append(res, i)
		}
	}
	return res
}

// 求一个顶点的度degree
func (a *AdjMatrix) Degree(v int) int {
	return len(a.Adjacency(v))
}

// 打印出邻接表表示的图
func (a *AdjMatrix) Print() {
	fmt.Printf("V: %d, E: %d\n", a.V, a.E)
	for i := 0; i < a.V; i++ {
		for j := 0; j < a.V; j++ {
			fmt.Printf("%d ", a.adj[i][j])
		}
		fmt.Println()
	}
}

func validateVertex(v, V int) {
	if v < 0 || v >= V {
		panic("文件内容行数列数与定义不相符\n")
	}
}

func main() {
	adj, err := NewAdjMatrix("g.txt")
	if err != nil {
		fmt.Print(err)
	}
	adj.Print()
}
