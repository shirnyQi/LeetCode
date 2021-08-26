package minNumber

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

/*
剑指 Offer 45. 把数组排成最小的数
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。



示例 1:

输入: [10,2]
输出: "102"
示例 2:

输入: [3,30,34,5,9]
输出: "3033459"

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sort-algorithms/oshq72/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return strconv.Itoa(a)+strconv.Itoa(b) < strconv.Itoa(b)+strconv.Itoa(a)

	})
	var res = ""
	for _, v := range nums {
		res += strconv.Itoa(v)
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(minNumber([]int{10, 2}))           // 102
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9})) // 3033459
}
