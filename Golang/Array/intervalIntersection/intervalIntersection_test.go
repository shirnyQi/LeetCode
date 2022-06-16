package intervalIntersection

import (
	"fmt"
	"testing"
)

/*
986. 区间列表的交集
给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
返回这 两个区间列表的交集 。
形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

示例 1：
输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

示例 2：
输入：firstList = [[1,3],[5,9]], secondList = []
输出：[]

示例 3：
输入：firstList = [], secondList = [[4,8],[10,12]]
输出：[]

示例 4：
输入：firstList = [[1,7]], secondList = [[3,10]]
输出：[[3,7]]

link:https://leetcode-cn.com/problems/interval-list-intersections/
*/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int

	i, j := 0, 0
	beg, end := 0, 0
	for i < len(firstList) && j < len(secondList) {
		f, s := firstList[i], secondList[j]

		if s[0] > f[1] {
			i++
			continue
		}

		if s[1] < f[0] {
			j++
			continue
		}

		if f[0] <= s[0] {
			beg = s[0]
			if f[1] <= s[1] {
				end = f[1]
				i++
			} else {
				end = s[1]
				j++
			}
		} else {
			beg = f[0]
			if f[1] <= s[1] {
				end = f[1]
				i++
			} else {
				end = s[1]
				j++
			}
		}
		res = append(res, []int{beg, end})
	}

	return res
}

func TestName(t *testing.T) {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
	fmt.Println(intervalIntersection([][]int{}, [][]int{{4, 8}, {10, 12}}))
	fmt.Println(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}))
}
