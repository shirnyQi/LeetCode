import "container/list"

func maxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	var maxIsland int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 0 {
				maxIsland = maxArea(maxIsland, bsf2(grid, NewPoint(i, j), rows, cols))
			}
		}
	}
	return maxIsland
}

type Point [2]int

func NewPoint(x, y int) Point {
	return [2]int{x, y}
}

func maxArea(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bsf2(grid [][]int, point Point, rows, cols int) int {
	var res []Point

	queue := list.New()
	queue.PushBack(point)
	grid[point[0]][point[1]] = 0

	for queue.Len() > 0 {
		lenq := queue.Len()
		for i := 0; i < lenq; i++ {
			crtPoint := queue.Remove(queue.Front()).(Point)
			res = append(res, crtPoint)
			x, y := crtPoint[0], crtPoint[1]
			if x-1 > -1 && grid[x-1][y] != 0 {
				grid[x-1][y] = 0
				queue.PushBack(NewPoint(x-1, y))
			}
			if x+1 < rows && grid[x+1][y] != 0 {
				grid[x+1][y] = 0
				queue.PushBack(NewPoint(x+1, y))
			}
			if y-1 > -1 && grid[x][y-1] != 0 {
				grid[x][y-1] = 0
				queue.PushBack(NewPoint(x, y-1))
			}
			if y+1 < cols && grid[x][y+1] != 0 {
				grid[x][y+1] = 0
				queue.PushBack(NewPoint(x, y+1))
			}
		}
	}
	return len(res)
}
