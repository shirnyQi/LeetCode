package generateParenthesis

import (
	"fmt"
	"testing"
)

/*
 括号
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
说明：解集不能包含重复的子集。
例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
link:https://leetcode-cn.com/problems/bracket-lcci/
*/

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	var res []string

	var dfs func(path string, leftCnt, rightCnt int)
	dfs = func(path string, leftCnt, rightCnt int) {
		if leftCnt > n || leftCnt < rightCnt {
			return
		}
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		dfs(path+"(", leftCnt+1, rightCnt)
		dfs(path+")", leftCnt, rightCnt+1)
	}
	dfs("", 0, 0)
	return res
}

func TestName(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
