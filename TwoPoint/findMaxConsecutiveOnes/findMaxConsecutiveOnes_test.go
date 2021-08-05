package findMaxConsecutiveOnes

import (
	"fmt"
	"testing"
)

/*
给定一个二进制数组，你可以最多将 1 个 0 翻转为 1，找出其中最大连续 1 的个数。
示例 1：
输入：[1,0,1,1,0]
输出：4
解释：翻转第一个 0 可以得到最长的连续 1。
     当翻转以后，最大连续 1 的个数为 4。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/rirsze/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// standard
func findMaxConsecutiveOnes(nums []int) int {
	var res int

	target := 1
	left, right := 0, 0

	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			target--
		}

		for target < 0 {
			if nums[left] == 0 {
				target++
			}
			left++
		}
		res = max(res, right-left+1)
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
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0}))          //4
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1, 1, 1})) //6
}
