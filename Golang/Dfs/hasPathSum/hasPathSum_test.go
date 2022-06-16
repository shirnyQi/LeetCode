package hasPathSum

import "testing"

/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e8gh3h/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	Node7   = &TreeNode{7, nil, nil}
	Node2   = &TreeNode{2, nil, nil}
	Node1   = &TreeNode{1, nil, nil}
	Node11  = &TreeNode{11, Node7, Node2}
	Node4   = &TreeNode{4, Node11, nil}
	Node5   = &TreeNode{5, Node4, Node8}
	Node8   = &TreeNode{8, Node13, Node4_2}
	Node13  = &TreeNode{13, nil, nil}
	Node4_2 = &TreeNode{4, nil, Node1}

	targetSum = 22
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Right == nil && root.Left == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

}

func Test_HasPathSum(t *testing.T) {
	got := hasPathSum(Node5, targetSum)
	if true != got {
		t.Errorf("want:%v, got:%v", true, got)
	}
}
