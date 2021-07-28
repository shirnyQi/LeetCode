package moveZeroes

import (
	"fmt"
	"testing"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

*/

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	slow, fast := 0, 0
	for slow < len(nums) {
		if fast == len(nums) && slow < fast {
			nums[slow] = 0
			slow++
			continue
		}
		if nums[fast] == 0 {
			fast++
			continue
		}

		nums[slow] = nums[fast]
		fast++
		slow++
	}
	fmt.Println(nums)
}

func Test_moveZeroes(t *testing.T) {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 1})
}
