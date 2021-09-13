package findKthBit

import (
	"fmt"
	"testing"
)

/*
1545. 找出第 N 个二进制字符串中的第 K 位
给你两个正整数 n 和 k，二进制字符串  Sn 的形成规则如下：

S1 = "0"
当 i > 1 时，Si = Si-1 + "1" + reverse(invert(Si-1))
其中 + 表示串联操作，reverse(x) 返回反转 x 后得到的字符串，而 invert(x) 则会翻转 x 中的每一位（0 变为 1，而 1 变为 0）。
例如，符合上述描述的序列的前 4 个字符串依次是：
S1 = "0"
S2 = "011"
S3 = "0111001"
S4 = "011100110110001"
请你返回  Sn 的 第 k 位字符 ，题目数据保证 k 一定在 Sn 长度范围以内。

示例 1：
输入：n = 3, k = 1
输出："0"
解释：S3 为 "0111001"，其第 1 位为 "0" 。

示例 2：
输入：n = 4, k = 11
输出："1"
解释：S4 为 "011100110110001"，其第 11 位为 "1" 。

示例 3：
输入：n = 1, k = 1
输出："0"

示例 4：
输入：n = 2, k = 3
输出："1"

link:https://leetcode-cn.com/problems/find-kth-bit-in-nth-binary-string/
*/

// 模拟法
func findKthBit(n int, k int) byte {
	s := []byte{'0'}

	var change func(i int)
	change = func(i int) {
		if i > n {
			return
		}
		s = append(append(s, '1'), reverse(s)...)
		change(i + 1)
	}

	change(2)
	return s[k-1]
}

func reverse(s []byte) []byte {
	tmp := make([]byte, len(s), len(s))
	l, r := 0, len(s)-1

	for l <= r {
		if l == r {
			tmp[l] = invert(s[l])
		} else {
			tmp[l] = invert(s[l])
			tmp[r] = invert(s[r])
			tmp[l], tmp[r] = tmp[r], tmp[l]
		}
		l++
		r--
	}
	return tmp
}

func invert(s byte) byte {
	if s == '0' {
		return '1'
	}
	return '0'
}

func findKthBit2(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	mid := 1 << (n - 1)

	// 如果是正中间，直接返回1
	if k == mid {
		return '1'
	}

	// 如果是前半部分，则递归未翻转前的S
	if k < mid {
		return findKthBit2(n-1, k)
	}

	// 如果是后半部，则递归翻转部分的S
	return invert(findKthBit2(n-1, (1<<n)-k))
}

func TestName(t *testing.T) {
	fmt.Println(findKthBit(3, 1))  // 48
	fmt.Println(findKthBit(4, 11)) // 49
	fmt.Println(findKthBit(1, 1))  // 48
	fmt.Println(findKthBit(2, 3))  // 49

	fmt.Println(findKthBit2(3, 1))  // 48
	fmt.Println(findKthBit2(4, 11)) // 49
	fmt.Println(findKthBit2(1, 1))  // 48
	fmt.Println(findKthBit2(2, 3))  // 49
}
