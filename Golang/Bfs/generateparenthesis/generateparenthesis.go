package generateparenthesis

/*
LC T22:
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

1、左括号数目小于等于n
2、右括号数目小于等于左括号
*/

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	dfs(n-1, n, 1, "(") //递归开始
	return res
}

func dfs(m, n, tag int, s string) {
	if m == 0 && n == 0 && tag == 0 {
		res = append(res, s) //满足结果退出递归
		return
	}
	if m < 0 || n < 0 || tag < 0 {
		return //剪枝
	}
	dfs(m, n-1, tag-1, s+")") //下一层，将')'视为-1，'('视为1，当tag为负数时为非法序列
	dfs(m-1, n, tag+1, s+"(")
}
