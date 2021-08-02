package minWindow

import (
	"fmt"
	"math"
	"testing"
)

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minLenth(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	return b
}

func Test_minWindow(t *testing.T) {
	fmt.Println(minWindow2("cabwefgewcwaefgcf", "cae")) // cwae
	fmt.Println(minWindow2("ADOBECODEBANC", "ABC"))     // BANC
	fmt.Println(minWindow2("a", "a"))                   // a
	fmt.Println(minWindow2("a", "aa"))                  // ""
	fmt.Println(minWindow2("a", "b"))                   // ""
	fmt.Println(minWindow2("ab", "a"))                  // a
}

func minWindow2(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		// cnt记录当前子串中，t存在的每个元素的个数
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		fmt.Println(cnt)
		for check() && l <= r {
			// 当前子串比原来的短
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			// 左指针前进，但如果左指针对于的字符在t中存在，那么cnt中将该元素--，否则前进
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
