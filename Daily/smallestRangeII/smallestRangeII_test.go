package smallestRangeII

import (
	"fmt"
	"sort"
	"testing"
)

/*
910. 最小差值 II
给你一个整数数组 A，对于每个整数 A[i]，可以选择 x = -K 或是 x = K （K 总是非负整数），并将 x 加到 A[i] 中。
在此过程之后，得到数组 B。
返回 B 的最大值和 B 的最小值之间可能存在的最小差值。

示例 1：
输入：A = [1], K = 0
输出：0
解释：B = [1]

示例 2：
输入：A = [0,10], K = 2
输出：6
解释：B = [2,8]

示例 3：
输入：A = [1,3,6], K = 3
输出：3
解释：B = [4,6,3]

link:https://leetcode-cn.com/problems/smallest-range-ii/
*/

func smallestRangeII(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)

	size := len(nums)
	nMin, nMax := nums[0], nums[size-1]
	res := nMax - nMin

	for i := 0; i < size-1; i++ {
		a, b := nums[i], nums[i+1]
		res = min(res, max(nMax-k, a+k)-min(nMin+k, b-k))
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

func TestName(t *testing.T) {
	fmt.Println(smallestRangeII([]int{1}, 0))       // 0
	fmt.Println(smallestRangeII([]int{0, 10}, 2))   // 6
	fmt.Println(smallestRangeII([]int{1, 3, 6}, 3)) // 3
}
