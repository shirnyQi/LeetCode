package removeDuplicates2

import (
	"fmt"
	"testing"
)

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/owzb55/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var (
	nums1 = []int{1, 1, 1, 2, 2, 3}
	nums2 = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	nums3 = []int{1, 1, 1, 1, 2, 2, 3}
)

func removeDuplicates(nums []int) int {
	fmt.Println(nums)
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	repeat := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
			repeat = 1
		} else {
			repeat++
			if repeat == 2 {
				nums[slow+1] = nums[slow]
				slow++
			}
		}
	}
	if repeat == 2 && nums[len(nums)-1] == nums[slow-1] {
		nums[slow] = nums[len(nums)-1]
	}
	fmt.Println(nums)

	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(removeDuplicates(nums3))
}
