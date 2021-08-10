package shortestBridge

import (
	"container/list"
	"fmt"
	"testing"
)

/*
934. 最短的桥
在给定的二维二进制数组 A 中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。）
现在，我们可以将 0 变为 1，以使两座岛连接起来，变成一座岛。
返回必须翻转的 0 的最小数目。（可以保证答案至少是 1 。）

示例 1：
输入：A = [[0,1],[1,0]]
输出：1

示例 2：
输入：A = [[0,1,0],[0,0,0],[0,0,1]]
输出：2

示例 3：
输入：A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-bridge
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Point struct {
	x int
	y int
}

func shortestBridge(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	islandA := getIsLand(grid)

	var res int = -1
	for islandA.Len() > 0 {
		res++
		length := islandA.Len()
		for i := 0; i < length; i++ {
			ponit := islandA.Remove(islandA.Front()).(Point)
			x, y := ponit.x, ponit.y
			if x-1 > 0 && grid[x-1][y] != 2 {
				if grid[x-1][y] == 1 {
					return res
				}
				grid[x-1][y] = 2
				islandA.PushBack(Point{x - 1, y})
			}
			if x+1 < rows && grid[x+1][y] != 2 {
				if grid[x+1][y] == 1 {
					return res
				}
				grid[x+1][y] = 2
				islandA.PushBack(Point{x + 1, y})
			}
			if y-1 > 0 && grid[x][y-1] != 2 {
				if grid[x][y-1] == 1 {
					return res
				}
				grid[x][y-1] = 2
				islandA.PushBack(Point{x, y - 1})
			}
			if y+1 < cols && grid[x][y+1] != 2 {
				if grid[x][y+1] == 1 {
					return res
				}
				grid[x][y+1] = 2
				islandA.PushBack(Point{x, y + 1})
			}
		}
	}
	return res
}

func getIsLand(grid [][]int) *list.List {
	islandA := list.New()
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				dfs(grid, Point{i, j}, islandA, rows, cols)
				return islandA
			}
		}
	}
	return islandA
}

func dfs(grid [][]int, point Point, islandA *list.List, rows, cols int) {
	x, y := point.x, point.y
	if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == 2 {
		return
	}
	if grid[x][y] == 0 {
		return
	}
	if grid[x][y] == 1 {
		islandA.PushBack(point)
		grid[x][y] = 2
	}

	dfs(grid, Point{x - 1, y}, islandA, rows, cols)
	dfs(grid, Point{x + 1, y}, islandA, rows, cols)
	dfs(grid, Point{x, y - 1}, islandA, rows, cols)
	dfs(grid, Point{x, y + 1}, islandA, rows, cols)
}

func TestName(t *testing.T) {
	var grid [][]int
	grid = [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println(shortestBridge(grid)) //1
	grid = [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1},
	}
	fmt.Println(shortestBridge(grid)) //2
	grid = [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(shortestBridge(grid)) //1

	grid = [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	fmt.Println(shortestBridge(grid)) //1
}
