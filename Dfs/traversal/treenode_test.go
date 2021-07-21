package traversal

import (
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	t.Logf("前序：%v", preorderTraversal(node4))
	t.Logf("前序：%v", preorderTraversalStack(node4))
}

func Test_inorderTraversal(t *testing.T) {
	t.Logf("中序：%v", inorderTraversal(node4))
	t.Logf("中序：%v", inorderTraversalStack(node4))
}

func Test_postorderTraversal(t *testing.T) {
	t.Logf("后序：%v", postorderTraversal(node4))
	t.Logf("后序：%v", postorderTraversalStack(node4))
}
