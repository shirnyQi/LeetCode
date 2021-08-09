package minSwaps

import (
	"fmt"
	"math"
	"testing"
)

/*

给出一个二进制数组 data，你需要通过交换位置，将数组中 任何位置 上的 1 组合到一起，并返回所有可能中所需 最少的交换次数。

示例 1：
	输入：[1,0,1,0,1]
	输出：1
	解释：
	有三种可能的方法可以把所有的 1 组合在一起：
	[1,1,1,0,0]，交换 1 次；
	[0,1,1,1,0]，交换 2 次；
	[0,0,1,1,1]，交换 1 次。
	所以最少的交换次数为 1。

示例 2：
	输入：[0,0,0,1,0]
	输出：0
	解释：
	由于数组中只有一个 1，所以不需要交换。

示例 3：
	输入：[1,0,1,0,1,0,0,1,1,0,1]
	输出：3
	解释：
	交换 3 次，一种可行的只用 3 次交换的解决方案是 [0,0,0,0,0,1,1,1,1,1,1]。

https://leetcode-cn.com/problems/minimum-swaps-to-group-all-1s-together/
*/

func minSwaps(data []int) int {
	var res = math.MaxInt32

	// 先统计有多少个1，以此作为窗口
	var windows int
	for _, v := range data {
		windows += v
	}

	var count int
	for left, right := 0, 0; right < len(data); right++ {
		count += data[right]
		if right-left+1 == windows {
			if res > windows-count {
				res = windows - count
			}
			count -= data[left]
			left++
		}
	}
	if res == math.MaxInt32 {
		res = 0
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(minSwaps([]int{1, 0, 1, 0, 1}))                   // 1
	fmt.Println(minSwaps([]int{0, 0, 0, 1, 0}))                   // 0
	fmt.Println(minSwaps([]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1})) // 3
}
