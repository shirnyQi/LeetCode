package numEnclaves

import (
	"container/list"
	"fmt"
	"testing"
)

/*
BFS:飞地的数量

给出一个二维数组 A，每个单元格为 0（代表海）或 1（代表陆地）。
移动是指在陆地上从一个地方走到另一个地方（朝四个方向之一）或离开网格的边界。
返回网格中无法在任意次数的移动中离开网格边界的陆地单元格的数量。

示例 1：
输入：[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：
有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。

示例 2：
输入：[[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：
所有 1 都在边界上或可以到达边界。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/bfs/e66mec/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func numEnclaves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != 0 && j != 0 && i != rows-1 && j != cols-1 {
				continue
			}
			if grid[i][j] == 1 {
				bfs(grid, Point{i, j}, rows, cols)
			}

		}
	}

	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

type Point struct {
	x int
	y int
}

func bfs(grid [][]int, point Point, rows, cols int) {
	queue := list.New()
	queue.PushBack(point)
	grid[point.x][point.y] = 2

	xShift := []int{-1, 1, 0, 0}
	yShift := []int{0, 0, -1, 1}

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			point := queue.Remove(queue.Front()).(Point)
			x, y := point.x, point.y
			for j := 0; j < 4; j++ {
				nextx, nexty := x+xShift[j], y+yShift[j]
				if nextx >= 0 && nextx < rows && nexty >= 0 && nexty < cols && grid[nextx][nexty] == 1 {
					queue.PushBack(Point{nextx, nexty})
					grid[nextx][nexty] = 2
				}
			}
		}
	}
}

func TestName(t *testing.T) {
	fmt.Println(numEnclaves([][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	})) // 3

	fmt.Println(numEnclaves([][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	})) // 0
}
