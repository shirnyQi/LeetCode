package findLengthOfLCIS

import (
	"fmt"
	"testing"
)

var (
	nums1 = []int{1, 3, 5, 4, 7}
	nums2 = []int{2, 2, 2, 2, 2}
)

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var res, tmp = 1, 1

	var pre, next int = 0, 1
	for pre < len(nums) && next < len(nums) {
		if nums[pre] < nums[next] {
			tmp++
		} else {
			res = max(res, tmp)
			tmp = 1
		}
		pre = next
		next++
	}

	return max(res, tmp)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Test_findLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS(nums1)) //3
	fmt.Println(findLengthOfLCIS(nums2)) //1
}
