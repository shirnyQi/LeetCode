package longestIncreasingPath

import (
	"container/list"
	"fmt"
	"testing"
)

/*
329. 矩阵中的最长递增路径
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。


示例 1：
输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
输出：4
解释：最长递增路径为 [1, 2, 6, 9]。

示例 2：
输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
输出：4
解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。

示例 3：
输入：matrix = [[1]]
输出：1

link:https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/
*/

// bfs
func longestIncreasingPath(matrix [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	rows, cols := len(matrix), len(matrix[0])

	degrees := make([][]int, rows, rows)

	for i := 0; i < rows; i++ {
		degrees[i] = make([]int, cols, cols)
		for j := 0; j < cols; j++ {
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if x >= 0 && x < rows && y >= 0 && y < cols && matrix[i][j] < matrix[x][y] {
					degrees[i][j]++
				}
			}
		}
	}

	var res int

	q := list.New()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if degrees[i][j] == 0 {
				q.PushBack([]int{i, j})
			}
		}
	}

	for q.Len() > 0 {
		res++
		lenth := q.Len()
		for i := 0; i < lenth; i++ {
			point := q.Remove(q.Front()).([]int)
			for _, dir := range dirs {
				x, y := point[0]+dir[0], point[1]+dir[1]
				if x >= 0 && x < rows && y >= 0 && y < cols && matrix[point[0]][point[1]] > matrix[x][y] {
					degrees[x][y]--
					if degrees[x][y] == 0 {
						q.PushBack([]int{x, y})
					}
				}
			}
		}
	}

	return res
}

// dfs
func longestIncreasingPathDfs(matrix [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rows, cols := len(matrix), len(matrix[0])

	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, cols)
	}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if memo[row][col] != 0 {
			return memo[row][col]
		}
		memo[row][col]++

		for _, dir := range dirs {
			x, y := row+dir[0], col+dir[1]
			if x >= 0 && x < rows && y >= 0 && y < cols && matrix[row][col] < matrix[x][y] {
				memo[row][col] = max(memo[row][col], dfs(x, y)+1)
			}
		}
		return memo[row][col]
	}

	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})) // 4
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})) // 4
	fmt.Println(longestIncreasingPath([][]int{{1}}))                             // 1

	fmt.Println(longestIncreasingPathDfs([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})) // 4
	fmt.Println(longestIncreasingPathDfs([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})) // 4
	fmt.Println(longestIncreasingPathDfs([][]int{{1}}))                             // 1
}
