package Water_Puzzle

func WaterPuzzle() {
	var queue []int
	visited := make(map[int]bool, 100)
	queue = append(queue, 0)
	visited[0] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		a, b := cur/10, cur%10
		// max a = 5, max b = 3
		var nexts []int
		// TODO: nexts

		for _, next := range nexts {
			if !visited[next] {
				queue = append(queue, next)
				visited[next] = true

				if next/10 == 4 || next%10 == 4 {
					return
				}
			}
		}
	}
}
