package maxSubArrayLen

import (
	"fmt"
	"math"
	"testing"
)

/*
325. 和等于 k 的最长子数组长度
给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。

示例 1:
输入: nums = [1,-1,5,-2,3], k = 3
输出: 4
解释: 子数组 [1, -1, 5, -2] 和等于 3，且长度最长。

示例 2:
输入: nums = [-2,-1,2,1], k = 1
输出: 2
解释: 子数组 [-1, 2] 和等于 1，且长度最长。

link:https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k/
*/

func maxSubArrayLen(nums []int, k int) int {
	size := len(nums)
	var preSum = make([]int, size+1, size+1)
	preSum[0] = 0

	mIndex := map[int]int{}
	mIndex[0] = 0
	var res = math.MinInt32
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		if _, ok := mIndex[preSum[i+1]]; !ok {
			mIndex[preSum[i+1]] = i + 1
		}
		target := preSum[i+1] - k
		if _, ok := mIndex[target]; ok {
			res = max(res, i-mIndex[target]+1)
		}
	}

	if res == math.MinInt32 {
		return 0
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
	fmt.Println(maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3)) // 4
	fmt.Println(maxSubArrayLen([]int{-2, -1, 2, 1}, 1))    // 2
}
