package subarraysWithKDistinct

import (
	"fmt"
	"testing"
)

/*
K 个不同整数的子数组

给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定不同的子数组为好子数组。
（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
返回 A 中好子数组的数目。

示例 1：

输入：A = [1,2,1,2,3], K = 2
输出：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].

示例 2：

输入：A = [1,2,1,3,4], K = 3
输出：3
解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/riq25e/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func subarraysWithKDistinct(nums []int, k int) (ans int) {
	n := len(nums)
	num1 := make([]int, n+1)
	num2 := make([]int, n+1)

	// 第一个左指针表示极大的包含 k 个不同整数的区间的左端点，
	// 第二个左指针表示极大的包含 k-1 个不同整数的区间的左端点。
	// 而恰好存在 K 个不同整数的子区间的个数 = 最多存在 K 个不同整数的子区间的个数-最多存在 K-1 个不同整数的子区间的个数
	var tot1, tot2, left1, left2 int
	for _, v := range nums {
		if num1[v] == 0 {
			tot1++
		}
		num1[v]++
		if num2[v] == 0 {
			tot2++
		}
		num2[v]++
		for tot1 > k {
			num1[nums[left1]]--
			if num1[nums[left1]] == 0 {
				tot1--
			}
			left1++
		}
		for tot2 > k-1 {
			num2[nums[left2]]--
			if num2[nums[left2]] == 0 {
				tot2--
			}
			left2++
		}
		ans += left2 - left1
	}
	return ans
}

func TestName(t *testing.T) {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2)) //7
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3)) //3
}
