package minSubArrayLen

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

