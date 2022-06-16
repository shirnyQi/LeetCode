package findAnagrams

import (
	"fmt"
	"testing"
)

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指字母相同，但排列不同的字符串。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 错位解法
func findAnagrams(s string, p string) []int {
	var res []int

	targetMap := map[byte]int{}
	for _, s := range []byte(p) {
		targetMap[s]++
	}

	left, right := 0, 0

	charArray := []byte(s)
	for ; right < len(charArray); right++ {
		_, ok1 := targetMap[charArray[right]]
		if ok1 {
			targetMap[charArray[right]]--
			if targetMap[charArray[right]] >= 0 {
				continue
			}
		}
		fmt.Println("1", targetMap, left, right)
		for left < right {
			if _, ok2 := targetMap[charArray[left]]; ok2 {
				targetMap[charArray[left]]++
			}
			left++
		}

		fmt.Println("2", targetMap, left, right)
		if check(targetMap, p) {
			res = append(res, left-(len(p)))
		}
	}

	return res
}

func check(m map[byte]int, p string) bool {
	m0 := map[byte]int{}
	for _, s := range []byte(p) {
		m0[s]++
	}

	for k, v := range m {
		if v != m0[k] {
			return false
		}
	}
	return true
}

func Test_findAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("abab", "ab"))        //[0,1,2]

	fmt.Println(findAnagrams2("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams2("abab", "ab"))        //[0,1,2]

	fmt.Println(findAnagrams3("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams3("abab", "ab"))        //[0,1,2]
}

// 暴力解法
func findAnagrams2(s string, p string) []int {
	var res []int

	targetMap := map[byte]int{}
	for _, s := range []byte(p) {
		targetMap[s]++
	}

	left, right := 0, -1

	charArray := []byte(s)
	for left < len(charArray) {
		if right+1 < len(charArray) && right-left+1 < len(p) {
			right++
		} else {
			left++
		}
		if right-left+1 == len(p) && checkString(s, left, right, targetMap) {
			res = append(res, left)
		}
	}
	return res
}

func checkString(s string, left, right int, target map[byte]int) bool {
	m0 := map[byte]int{}
	for i := left; i <= right; i++ {
		m0[[]byte(s)[i]]++
	}

	for k, v := range m0 {
		if v != target[k] {
			return false
		}
	}
	return true
}

// O(n)解法
func findAnagrams3(s string, p string) []int {
	pm := map[byte]int{}
	for i := 0; i < len(p); i++ {
		pm[p[i]]++
	}
	sm := map[byte]int{}
	var match int
	var ans []int
	for l, r := 0, 0; r < len(s); r++ {
		sm[s[r]]++
		if _, ok := pm[s[r]]; ok {
			if pm[s[r]] == sm[s[r]] {
				match++
			}
			if match == len(pm) {
				ans = append(ans, l)
			}
		}
		if r-l+1 == len(p) {
			if _, ok := pm[s[l]]; ok && sm[s[l]] == pm[s[l]] {
				match--
			}
			sm[s[l]]--
			l++
		}
	}
	return ans
}
