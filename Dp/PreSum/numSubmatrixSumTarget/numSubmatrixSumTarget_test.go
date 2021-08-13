package numSubmatrixSumTarget

import (
	"fmt"
	"testing"
)

/*
1074. 元素和为目标值的子矩阵数量
给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。
子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。
如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。

示例 1：
输入：matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
输出：4
解释：四个只含 0 的 1x1 子矩阵。

示例 2：
输入：matrix = [[1,-1],[-1,1]], target = 0
输出：5
解释：两个 1x2 子矩阵，加上两个 2x1 子矩阵，再加上一个 2x2 子矩阵。

示例 3：
输入：matrix = [[904]], target = 0
输出：0

link:https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/
*/

func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{0: 1}
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			ans += mp[pre-k]
		}
		mp[pre]++
	}
	return
}

func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	for i := range matrix { // 枚举上边界
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				sum[c] += v // 更新每列的元素和
			}
			ans += subarraySum(sum, target)
		}
	}
	return
}

func TestName(t *testing.T) {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0)) // 4
	fmt.Println(numSubmatrixSumTarget([][]int{{1, -1}, {-1, 1}}, 0))                // 5
	fmt.Println(numSubmatrixSumTarget([][]int{{904}}, 0))                           // 0
}
