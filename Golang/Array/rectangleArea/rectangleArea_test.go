package rectangleArea

import (
	"fmt"
	"sort"
	"testing"
)

/*
850. 矩形面积 II
我们给出了一个（轴对齐的）矩形列表 rectangles 。 对于 rectangle[i] = [x1, y1, x2, y2]，其中（x1，y1）是矩形 i 左下角的坐标，（x2，y2）是该矩形右上角的坐标。
找出平面中所有矩形叠加覆盖后的总面积。 由于答案可能太大，请返回它对 10 ^ 9 + 7 取模的结果。

示例 1：
输入：[[0,0,2,2],[1,0,2,3],[1,0,3,1]]
输出：6
解释：如图所示。

示例 2：
输入：[[0,0,1000000000,1000000000]]
输出：49
解释：答案是 10^18 对 (10^9 + 7) 取模的结果， 即 (10^9)^2 → (-7)^2 = 49 。

提示：
1 <= rectangles.length <= 200
rectanges[i].length = 4
0 <= rectangles[i][j] <= 10^9
矩形叠加覆盖后的总面积不会超越 2^63 - 1 ，这意味着可以用一个 64 位有符号整数来保存面积结果。

link:https://leetcode-cn.com/problems/rectangle-area-ii/submissions/
*/

// 暴力解超出内存
func rectangleAreaOutMemery(rectangles [][]int) int {
	//mod := 1e9 + 7

	grid := make([][]int, 100, 100)

	for i := 0; i < 100; i++ {
		grid[i] = make([]int, 100, 100)
	}

	var res int

	for _, rect := range rectangles {
		x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				if grid[i][j] == 1 {
					continue
				}
				grid[i][j] = 1
				res++
			}
		}
	}

	return res
}

// 先切再拼
func rectangleArea(rectangles [][]int) int {
	const mod = 1e9 + 7

	xSplit, ySplit := make([]int, 0, len(rectangles)), make([]int, 0, len(rectangles))
	for _, v := range rectangles {
		xSplit = appendNoDouble(xSplit, v[0])
		ySplit = appendNoDouble(ySplit, v[1])
		xSplit = appendNoDouble(xSplit, v[2])
		ySplit = appendNoDouble(ySplit, v[3])
	}
	sort.Ints(xSplit)
	sort.Ints(ySplit)
	area := 0
	for i := 0; i+1 < len(xSplit); i++ {
		for j := 0; j+1 < len(ySplit); j++ {
			area += getArea(rectangles, xSplit[i], ySplit[j], xSplit[i+1], ySplit[j+1])
			area %= mod
		}
	}
	return area

}

func getArea(rectangles [][]int, x1 int, y1 int, x2 int, y2 int) int {
	for _, v := range rectangles {
		if v[0] <= x1 && v[1] <= y1 && v[2] >= x2 && v[3] >= y2 {
			return (x2 - x1) * (y2 - y1)
		}
	}
	return 0
}

func appendNoDouble(split []int, i int) []int {
	for _, v := range split {
		if v == i {
			return split
		}
	}
	return append(split, i)
}

func TestName(t *testing.T) {
	fmt.Println(rectangleAreaOutMemery([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}))

	fmt.Println(rectangleArea([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}))
	fmt.Println(rectangleArea([][]int{{0, 0, 1000000000, 1000000000}}))
}
