package largestValsFromLabels

import (
	"fmt"
	"sort"
	"testing"
)

/*
受标签影响的最大值
我们有一个项的集合，其中第 i 项的值为 values[i]，标签为 labels[i]。
我们从这些项中选出一个子集 S，这样一来：
|S| <= num_wanted
对于任意的标签 L，子集 S 中标签为 L 的项的数目总满足 <= use_limit。
返回子集 S 的最大可能的 和。

示例 1：
输入：values = [5,4,3,2,1], labels = [1,1,2,2,3], num_wanted = 3, use_limit = 1
输出：9
解释：选出的子集是第一项，第三项和第五项。

示例 2：
输入：values = [5,4,3,2,1], labels = [1,3,3,3,2], num_wanted = 3, use_limit = 2
输出：12
解释：选出的子集是第一项，第二项和第三项。

示例 3：
输入：values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 1
输出：16
解释：选出的子集是第一项和第四项。

示例 4：
输入：values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 2
输出：24
解释：选出的子集是第一项，第二项和第四项。

link:https://leetcode-cn.com/problems/largest-values-from-labels/
*/
type item struct {
	value int
	label int
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	size := len(values)
	items := make([]item, size, size)
	for i := 0; i < size; i++ {
		items[i] = item{values[i], labels[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})
	var res int
	var lenth, i int
	usedMap := make(map[int][]int, 0)
	for lenth < numWanted && i < size {
		value, label := items[i].value, items[i].label
		i++
		if len(usedMap[label]) < useLimit {
			res += value
			usedMap[label] = append(usedMap[label], value)
			lenth++
		}
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1)) // 9
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2)) // 12
	fmt.Println(largestValsFromLabels([]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 1)) // 16
	fmt.Println(largestValsFromLabels([]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 2)) // 24
}
