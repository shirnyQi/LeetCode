package restoreArray

import (
	"fmt"
	"testing"
)

/*
存在一个由 n 个不同元素组成的整数数组 nums ，但你已经记不清具体内容。好在你还记得 nums 中的每一对相邻元素。

给你一个二维整数数组 adjacentPairs ，大小为 n - 1 ，其中每个 adjacentPairs[i] = [ui, vi] 表示元素 ui 和 vi 在 nums 中相邻。

题目数据保证所有由元素 nums[i] 和 nums[i+1] 组成的相邻元素对都存在于 adjacentPairs 中，存在形式可能是 [nums[i], nums[i+1]] ，也可能是 [nums[i+1], nums[i]] 。这些相邻元素对可以 按任意顺序 出现。

返回 原始数组 nums 。如果存在多种解答，返回 其中任意一个 即可。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func restoreArray(adjacentPairs [][]int) []int {
	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}

	nodeMap := make(map[int][]int, 0)
	var result []int

	var start int
	for _, pair := range adjacentPairs {
		if nodeMap[pair[0]] == nil {
			nodeMap[pair[0]] = []int{pair[1]}
		} else {
			nodeMap[pair[0]] = append(nodeMap[pair[0]], pair[1])
		}
		if nodeMap[pair[1]] == nil {
			nodeMap[pair[1]] = []int{pair[0]}
		} else {
			nodeMap[pair[1]] = append(nodeMap[pair[1]], pair[0])
		}
	}

	for k, v := range nodeMap {
		if len(v) == 1 {
			start = k
			break
		}
	}

	result = append(result, start)
	pre := start
	start = nodeMap[start][0]

	for len(nodeMap[start]) == 2 {
		result = append(result, start)
		nextList := nodeMap[start]
		if nextList[0] == pre {
			pre = start
			start = nextList[1]
		} else {
			pre = start
			start = nextList[0]
		}
	}
	result = append(result, start)

	return result
}

func Test_restoreArray(t *testing.T) {
	fmt.Println(restoreArray([][]int{{4, -10}, {-1, 3}, {4, -3}, {-3, 3}}))
}
