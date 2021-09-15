package largestRectangleArea

import (
	"fmt"
	"testing"
)

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。


示例 1:
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

示例 2：
输入： heights = [2,4]
输出： 4
*/

// 双指针解法
func largestRectangleAreaTwoPoint(heights []int) int {
	var maxArea int

	for crtIndex := range heights {
		if heights[crtIndex]*len(heights) <= maxArea {
			continue
		}
		var prefix, next int

		for i := crtIndex; i > -1; i-- {
			if heights[i] < heights[crtIndex] {
				prefix = i + 1
				break
			}
			prefix = i
		}

		for i := crtIndex; i < len(heights); i++ {
			if heights[i] < heights[crtIndex] {
				next = i - 1
				break
			}
			next = i
		}

		maxArea = max(maxArea, heights[crtIndex]*(next-prefix+1))
	}

	return maxArea

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 单调栈解法
func largestRectangleArea(heights []int) int {
	n := len(heights)

	// left: 找到当前位置开始，向左第一个比自身小的数的下标
	// right: 找到当前位置开始，向右第一个比自身小的数的下标

	left, right := make([]int, n), make([]int, n)
	mono_stack := []int{} // 单调递减栈
	for i := 0; i < n; i++ {
		// 如果栈顶元素比当前元素大或者相等，则出栈，保证严格递减
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}

		// 初始左侧下标为-1，而后保存每个栈顶的下标
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		// 下标压栈
		mono_stack = append(mono_stack, i)
	}

	mono_stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			right[i] = n
		} else {
			right[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}

	// 从当前位置开始，分别取左右各自比自己小的第一个数的下标，即为当前面积
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func TestName(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))

	fmt.Println(largestRectangleAreaTwoPoint([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleAreaTwoPoint([]int{2, 4}))
}
