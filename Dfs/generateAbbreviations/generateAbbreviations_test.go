package generateAbbreviations

import (
	"fmt"
	"strconv"
	"testing"
)

/*
320. 列举单词的全部缩写
单词的 广义缩写词 可以通过下述步骤构造：先取任意数量的不重叠的子字符串，再用它们各自的长度进行替换。例如，"abcde" 可以缩写为 "a3e"（"bcd" 变为 "3" ），"1bcd1"（"a" 和 "e" 都变为 "1"），"23"（"ab" 变为 "2" ，"cde" 变为 "3" ）。
给你一个字符串 word ，返回一个由所有可能 广义缩写词 组成的列表。按 任意顺序 返回答案。

示例 1：
输入：word = "word"
输出：["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]

示例 2：
输入：word = "a"
输出：["1","a"]

link:https://leetcode-cn.com/problems/generalized-abbreviation/
*/

func generateAbbreviations(word string) []string {
	var res []string

	size := len(word)

	var dfs func(s string, index int, preCnt int)
	dfs = func(s string, index int, preCnt int) {
		if index == size {
			res = append(res, s)
			return
		}

		if preCnt != 0 {
			nextCnt := preCnt + 1
			nextCntStr := strconv.Itoa(nextCnt)
			if preCnt < 10 {
				dfs(string(s[:len(s)-1])+nextCntStr, index+1, nextCnt)
			} else {
				dfs(string(s[:len(s)-2])+nextCntStr, index+1, nextCnt)
			}

		} else {
			dfs(s+"1", index+1, 1)
		}

		dfs(s+string(word[index]), index+1, 0)
	}
	dfs("", 0, 0)
	return res
}

func TestName(t *testing.T) {
	fmt.Println(generateAbbreviations("word"))
	fmt.Println(generateAbbreviations("a"))
}
