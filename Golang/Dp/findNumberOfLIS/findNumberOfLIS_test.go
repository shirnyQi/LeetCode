package findNumberOfLIS

import (
	"fmt"
	"testing"
)

/*
最长递增子序列的个数
给定一个未排序的整数数组，找到最长递增子序列的个数。

示例 1:
输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。

示例 2:
输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5o6mrv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func findNumberOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	dp := make([]int, size, size)
	count := make([]int, size, size)
	dp[0], count[0] = 1, 1

	var maxLenth = 1
	for i := 1; i < size; i++ {
		dp[i], count[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		maxLenth = max(maxLenth, dp[i])
		fmt.Println(dp, count)
	}

	var res int
	for i, v := range count {
		if dp[i] == maxLenth {
			res += v
		}
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
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7})) //2
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2})) //5
	fmt.Println(findNumberOfLIS([]int{1, 2}))          //1
}
