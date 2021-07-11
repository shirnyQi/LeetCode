import "container/list"

/*
给定一个包含了一些 0 和 1 的非空二维数组grid 。

一个岛屿是由一些相邻的1(代表土地)构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
你可以假设grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

示例1：
输入：
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],
[0,1,1,0,1,0,0,0,0,0,0,0,0],
[0,1,0,0,1,1,0,0,1,0,1,0,0],
[0,1,0,0,1,1,0,0,1,1,1,0,0],
[0,0,0,0,0,0,0,0,0,0,1,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],
[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6

示例2：
输入：
[[0,0,0,0,0,0,0,0]]
输出：0
*/

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
