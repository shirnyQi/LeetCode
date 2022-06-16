package lengthOfLongestSubstringTwoDistinct

import (
	"fmt"
	"testing"
)

/*
159. 至多包含两个不同字符的最长子串

给定一个字符串 s ，找出 至多 包含两个不同字符的最长子串 t ，并返回该子串的长度。

示例 1:
输入: "eceba"
输出: 3
解释: t 是 "ece"，长度为3。

示例 2:
输入: "ccaabbb"
输出: 5
解释: t 是 "aabbb"，长度为5。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLongestSubstringTwoDistinct(s string) int {
	array := []byte(s)
	m := map[byte]int{}

	var res, left int
	for right, v := range array {
		m[v]++
		for len(m) > 2 {
			m[array[left]]--
			if m[array[left]] == 0 {
				delete(m, array[left])
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))   //3
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb")) //5
}
