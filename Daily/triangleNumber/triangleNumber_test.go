package triangleNumber

import (
	"fmt"
	"sort"
	"testing"
)

/*
给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
示例 1:
输入: [2,2,3,4]
输出: 3
解释:
有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3
注意:
数组长度不超过1000。
数组里整数的范围为 [0, 1000]。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-triangle-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 二分法：先将数组递增排序，接着用二分法快速找到不满足的下标，提升速度
func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)

	var res int

	size := len(nums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			left, k, right := j+1, j, size-1
			for left <= right {
				mid := (left + right) / 2
				if nums[i]+nums[j] > nums[mid] {
					left = mid + 1
					k = mid
				} else {
					right = mid - 1
				}
			}
			res += k - j
		}
	}
	return res
}

// 双指针：固定一个边，另外两个边看成双指针挨个迁移有序数组
func triangleNumber2(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v+nums[j] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))  // 3
	fmt.Println(triangleNumber2([]int{2, 2, 3, 4})) // 3
}
