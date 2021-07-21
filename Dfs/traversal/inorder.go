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
