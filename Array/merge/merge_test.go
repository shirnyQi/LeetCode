package merge

import (
	"fmt"
	"sort"
	"testing"
)

/*
合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。


示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/array-and-string/c5tv3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	size := len(intervals)
	if size == 1 {
		return intervals
	}

	var res [][]int
	start, end, index := intervals[0][0], intervals[0][1], 1
	res = append(res, []int{start, end})
	for index < size {
		if end >= intervals[index][0] {
			end = max(end, intervals[index][1])
			res[len(res)-1][1] = end
		} else {
			start, end = intervals[index][0], intervals[index][1]
			res = append(res, []int{start, end})
		}
		index++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1,5]]
}
