package LeetCode_773

import (
	"bytes"
	"strconv"
)

func slidingPuzzle(board [][]int) int {
	queue := make([]string, 0)           // 队列中存储的元素是每一次board的状态
	visited := make(map[string]int)      // 记录达到当前顶点所处的状态对应走了多少步
	initialState := boardToString(board) // 二维数组转化成字符串，初始状态
	if initialState == "123450" {        // 初始状态如果等于终止状态
		return 0
	}

	// BFS
	queue = append(queue, initialState) // 向队列中添加初始状态
	visited[initialState] = 0           // 初始状态0步
	for len(queue) > 0 {                // 只要队列不为空就继续执行循环
		cur := queue[0]              // 取出队首元素
		queue = queue[1:]            // 移除队首元素
		nexts := getNexts(cur)       // 根据cur找到下一个其他的状态
		for _, next := range nexts { // 对每一个状态进行判断
			if _, ok := visited[next]; !ok { // 没有被访问过
				queue = append(queue, next)      // 添加到队列
				visited[next] = visited[cur] + 1 // 记录到达该状态的步数
				if next == "123450" {            // 如果是最终状态
					return visited[next] // 返回达到终止状态的步骤
				}
			}
		}
	}
	return -1 // 没有找到终止状态
}

func getNexts(s string) []string { // 传入一个状态，返回对应的其他状态
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 对某一个格子四个方向相对的位移
	cur := stringToBoard(s)                           // 将当前状态恢复成2*3的二维数组
	var zero int                                      // 先找到空的格子
	for zero = 0; zero < 6; zero++ {
		if cur[zero/3][zero%3] == 0 {
			break
		}
	}
	var res []string
	zx, zy := zero/3, zero%3 // 空格子位置转化成二维坐标
	for d := 0; d < 4; d++ { // 滑动到空格子位置,就是和空格子交换位置
		nextx := zx + dirs[d][0]  // 空格子可以移动上下左右四个方向
		nexty := zy + dirs[d][1]  // 获取空格子相邻位置的坐标(nextx,nexty)
		if inArea(nextx, nexty) { // 保证位置是合法的
			swap(cur, zx, zy, nextx, nexty)       // 空格子和相邻的交换位置
			res = append(res, boardToString(cur)) // 交换之后就形成了下一个牌面的情况
			swap(cur, zx, zy, nextx, nexty)       // 牌面恢复成原来的位置，再寻找下一个牌面
		}
	}
	return res
}

func swap(board [][]int, x1, y1, x2, y2 int) {
	t := board[x1][y1]
	board[x1][y1] = board[x2][y2]
	board[x2][y2] = t
}

func inArea(x, y int) bool {
	return x >= 0 && x < 2 && y >= 0 && y < 3
}

func boardToString(board [][]int) string {
	buffer := bytes.Buffer{}

	// 把每个数字转化成字符，组成字符串
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			buffer.WriteString(strconv.Itoa(board[i][j]))
		}
	}
	return buffer.String()
}

func stringToBoard(s string) [][]int {
	board := make([][]int, 0)
	for range make([]int, 2) { // 初始化空的二维数组
		board = append(board, make([]int, 3))
	}

	for i := 0; i < 6; i++ {
		board[i/3][i%3] = int(s[i] - '0') // 一位索引转化成二维索引
	}
	return board
}
