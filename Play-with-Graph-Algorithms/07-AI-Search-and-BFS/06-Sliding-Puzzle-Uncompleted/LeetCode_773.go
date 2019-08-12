package LeetCode_773

import (
	"bytes"
	"strconv"
)

func slidingPuzzle(board [][]int) int {
	queue := make([]string, 0)
	visited := make(map[string]int)
	initialState := boardToString(board)
	if initialState == "123450" {
		return 0
	}

	// BFS
	queue = append(queue, initialState)
	visited[initialState] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		nexts := getNexts(cur)
		for _, next := range nexts {
			if _, ok := visited[next]; !ok {
				queue = append(queue, next)
				visited[next] = visited[cur] + 1
				if next == "123450" {
					return visited[next]
				}
			}
		}
	}
	return -1
}

func getNexts(s string) []string {

}

func boardToString(board [][]int) string {
	buffer := bytes.Buffer{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			buffer.WriteString(strconv.Itoa(board[i][j]))
		}
	}
	return buffer.String()
}
