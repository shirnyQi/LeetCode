package traversal

// 二叉树中序遍历递归实现
func inorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return res

}

// 二叉树中序遍历栈实现
func inorderTraversalStack(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		vals = append(vals, node.Val)
		node = node.Right
		stack = stack[:len(stack)-1]
	}
	return
}
