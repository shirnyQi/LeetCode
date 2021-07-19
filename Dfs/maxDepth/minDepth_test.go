package maxDepth

import "testing"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left)+1, minDepth(root.Right)+1)
}

func Test_MinDepth(t *testing.T) {
	got := minDepth(Node3)
	if 2 != got {
		t.Errorf("want:%v, got:%v", 2, got)
	}
}
