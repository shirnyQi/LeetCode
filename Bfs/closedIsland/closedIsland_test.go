package closedIsland

import (
	"container/list"
	"fmt"
	"testing"
)

/*
BFS:统计封闭岛屿的数目

有一个二维矩阵 grid ，每个位置要么是陆地（记号为 0 ）要么是水域（记号为 1 ）。
我们从一块陆地出发，每次可以往上下左右 4 个方向相邻区域走，能走到的所有陆地区域，我们将其称为一座「岛屿」。
如果一座岛屿 完全 由水域包围，即陆地边缘上下左右所有相邻区域都是水域，那么我们将其称为 「封闭岛屿」。
请返回封闭岛屿的数目。



示例 1：
输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。

示例 2：
输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1

示例 3：
输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/bfs/e6ko9i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func closedIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	// BFS先将所有非封闭的岛屿（0）标记为1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != 0 && j != 0 && i != rows-1 && j != cols-1 {
				continue
			}
			if grid[i][j] == 0 {
				bfs(grid, Point{i, j}, rows, cols)
			}

		}
	}

	// 再用DFS查找封闭岛屿
	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				dfs(grid, Point{i, j}, rows, cols)
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
	grid[point.x][point.y] = 1

	xShift := []int{-1, 1, 0, 0}
	yShift := []int{0, 0, -1, 1}

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			point := queue.Remove(queue.Front()).(Point)
			x, y := point.x, point.y
			for j := 0; j < 4; j++ {
				nextx, nexty := x+xShift[j], y+yShift[j]
				if nextx >= 0 && nextx < rows && nexty >= 0 && nexty < cols && grid[nextx][nexty] == 0 {
					queue.PushBack(Point{nextx, nexty})
					grid[nextx][nexty] = 1
				}
			}
		}
	}
}

func dfs(grid [][]int, point Point, rows, cols int) {
	x, y := point.x, point.y

	if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == 1 {
		return
	}

	grid[x][y] = 1
	dfs(grid, Point{x - 1, y}, rows, cols)
	dfs(grid, Point{x + 1, y}, rows, cols)
	dfs(grid, Point{x, y - 1}, rows, cols)
	dfs(grid, Point{x, y + 1}, rows, cols)
}

func TestName(t *testing.T) {
	var grid [][]int

	grid = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}}
	fmt.Println(closedIsland(grid)) // 2

	grid = [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0}}
	fmt.Println(closedIsland(grid)) // 1

	grid = [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1}}
	fmt.Println(closedIsland(grid)) // 2
}
