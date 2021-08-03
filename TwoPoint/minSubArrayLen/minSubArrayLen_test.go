package minSubArrayLen

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

 

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/ri9ec6/
*/
import (
	"fmt"
	"math"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	var res, count int
	res = math.MaxInt32
	for left, right := 0, 0; right < len(nums) && left < len(nums); {
		count += nums[right]
		right++
		for count >= target {
			count -= nums[left]
			left++
			res = min(res, right-left+1)
		}
	}
	if res == math.MaxInt32 {
		res = 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          // 3
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}

