package colorBorder

import (
	"container/list"
	"fmt"
	"testing"
)

/*
给出一个二维整数网格 grid，网格中的每个值表示该位置处的网格块的颜色。
只有当两个网格块的颜色相同，而且在四个方向中任意一个方向上相邻时，它们属于同一连通分量。
连通分量的边界是指连通分量中的所有与不在分量中的正方形相邻（四个方向上）的所有正方形，或者在网格的边界上（第一行/列或最后一行/列）的所有正方形。
给出位于 (r0, c0) 的网格块和颜色 color，使用指定颜色 color 为所给网格块的连通分量的边界进行着色，并返回最终的网格 grid 。

示例 1：
输入：grid = [[1,1],[1,2]], r0 = 0, c0 = 0, color = 3
输出：[[3, 3], [3, 2]]

示例 2：
输入：grid = [[1,2,2],[2,3,2]], r0 = 0, c0 = 1, color = 3
输出：[[1, 3, 3], [2, 3, 3]]

示例 3：
输入：grid = [[1,1,1],[1,1,1],[1,1,1]], r0 = 1, c0 = 1, color = 2
输出：[[2, 2, 2], [2, 1, 2], [2, 2, 2]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coloring-a-border
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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

func TestName(t *testing.T) {
	var grid [][]int
	grid = [][]int{
		{1, 1},
		{1, 2},
	}
	fmt.Println(colorBorder(grid, 0, 0, 3)) // [[3, 3], [3, 2]]

	grid = [][]int{
		{1, 2, 2},
		{2, 3, 3},
	}
	fmt.Println(colorBorder(grid, 0, 1, 3)) // [[1, 3, 3], [2, 3, 3]]

	grid = [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Println(colorBorder(grid, 1, 1, 2)) // [[2, 2, 2], [2, 1, 2], [2, 2, 2]]

	grid = [][]int{
		{1, 1, 1, 1},
		{0, 0, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 1},
	}
	fmt.Println(colorBorder(grid, 0, 0, 2)) // [[2, 2, 2, 2], [0, 0, 2, 2], [0, 2, 1, 2], [2, 2, 2, 2]]
}
