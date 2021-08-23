package verifyPostorder

import (
	"fmt"
	"testing"
)

/*
剑指 Offer 33. 二叉搜索树的后序遍历序列
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

link:https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
*/

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}

	size := len(postorder)

	root := postorder[size-1]

	// find right child and left child line
	index := 0
	for i, v := range postorder {
		if v >= root {
			index = i
			break
		}
	}

	leftChild := postorder[:index]
	rightChild := postorder[index : size-1]

	for _, v := range leftChild {
		if v > root {
			return false
		}
	}

	for _, v := range rightChild {
		if v < root {
			return false
		}
	}

	return verifyPostorder(leftChild) && verifyPostorder(rightChild)
}

func TestName(t *testing.T) {
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5})) // false
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5})) // true
}
