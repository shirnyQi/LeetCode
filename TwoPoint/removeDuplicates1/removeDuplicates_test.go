package removeDuplicates1

import (
	"fmt"
	"testing"
)

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/owkrrm/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
var (
	nums1 = []int{1, 1, 2}
	nums2 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
)

func removeDuplicates(nums []int) int {
	fmt.Println(nums)
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	fmt.Println(nums)
	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
}
