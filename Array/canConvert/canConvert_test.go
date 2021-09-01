package canConvert

import (
	"fmt"
	"testing"
)

/*
1153. 字符串转化
给出两个长度相同的字符串，分别是 str1 和 str2。请你帮忙判断字符串 str1 能不能在 零次 或 多次 转化后变成字符串 str2。
每一次转化时，将会一次性将 str1 中出现的 所有 相同字母变成其他 任何 小写英文字母（见示例）。
只有在字符串 str1 能够通过上述方式顺利转化为字符串 str2 时才能返回 True，否则返回 False。​​

示例 1：
输入：str1 = "aabcc", str2 = "ccdee"
输出：true
解释：将 'c' 变成 'e'，然后把 'b' 变成 'd'，接着再把 'a' 变成 'c'。注意，转化的顺序也很重要。

示例 2：
输入：str1 = "leetcode", str2 = "codeleet"
输出：false
解释：我们没有办法能够把 str1 转化为 str2。
*/

func canConvert(str1 string, str2 string) bool {
	if str1 == str2 { // 如果相等，不需要转换
		return true
	}
	m := map[byte]byte{}      // 保存 str1 中字符对应 str2 中的字符
	exists := map[byte]bool{} // 保存 str2 中不同字符的个数
	for i := 0; i < len(str1); i++ {
		if _, ok := m[str1[i]]; !ok {
			m[str1[i]] = str2[i]
			exists[str2[i]] = true
		} else if m[str1[i]] != str2[i] {
			return false
		}
	}
	return len(exists) < 26
}

func TestName(t *testing.T) {
	fmt.Println(canConvert("aabcc", "ccdee"))
	fmt.Println(canConvert("leetcode", "codeleet"))
}
