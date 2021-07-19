package maxDepth

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	Node15 = &TreeNode{15, nil, nil}
	Node7  = &TreeNode{7, nil, nil}
	Node20 = &TreeNode{20, Node15, Node7}
	Node9  = &TreeNode{9, nil, nil}
	Node3  = &TreeNode{3, Node9, Node20}

	want = 3
)

func maxDepth(root *TreeNode) int {
	return Dfs(root)
}

func Dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(Dfs(node.Left)+1, Dfs(node.Right)+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_MaxDepth(t *testing.T) {
	got := maxDepth(Node3)
	if want != got {
		t.Errorf("want:%v, got:%v", want, got)
	}
}
