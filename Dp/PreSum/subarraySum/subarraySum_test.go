package subarraySum

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5m16vg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func subarraySum(nums []int, k int) int {
	size := len(nums)
	var preSum = make([]int, size+1, size+1)
	preSum[0] = 0

	mIndex := map[int]int{}
	mIndex[0] = 1
	var res int
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v

		target := preSum[i+1] - k
		if _, ok := mIndex[target]; ok {
			res += mIndex[target]
		}
		mIndex[preSum[i+1]]++
	}

	return res
}

func TestName(t *testing.T) {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2)) // 2
}
