package maxSatisfied

import (
	"fmt"
	"testing"
)

/*
今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。

在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。

书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。

请你返回这一天营业下来，最多有多少客户能够感到满意。


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/rlkdli/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var (
	customers = []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy    = []int{0, 1, 0, 1, 0, 1, 0, 1}
	X         = 3
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	size := len(customers)

	preNums := make([]int, size+1, size+1)
	origin := 0
	for i := 0; i < size; i++ {
		origin += customers[i] * (1 - grumpy[i])
		preNums[i+1] = preNums[i] + customers[i]*grumpy[i]
	}
	fmt.Println(origin, preNums)

	maxAngryCount := 0
	for left := 0; left < size-minutes+1; left++ {
		maxAngryCount = max(maxAngryCount, preNums[left+minutes]-preNums[left])
	}
	return origin + maxAngryCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_(t *testing.T) {
	fmt.Println(maxSatisfied(customers, grumpy, X)) // 16
	//fmt.Println(maxSatisfied([]int{2, 6, 6, 9}, []int{0, 0, 1, 1}, 1)) // 17
}
