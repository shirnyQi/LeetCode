package Bfs

import "container/list"

/*
给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。



提示：

输出坐标的顺序不重要
m 和 n 都小于150
*/

var (
	h = [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4}}

	want = [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}
)

func pacificAtlantic(heights [][]int) [][]int {
	var result [][]int

	rows := len(heights)
	cols := len(heights[0])

	pacific := initGrahp(rows, cols, -1)
	atlantic := initGrahp(rows, cols, 0)

	for i := 0; i < rows; i++ {
		bfs(heights, pacific, NewPoint(i, 0))
	}

	for i := 0; i < cols; i++ {
		bfs(heights, pacific, NewPoint(0, i))
	}

	for i := 0; i < rows; i++ {
		bfs(heights, atlantic, NewPoint(i, cols-1))
	}

	for i := 0; i < cols; i++ {
		bfs(heights, atlantic, NewPoint(rows-1, i))
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] == atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

type Point [2]int

func NewPoint(x, y int) Point {
	return [2]int{x, y}
}

func initGrahp(rows, cols, value int) [][]int {
	var res = make([][]int, rows, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]int, 0, cols)
		for j := 0; j < cols; j++ {
			res[i] = append(res[i], value)
		}
	}
	return res
}

func bfs(heights, oceans [][]int, point Point) {
	rows := len(heights)
	cols := len(heights[0])

	queue := list.New()
	queue.PushBack(point)
	oceans[point[0]][point[1]] = 1

	for queue.Len() > 0 {
		qlen := queue.Len()
		for i := 0; i < qlen; i++ {
			p := queue.Remove(queue.Front()).(Point)
			x, y := p[0], p[1]
			oceans[x][y] = 1

			if x-1 > -1 && heights[x-1][y] >= heights[x][y] {
				if oceans[x-1][y] != 1 {
					oceans[x-1][y] = 1
					queue.PushBack(NewPoint(x-1, y))
				}
			}

			if x+1 < rows && heights[x+1][y] >= heights[x][y] {
				if oceans[x+1][y] != 1 {
					oceans[x+1][y] = 1
					queue.PushBack(NewPoint(x+1, y))
				}
			}

			if y-1 > -1 && heights[x][y-1] >= heights[x][y] {
				if oceans[x][y-1] != 1 {
					oceans[x][y-1] = 1
					queue.PushBack(NewPoint(x, y-1))
				}
			}

			if y+1 < cols && heights[x][y+1] >= heights[x][y] {
				if oceans[x][y+1] != 1 {
					oceans[x][y+1] = 1
					queue.PushBack(NewPoint(x, y+1))
				}
			}
		}
	}

}
