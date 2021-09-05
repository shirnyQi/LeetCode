package isCoverd

import (
	"fmt"
	"testing"
)

/*
1893. 检查是否区域内所有整数都被覆盖
给你一个二维整数数组 ranges 和两个整数 left 和 right 。每个 ranges[i] = [starti, endi] 表示一个从 starti 到 endi 的 闭区间 。
如果闭区间 [left, right] 内每个整数都被 ranges 中 至少一个 区间覆盖，那么请你返回 true ，否则返回 false 。
已知区间 ranges[i] = [starti, endi] ，如果整数 x 满足 starti <= x <= endi ，那么我们称整数x 被覆盖了。

示例 1：
输入：ranges = [[1,2],[3,4],[5,6]], left = 2, right = 5
输出：true
解释：2 到 5 的每个整数都被覆盖了：
- 2 被第一个区间覆盖。
- 3 和 4 被第二个区间覆盖。
- 5 被第三个区间覆盖。

示例 2：
输入：ranges = [[1,10],[10,20]], left = 21, right = 21
输出：false
解释：21 没有被任何一个区间覆盖。
*/

// 暴力求解
func isCovered(ranges [][]int, left int, right int) bool {
	for left <= right {
		have := false
		for _, v := range ranges {
			if left >= v[0] && left <= v[1] {
				have = true
				break
			}
		}
		if have == false {
			return false
		}
		left++
	}
	return true
}

// 差分数组
func isCovered2(ranges [][]int, left, right int) bool {
	diff := [52]int{} // 差分数组
	for _, r := range ranges {
		diff[r[0]]++
		diff[r[1]+1]--
	}
	cnt := 0 // 前缀和
	for i := 1; i <= 50; i++ {
		cnt += diff[i]
		if cnt <= 0 && left <= i && i <= right {
			return false
		}
	}
	return true
}

func TestName(t *testing.T) {
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5)) // true
	fmt.Println(isCovered([][]int{{1, 10}, {10, 20}}, 21, 21))    // false

	fmt.Println(isCovered2([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5)) // true
	fmt.Println(isCovered2([][]int{{1, 10}, {10, 20}}, 21, 21))    // false
}
