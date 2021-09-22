package longestPalindromeSubseq

import (
	"fmt"
	"testing"
)

/*
1682. 最长回文子序列 II
字符串 s 的某个子序列符合下列条件时，称为“好的回文子序列”：
它是 s 的子序列。
它是回文序列（反转后与原序列相等）。
长度为偶数。
除中间的两个字符外，其余任意两个连续字符不相等。
例如，若 s = "abcabcabb"，则 "abba" 可称为“好的回文子序列”，而 "bcb" （长度不是偶数）和 "bbbb" （含有相等的连续字符）不能称为“好的回文子序列”。
给定一个字符串 s， 返回 s 的最长“好的回文子序列”的长度。


示例 1:
输入: s = "bbabab"
输出: 4
解释: s 的最长“好的回文子序列”是 "baab"。

示例 2:
输入: s = "dcbccacdb"
输出: 4
解释: The longest good palindromic subsequence of s is "dccd".
*/

// error
func longestPalindromeSubseqerror(s string) int {
	array := []byte(s)
	lastMap := make(map[byte]int, 0)
	var res int
	for i := range array {
		lastMap[array[i]] = i
		l, r := i, i+1
		m := make(map[byte]bool, 0)
		for l >= 0 && r < len(array) {
			if array[l] != array[r] {
				r++
				continue
			}
			if m[array[l]] {
				break
			}
			m[array[l]] = true
			if len(m)*2 > res {
				res = len(m) * 2
			}

			isFind := false
			for r < len(array) {
				if _, ok := lastMap[array[r]]; ok && l > lastMap[array[r]] {
					isFind = true
					l = lastMap[array[r]]
					break
				}
				r++
			}
			if !isFind {
				break
			}
		}

	}
	return res
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	memo := make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, 26)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dp func(i, j, c int) int
	dp = func(i, j, c int) int {
		if memo[i][j][c] != -1 {
			return memo[i][j][c]
		}
		if i >= j {
			memo[i][j][c] = 0
			return 0
		}
		if s[i] != byte('a'+c) {
			memo[i][j][c] = dp(i+1, j, c)
			return memo[i][j][c]
		}
		if s[j] != byte('a'+c) {
			memo[i][j][c] = dp(i, j-1, c)
			return memo[i][j][c]
		}
		res := -1
		for cc := 0; cc < 26; cc++ {
			if cc != c {
				res = max(res, 2+dp(i+1, j-1, cc))
			}
		}
		memo[i][j][c] = res
		return res
	}
	ans := -1
	for cc := 0; cc < 26; cc++ {
		ans = max(ans, dp(0, n-1, cc))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(longestPalindromeSubseq("bbabab"))
	fmt.Println(longestPalindromeSubseq("dcbccacdb"))
	fmt.Println(longestPalindromeSubseq("ftubcyofeqcchaobkcaocpkxfc"))
}
