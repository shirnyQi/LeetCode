package checkValidString

import (
	"fmt"
	"testing"
)

/*
678. 有效的括号字符串
给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

任何左括号 ( 必须有相应的右括号 )。
任何右括号 ) 必须有相应的左括号 ( 。
左括号 ( 必须在对应的右括号之前 )。
* 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
一个空字符串也被视为有效字符串。

示例 1:
输入: "()"
输出: True

示例 2:
输入: "(*)"
输出: True

示例 3:
输入: "(*))"
输出: True

link:https://leetcode-cn.com/problems/valid-parenthesis-string/
*/

func checkValidString(s string) bool {
	// 利用两个栈分别存放左括号和*下标
	var leftStk, asteriskStk []int
	for i, ch := range s {
		if ch == '(' {
			leftStk = append(leftStk, i)
		} else if ch == '*' {
			asteriskStk = append(asteriskStk, i)
		} else if len(leftStk) > 0 {
			leftStk = leftStk[:len(leftStk)-1]
		} else if len(asteriskStk) > 0 {
			asteriskStk = asteriskStk[:len(asteriskStk)-1]
		} else {
			return false
		}
	}

	// 左括号栈和星栈都有可能有剩余，如用例5，这个时候需要将星栈默认为右括号栈，且下标必须小于左括号
	i := len(leftStk) - 1
	for j := len(asteriskStk) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if leftStk[i] > asteriskStk[j] {
			return false
		}
	}
	return i < 0
}

func checkValidStringDp(s string) bool {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i, ch := range s {
		if ch == '*' {
			dp[i][i] = true
		}
	}

	for i := 1; i < n; i++ {
		c1, c2 := s[i-1], s[i]
		dp[i-1][i] = (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*')
	}

	for i := n - 3; i >= 0; i-- {
		c1 := s[i]
		for j := i + 2; j < n; j++ {
			c2 := s[j]
			if (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*') {
				dp[i][j] = dp[i+1][j-1]
			}
			for k := i; k < j && !dp[i][j]; k++ {
				dp[i][j] = dp[i][k] && dp[k+1][j]
			}
		}
	}

	return dp[0][n-1]
}

func TestName(t *testing.T) {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(*)))"))
	fmt.Println(checkValidString("(((*))"))

	fmt.Println(checkValidStringDp("()"))
	fmt.Println(checkValidStringDp("(*)"))
	fmt.Println(checkValidStringDp("(*))"))
	fmt.Println(checkValidStringDp("(*)))"))
	fmt.Println(checkValidStringDp("(((*))"))
}
