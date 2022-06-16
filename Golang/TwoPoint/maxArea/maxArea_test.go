package maxArea

import (
	"fmt"
	"testing"
)

/*
盛最多水的容器

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：
输入：height = [1,1]
输出：1

示例 3：
输入：height = [4,3,2,1,4]
输出：16

示例 4：
输入：height = [1,2,1]
输出：2

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/od6y65/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	m := 0
	for i < j {
		// 计算当前最大面积
		cur := (j - i) * min(height[i], height[j])
		if cur > m {
			m = cur
		}

		// 移动较小的一侧指针
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func TestName(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //49
	fmt.Println(maxArea([]int{1, 1}))                      //1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             //16
	fmt.Println(maxArea([]int{1, 2, 1}))                   //2
}
