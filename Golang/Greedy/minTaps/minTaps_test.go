package minTaps

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
1326. 灌溉花园的最少水龙头数目
在 x 轴上有一个一维的花园。花园长度为 n，从点 0 开始，到点 n 结束。
花园里总共有 n + 1 个水龙头，分别位于 [0, 1, ..., n] 。
给你一个整数 n 和一个长度为 n + 1 的整数数组 ranges ，其中 ranges[i] （下标从 0 开始）表示：如果打开点 i 处的水龙头，可以灌溉的区域为 [i -  ranges[i], i + ranges[i]] 。
请你返回可以灌溉整个花园的 最少水龙头数目 。如果花园始终存在无法灌溉到的地方，请你返回 -1 。

示例 1：
输入：n = 5, ranges = [3,4,1,1,0,0]
输出：1
解释：
点 0 处的水龙头可以灌溉区间 [-3,3]
点 1 处的水龙头可以灌溉区间 [-3,5]
点 2 处的水龙头可以灌溉区间 [1,3]
点 3 处的水龙头可以灌溉区间 [2,4]
点 4 处的水龙头可以灌溉区间 [4,4]
点 5 处的水龙头可以灌溉区间 [5,5]
只需要打开点 1 处的水龙头即可灌溉整个花园 [0,5] 。

示例 2：
输入：n = 3, ranges = [0,0,0,0]
输出：-1
解释：即使打开所有水龙头，你也无法灌溉整个花园。

示例 3：
输入：n = 7, ranges = [1,2,1,0,2,1,0,1]
输出：3

示例 4：
输入：n = 8, ranges = [4,0,0,0,0,0,0,0,4]
输出：2

示例 5：
输入：n = 8, ranges = [4,0,0,0,4,0,0,0,4]
输出：1

link:https://leetcode-cn.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
*/

// 个人错误解答
func minTapsFail(n int, ranges []int) int {
	var cover [][]int

	for i, v := range ranges {
		cover = append(cover, []int{i - v, i + v})
	}

	sort.Slice(cover, func(i, j int) bool {
		return cover[i][0] < cover[j][0]
	})

	if cover[0][0] > 0 {
		return -1
	}

	var l, r, res int

	for _, v := range cover {
		beg, end := v[0], v[1]

		if beg <= l && end > r {
			r = end
			if r >= n {
				res++
				break
			}
		} else if beg > l {
			res++
			l = r + 1
		}
	}

	if r < n {
		return -1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 贪心算法
func minTaps(n int, ranges []int) int {
	var pre = make([]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		pre[i] = i
	}

	for i := 0; i <= n; i++ {
		l := max(i-ranges[i], 0)
		r := min(i+ranges[i], n)
		pre[r] = min(pre[r], l)
	}

	breakpoint, furthest, ans := n, math.MaxInt32, 0
	for i := n; i > 0; i-- {
		furthest = min(furthest, pre[i])
		if i == breakpoint {
			if furthest >= i {
				return -1
			}
			ans++
			breakpoint = furthest
		}

	}

	return ans
}

func TestName(t *testing.T) {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))          // 1
	fmt.Println(minTaps(3, []int{0, 0, 0, 0}))                // -1
	fmt.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))    // 3
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4})) // 2
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4})) // 1
	fmt.Println(minTaps(35,
		[]int{1, 0, 4, 0, 4, 1, 4, 3, 1, 1, 1, 2, 1, 4, 0, 3, 0, 3, 0, 3, 0, 5, 3, 0, 0, 1, 2, 1, 2, 4, 3, 0, 1, 0, 5, 2})) // 6

	fmt.Println(minTapsFail(5, []int{3, 4, 1, 1, 0, 0}))          // 1
	fmt.Println(minTapsFail(3, []int{0, 0, 0, 0}))                // -1
	fmt.Println(minTapsFail(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))    // 3
	fmt.Println(minTapsFail(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4})) // 2
	fmt.Println(minTapsFail(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4})) // 1
	fmt.Println(minTapsFail(35,
		[]int{1, 0, 4, 0, 4, 1, 4, 3, 1, 1, 1, 2, 1, 4, 0, 3, 0, 3, 0, 3, 0, 5, 3, 0, 0, 1, 2, 1, 2, 4, 3, 0, 1, 0, 5, 2})) // 6
}
