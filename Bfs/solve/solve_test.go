package solve

import (
	"container/list"
	"fmt"
	"testing"
)

/*
被围绕的区域

给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例 1：

输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

示例 2：
输入：board = [["X"]]
输出：[["X"]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Point struct {
	x int
	y int
}

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != 0 && j != 0 && i != rows-1 && j != cols-1 {
				continue
			}
			if board[i][j] == 'O' {
				bfs(board, Point{i, j})
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func bfs(board [][]byte, p Point) {
	rows, cols := len(board), len(board[0])
	queue := list.New()
	queue.PushBack(p)
	board[p.x][p.y] = '#'

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(Point)
			x, y := node.x, node.y
			if x-1 > 0 && board[x-1][y] == 'O' {
				board[x-1][y] = '#'
				queue.PushBack(Point{x - 1, y})
			}
			if x+1 < rows && board[x+1][y] == 'O' {
				board[x+1][y] = '#'
				queue.PushBack(Point{x + 1, y})
			}
			if y-1 > 0 && board[x][y-1] == 'O' {
				board[x][y-1] = '#'
				queue.PushBack(Point{x, y - 1})
			}
			if y+1 < cols && board[x][y+1] == 'O' {
				board[x][y+1] = '#'
				queue.PushBack(Point{x, y + 1})
			}
		}
	}
}

func TestName(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)

	board = [][]byte{
		{'X'},
	}
	solve(board)
	fmt.Println(board)

	board = [][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}
	solve(board)
	fmt.Println(board)
}
