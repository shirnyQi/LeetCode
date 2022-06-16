package productExceptSelf

import (
	"fmt"
	"testing"
)

/*
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:
输入: [1,2,3,4]
输出: [24,12,8,6]

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5ahred/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func productExceptSelf(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nil
	}

	var preMulti = make([]int, size+1, size+1)
	preMulti[0] = 1

	for i, v := range nums {
		preMulti[i+1] = preMulti[i] * v
	}

	t := 1
	for i := size - 1; i >= 0; i-- {
		preMulti[i] = preMulti[i] * t
		t = t * nums[i]
	}
	return preMulti[:size]
}

func TestName(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) // [24,12,8,6]
}
