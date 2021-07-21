package sumNumbers

import (
	"fmt"
	"testing"
)

/*
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e84677/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	node2 = &TreeNode{2, nil, nil}
	node3 = &TreeNode{3, nil, nil}
	node1 = &TreeNode{1, node2, node3}
	want1 = 25
)

var (
	node1_1 = &TreeNode{1, nil, nil}
	node5   = &TreeNode{5, nil, nil}
	node0   = &TreeNode{0, nil, nil}
	node9   = &TreeNode{9, node5, node1_1}
	node4   = &TreeNode{4, node9, node0}
	want2   = 1026
)

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, s int) {
		s = s*10 + root.Val
		fmt.Println(s)
		if root.Left == nil && root.Right == nil {
			ans += s
		}
		if root.Left != nil {
			dfs(root.Left, s)
		}
		if root.Right != nil {
			dfs(root.Right, s)
		}
	}
	dfs(root, 0)

	return ans
}

func Test_SumNumbers(t *testing.T) {
	var got int

	got = sumNumbers(node4)
	if want2 != got {
		t.Errorf("want:%v, got:%v", want2, got)
	}

	got = sumNumbers(node1)
	if want1 != got {
		t.Errorf("want:%v, got:%v", want1, got)
	}
}
