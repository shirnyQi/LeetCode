func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	rows, cols := len(grid), len(grid[0])
	direction := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	visit := make([][]bool, rows, rows)
	for i := 0; i < rows; i++ {
		visit[i] = make([]bool, cols, cols)
	}
	crtColor := grid[row][col]
	visit[row][col] = true
	queue := list.New()
	queue.PushBack(Point{row, col})

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			p := queue.Remove(queue.Front()).(Point)
			x, y := p.x, p.y
			for _, direct := range direction {
				nextx, nexty := x+direct.x, y+direct.y
				if nextx < 0 || nextx >= rows || nexty < 0 || nexty >= cols {
					// 邻居节点出界，则当前节点涂色
					grid[x][y] = color
					continue
				}

				if visit[nextx][nexty] {
					continue
				}

				// 对未访问过的访问过的联通分量标记
				if grid[nextx][nexty] == crtColor {
					queue.PushBack(Point{nextx, nexty})
					visit[nextx][nexty] = true
				} else {
					grid[x][y] = color
				}
			}
		}
	}

	return grid
}

type Point struct {
	x int
	y int
}
