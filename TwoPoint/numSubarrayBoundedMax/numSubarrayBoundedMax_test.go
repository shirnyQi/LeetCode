package numSubarrayBoundedMax

import (
	"fmt"
	"testing"
)

/*
795. 区间子数组个数
给定一个元素都是正整数的数组A ，正整数 L 以及 R (L <= R)。
求连续、非空且其中最大元素满足大于等于L 小于等于R的子数组个数。

例如 :
输入:
A = [2, 1, 4, 3]
L = 2
R = 3
输出: 3
解释: 满足条件的子数组: [2], [2, 1], [3].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-subarrays-with-bounded-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	cnt := count(nums, right) - count(nums, left-1)
	return cnt
}

func count(A []int, num int) int {
	ans := 0
	cur := 0
	for _, v := range A {
		if v <= num {
			cur += 1
		} else {
			cur = 0
		}
		ans += cur
	}
	return ans
}

func TestName(t *testing.T) {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))                      //3
	fmt.Println(numSubarrayBoundedMax([]int{1, 2, 3, 2, 1}, 2, 3))                   //13
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3, 5, 1, 2, 3, 2, 1, 4}, 2, 3)) //16
}
