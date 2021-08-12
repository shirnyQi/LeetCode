package checkSubarraySum

import (
	"fmt"
	"testing"
)

/*
连续的子数组和
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。
如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。

示例 1：
输入：nums = [23,2,4,6,7], k = 6
输出：true
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。

示例 2：
输入：nums = [23,2,6,4,7], k = 6
输出：true
解释：[23, 2, 6, 4, 7] 是大小为 5 的子数组，并且和为 42 。
42 是 6 的倍数，因为 42 = 7 * 6 且 7 是一个整数。

示例 3：
输入：nums = [23,2,6,4,7], k = 13
输出：false

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5mb0mt/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func checkSubarraySum(nums []int, k int) bool {
	size := len(nums)
	var preSum = make([]int, size+1, size+1)
	preSum[0] = 0

	mIndex := map[int]int{}
	mIndex[0] = 0

	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		decimal := preSum[i+1] % k

		index, ok := mIndex[decimal]
		if !ok {
			mIndex[decimal] = i + 1
		} else {
			if i != index {
				return true
			}
		}
	}

	return false
}

func TestName(t *testing.T) {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))  // true
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))  // true
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13)) // false
}
