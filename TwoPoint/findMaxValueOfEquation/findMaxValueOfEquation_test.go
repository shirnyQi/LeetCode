package findMaxValueOfEquation

import (
	"fmt"
	"math"
	"testing"
)

/*
1499. 满足不等式的最大值
给你一个数组 points 和一个整数 k 。数组中每个元素都表示二维平面上的点的坐标，并按照横坐标 x 的值从小到大排序。也就是说 points[i] = [xi, yi] ，并且在 1 <= i < j <= points.length 的前提下， xi < xj 总成立。
请你找出 yi + yj + |xi - xj| 的 最大值，其中 |xi - xj| <= k 且 1 <= i < j <= points.length。
题目测试数据保证至少存在一对能够满足 |xi - xj| <= k 的点。

示例 1：
输入：points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
输出：4
解释：前两个点满足 |xi - xj| <= 1 ，代入方程计算，则得到值 3 + 0 + |1 - 2| = 4 。第三个和第四个点也满足条件，得到值 10 + -10 + |5 - 6| = 1 。
没有其他满足条件的点，所以返回 4 和 1 中最大的那个。

示例 2：
输入：points = [[0,0],[3,0],[9,2]], k = 3
输出：3
解释：只有前两个点满足 |xi - xj| <= 3 ，代入方程后得到值 0 + 0 + |0 - 3| = 3 。

link:https://leetcode-cn.com/problems/max-value-of-equation/
*/

// 暴力解超时
func findMaxValueOfEquationTimeOut(points [][]int, k int) int {
	res := math.MinInt32
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			if abs(x2, x1) > k {
				break
			}
			res = max(res, y1+y2+abs(x2, x1))
		}
	}

	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 滑动窗口
func findMaxValueOfEquationWindows(points [][]int, k int) int {
	res := math.MinInt32
	l, r := 0, 1

	for r < len(points) {
		if l >= r {
			r++
			continue
		}
		x1, y1 := points[l][0], points[l][1]
		x2, y2 := points[r][0], points[r][1]

		if x2-x1 > k {
			l++
			continue
		}
		res = max(res, y1+y2+x2-x1)

		if y2-y1 > x2-x1 {
			l++
			continue
		}
		r++
	}
	r--
	for l < r {
		l++
		x1, y1 := points[l][0], points[l][1]
		x2, y2 := points[r][0], points[r][1]
		if x2-x1 > k {
			l++
			continue
		}
		res = max(res, y1+y2+x2-x1)
	}
	return res
}

// 队列
//i,j对称，不妨设j>i，xj-xi <= k，求 xj+yj+yi-xi
//维护单调递减队列，可以用O(1)的时间查到当前窗口的yi-xi最大值，i即队列头
func findMaxValueOfEquation(points [][]int, k int) int {
	queue := make([]int, 0)
	ret := math.MinInt64
	for j, point := range points {
		//不满足条件xj-xi<= k，删除队列头，跳到下一个
		for len(queue) != 0 && point[0]-points[queue[0]][0] > k {
			queue = queue[1:]
		}
		//更新ret，队列头元素i对应当前yi-xi最大值
		if len(queue) != 0 {
			//xj+yj+yi-xi
			tmp := point[0] + point[1] + points[queue[0]][1] - points[queue[0]][0]
			if tmp > ret {
				ret = tmp
			}
		}
		//维护单调减队列
		diff := point[1] - point[0]
		for len(queue) != 0 && diff > points[queue[len(queue)-1]][1]-points[queue[len(queue)-1]][0] {
			queue = queue[:len(queue)-1]
		}
		//无论如何，当前点入队列
		queue = append(queue, j)
	}
	return ret
}

func TestName(t *testing.T) {
	fmt.Println(findMaxValueOfEquationWindows([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
	fmt.Println(findMaxValueOfEquationWindows([][]int{{0, 0}, {3, 0}, {9, 2}}, 3))

	fmt.Println(findMaxValueOfEquationTimeOut([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
	fmt.Println(findMaxValueOfEquationTimeOut([][]int{{0, 0}, {3, 0}, {9, 2}}, 3))

	fmt.Println(findMaxValueOfEquation([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
	fmt.Println(findMaxValueOfEquation([][]int{{0, 0}, {3, 0}, {9, 2}}, 3))
}
