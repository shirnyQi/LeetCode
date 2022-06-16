package equalSubstring

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你两个长度相同的字符串，s 和 t。
将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。
如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

示例 1：
输入：s = "abcd", t = "bcdf", maxCost = 3
输出：3
解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。

示例 2：
输入：s = "abcd", t = "cdef", maxCost = 3
输出：1
解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。

示例 3：
输入：s = "abcd", t = "acde", maxCost = 0
输出：1
解释：a -> a, cost = 0，字符串未发生变化，所以最大长度为 1。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/rim36v/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type function interface {
	equalSubstring(s string, t string, maxCost int) int
}

type windows struct {
}

// 双指针法
func (w *windows) equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	diff := make([]int, n)
	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}

	var res, sum, left int

	for right, cost := range diff {
		sum += cost
		// sum相当于窗口，right一直右移，left先不动，直到找到窗口的边界值，此时移动左边
		for sum > maxCost {
			sum -= diff[left]
			left++
		}
		// 每移动一次比较一次最大值，相比于未找到窗口边界值一直continue这种操作而言，
		// 不需要考虑由于一直continue导致for循环退出的问题
		res = max(res, right-left+1)
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type binarySearch struct {
}

// 前缀和+二分查找
func (b *binarySearch) equalSubstring(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	accDiff := make([]int, n+1)
	for i, ch := range s {
		accDiff[i+1] = accDiff[i] + abs(int(ch)-int(t[i]))
	}
	for i := 1; i <= n; i++ {
		start := sort.SearchInts(accDiff[:i], accDiff[i]-maxCost)
		maxLen = max(maxLen, i-start)
	}
	return
}

func TestName(t *testing.T) {
	w := &windows{}
	fmt.Println(w.equalSubstring("abcd", "bcdf", 3)) //3
	fmt.Println(w.equalSubstring("abcd", "cdef", 3)) //1
	fmt.Println(w.equalSubstring("abcd", "acde", 0)) //1

	b := &binarySearch{}
	fmt.Println(b.equalSubstring("abcd", "bcdf", 3)) //3
	fmt.Println(b.equalSubstring("abcd", "cdef", 3)) //1
	fmt.Println(b.equalSubstring("abcd", "acde", 0)) //1
}
