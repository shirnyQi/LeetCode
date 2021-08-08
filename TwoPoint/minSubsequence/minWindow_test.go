package minSubsequence

import (
	"fmt"
	"math"
	"testing"
)

/*
给定字符串 S and T，找出 S 中最短的（连续）子串 W ，使得 T 是 W 的 子序列 。
如果 S 中没有窗口可以包含 T 中的所有字符，返回空字符串 ""。如果有不止一个最短长度的窗口，返回开始位置最靠左的那个。
示例 1：

输入：
S = "abcdebdde", T = "bde"
输出："bcde"
解释：
"bcde" 是答案，因为它在相同长度的字符串 "bdde" 出现之前。
"deb" 不是一个更短的答案，因为在窗口中必须按顺序出现 T 中的元素。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sliding-window-and-two-pointers/riwpld/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// 双指针超时+二分法， 尽快找到左侧边界
func minWindow(s1 string, s2 string) string {
	arr1, arr2 := []byte(s1), []byte(s2)
	var res string
	var left int

	for right, _ := range arr1 {
		if checkSubSeq(arr1[left:right+1], arr2) {
			l, r := left, right
			var k int
			for l <= r {
				mid := (l + r) / 2
				if checkSubSeq(arr1[mid:right+1], arr2) {
					l = mid + 1
					k = mid
				} else {
					r = mid - 1
				}
			}
			fmt.Println(left, l, k, r, right)
			if len(res) > right-k+1 || len(res) == 0 {
				res = string(arr1[k : right+1])
			}
			left = k + 1
		}
	}

	return res
}

func checkSubSeq(s, t []byte) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return i == len(s) && j == len(t)
}

// 滑动窗口，从右向左逆向扩散左边界，不同于正常的左指针向右靠拢收缩窗口
func minWindow2(s1 string, s2 string) string {
	if s1 == "" || s2 == "" || len(s1) < len(s2) {
		return ""
	}
	rer, rel := math.MaxInt32, 0 //返回值的坐标
	i, j := 0, 0                 // 分别作为s1，s2的坐标
	for i < len(s1) {
		if s1[i] == s2[j] {
			j++
		}
		if j == len(s2) {
			right := i
			j--
			for j >= 0 {
				if s2[j] == s1[i] {
					j--
				}
				i--
			}
			i++
			if right-i < rer-rel {
				rer = right
				rel = i
			}
			j = 0
		}
		i++
	}
	if rer == math.MaxInt32 {
		return ""
	}
	return s1[rel : rer+1]
}

func TestName(t *testing.T) {
	fmt.Println(minWindow("abcdebdde", "bde"))  //bcde
	fmt.Println(minWindow2("abcdebdde", "bde")) //bcde

	s1 := "ffynmlzesdshlvugsigobutgaetsnjlizvqjdpccdylclqcbghhixpjihximvhapymfkjxyyxfwvsfyctmhwmfjyjidnfryiyajmtakisaxwglwpqaxaicuprrvxybzdxunypzofhpclqiybgniqzsdeqwrdsfjyfkgmejxfqjkmukvgygafwokeoeglanevavyrpduigitmrimtaslzboauwbluvlfqquocxrzrbvvplsivujojscytmeyjolvvyzwizpuhejsdzkfwgqdbwinkxqypaphktonqwwanapouqyjdbptqfowhemsnsl"
	s2 := "ntimcimzah"
	fmt.Println(minWindow(s1, s2)) //"nevavyrpduigitmrimtaslzboauwbluvlfqquocxrzrbvvplsivujojscytmeyjolvvyzwizpuhejsdzkfwgqdbwinkxqypaph"
	fmt.Println(minWindow2(s1, s2))
}
