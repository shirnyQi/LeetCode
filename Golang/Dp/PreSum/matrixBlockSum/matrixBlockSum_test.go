package matrixBlockSum

import (
	"fmt"
	"testing"
)

/*
矩阵区域和
给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和：
i - k <= r <= i + k,
j - k <= c <= j + k 且
(r, c) 在矩阵内。

示例 1：
输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
输出：[[12,21,16],[27,45,33],[24,39,28]]

示例 2：
输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
输出：[[45,45,45],[45,45,45],[45,45,45]]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5aegb7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	P := make([][]int, m+1, m+1)
	P[0] = make([]int, n+1, n+1)
	for i := 1; i <= m; i++ {
		P[i] = make([]int, n+1, n+1)
		for j := 1; j <= n; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + mat[i-1][j-1]
		}
	}

	res := make([][]int, m, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			res[i][j] = get(P, m, n, i+k+1, j+k+1) -
				get(P, m, n, i-k, j+k+1) -
				get(P, m, n, i+k+1, j-k) +
				get(P, m, n, i-k, j-k)
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func get(P [][]int, m, n, x, y int) int {
	x1 := max(min(m, x), 0)
	y1 := max(min(n, y), 0)
	return P[x1][y1]
}

func TestName(t *testing.T) {
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)) // [[12,21,16],[27,45,33],[24,39,28]]
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2)) // [[45,45,45],[45,45,45],[45,45,45]]
}
