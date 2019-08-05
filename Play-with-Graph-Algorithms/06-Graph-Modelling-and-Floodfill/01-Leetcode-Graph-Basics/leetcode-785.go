package _1_Leetcode_Graph_Basics

func IsBipartite(graph [][]int) bool {
	// 二分图染色问题
	V := len(graph) // 图的顶点

	visited := make([]bool, V) // 是否访问
	colors := make([]int, V)   // 染色
	//for i := 0; i < V; i++ {
	//	visited[i] = false
	//	colors[i] = 0
	//}

	for v := 0; v < V; v++ {
		if !visited[v] {
			if !dfs(v, 0, visited, colors, graph) {
				return false
			}
		}
	}
	return true
}
func dfs(v, color int, visited []bool, colors []int, graph [][]int) bool {
	visited[v] = true
	colors[v] = color
	for _, w := range graph[v] {
		if !visited[w] {
			if !dfs(w, 1-color, visited, colors, graph) {
				return false
			}
		} else if colors[v] == colors[w] {
			return false
		}
	}
	return true
}
