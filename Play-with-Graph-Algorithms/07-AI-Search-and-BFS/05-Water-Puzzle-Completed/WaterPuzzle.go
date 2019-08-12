package Water_Puzzle

type WaterPuzzle struct {
	pre []int
	end int
}

func NewWaterPuzzle() *WaterPuzzle {
	pre := make([]int, 100)
	end := -1
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
		nexts = append(nexts, (5*10 + b))
		nexts = append(nexts, a*10+3)
		nexts = append(nexts, (0*10 + b))
		nexts = append(nexts, a*10+0)

		x := min(a, 3-b)
		nexts = append(nexts, (a-x)*10+(b+x))

		y := min(5-a, b)
		nexts = append(nexts, (a+y)*10+(b-y))

		for _, next := range nexts {
			if !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				pre[next] = cur

				if next/10 == 4 || next%10 == 4 {
					end = next
					break
				}
			}
		}
	}
	return &WaterPuzzle{
		pre: pre,
		end: end,
	}
}

func (w *WaterPuzzle) result() []int {
	var res []int
	if w.end == -1 {
		return res
	}

	cur := w.end
	for cur != 0 {
		res = append(res, cur)
		cur = w.pre[cur]
	}
	res = append(res, 0)
	reverse(res)
	return res
}

// 反转函数
func reverse(l []int) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
