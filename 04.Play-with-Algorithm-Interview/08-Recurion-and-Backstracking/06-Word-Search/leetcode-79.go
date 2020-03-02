package _6_Word_Search

// 辅助递归函数，输入是字符矩阵，以及标记矩阵上是否已经访问过的二维数组，以及当前要考察的坐标i，j，
// 要查找的单词word，以及这个单词当前要考察字符的下标
// 输出是布尔值，表示是否能在字符矩阵中找到单词
func existHelp(board [][]byte, visited [][]bool, i, j int, word string, idx int) bool {
	if idx == len(word) { // 如果要考察的单词字符下标已经等于单词长度
		return true // 说明单词上所有的字符都已经在矩阵中找到，返回true
	}
	// 否则先检查i，j是否在合法范围内
	// 或者i,j这个位置的字符已经访问过
	// 或者虽然没有访问过，但这个位置的字符不等于单词当前要考察的字符
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		visited[i][j] || board[i][j] != word[idx] {
		return false // 则返回false
	} // 如果不是以上情况，说明i，j这个位置的字符等于单词当前要考察的字符
	visited[i][j] = true // 于是把位置i，j标识为已访问过
	// 然后递归调用自己去访问上下左右四个位置,单词要考察的下标idx+1
	// 这四个调用只要有一个找到单词即可，于是把这四个递归调用或起来赋值给existed
	existed := existHelp(board, visited, i-1, j, word, idx+1) ||
		existHelp(board, visited, i+1, j, word, idx+1) ||
		existHelp(board, visited, i, j-1, word, idx+1) ||
		existHelp(board, visited, i, j+1, word, idx+1)
	visited[i][j] = false // 退递归后需要把当前位置的访问标记设置为false
	return existed        // 最后返回existed即可
}

// Time: O(m*n*3^k), Space: O(m*n)
func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 ||
		board[0] == nil || len(board[0]) == 0 { // 处理边界情况，字符矩阵为空或长度为0
		return false // 则返回false
	}
	m, n := len(board), len(board[0]) // 否则把矩阵两个维度的大小分别赋值给m和n
	visited := make([][]bool, m)      // 然后定义一个二维辅助数组
	for i := range visited {          // 去记录字符矩阵上各个位置的字符是否已经访问过
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ { // 接下来遍历字符矩阵
		for j := 0; j < n; j++ { // 从位置i，j开始查找单词
			if existHelp(board, visited, i, j, word, 0) {
				return true // 如果找到就返回true
			}
		}
	} // 循环结束后还没有返回true则说明字符矩阵中不存在这个单词
	return false // 返回false
}
