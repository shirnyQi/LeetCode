package findMaxLength

import (
	"fmt"
	"testing"
)

/*
连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1:
输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。

示例 2:
输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dynamic-programming-1-plus/5mzb7c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// 暴力超时解法，垃圾
func findMaxLengthTimeOut(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	preSum := make([]int, size+1, size+1)
	preSum[0] = 0
	var res int
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		for j := 0; j < i; j++ {
			count1 := preSum[i+1] - preSum[j]
			count2 := i + 1 - j - count1
			if count1 == count2 {
				res = max(res, i+1-j)
			}
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

func findMaxLength(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	m := map[int]int{}
	m[0] = -1
	var res, count int
	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}
		index, ok := m[count]
		if !ok {
			m[count] = i
		} else {
			res = max(res, i-index)
		}
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(findMaxLength([]int{0, 1}))                // 2
	fmt.Println(findMaxLength([]int{0, 1, 0}))             // 2
	fmt.Println(findMaxLength([]int{0, 1, 0, 0}))          // 2
	fmt.Println(findMaxLength([]int{0, 1, 0, 0, 1}))       // 4
	fmt.Println(findMaxLength([]int{0, 1, 0, 0, 1, 1}))    // 6
	fmt.Println(findMaxLength([]int{0, 1, 0, 0, 1, 1, 1})) // 6

	fmt.Println("====")

	fmt.Println(findMaxLengthTimeOut([]int{0, 1}))                // 2
	fmt.Println(findMaxLengthTimeOut([]int{0, 1, 0}))             // 2
	fmt.Println(findMaxLengthTimeOut([]int{0, 1, 0, 0}))          // 2
	fmt.Println(findMaxLengthTimeOut([]int{0, 1, 0, 0, 1}))       // 4
	fmt.Println(findMaxLengthTimeOut([]int{0, 1, 0, 0, 1, 1}))    // 6
	fmt.Println(findMaxLengthTimeOut([]int{0, 1, 0, 0, 1, 1, 1})) // 6
}
